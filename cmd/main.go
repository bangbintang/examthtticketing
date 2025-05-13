package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"ticketing-konser/internal/handlers"
	"ticketing-konser/internal/middleware"
	"ticketing-konser/internal/models"
	"ticketing-konser/internal/repository"
	"ticketing-konser/internal/service"
	"ticketing-konser/internal/utils"
)

func main() {
	// 1. Load konfigurasi terlebih dahulu
	if err := loadConfig(); err != nil {
		log.Printf("Warning: gagal memuat file konfigurasi, menggunakan environment variable saja: %v", err)
	}

	// 2. Inisialisasi database dan JWT
	db, jwtUtil, err := initDatabaseAndJWT()
	if err != nil {
		log.Fatalf("Gagal inisialisasi database dan JWT: %v", err)
	}

	// 3. Jalankan auto migrate
	err = db.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Event{},
		&models.Ticket{},
		&models.Transaction{},
		&models.Review{},
		&models.Notification{},
	)
	if err != nil {
		log.Fatalf("Gagal melakukan auto migrate: %v", err)
	}

	// 4. Inisialisasi dependency dan router
	router, err := initApp(db, jwtUtil)
	if err != nil {
		log.Fatalf("Gagal inisialisasi aplikasi: %v", err)
	}

	// 5. Jalankan server dengan dukungan graceful shutdown
	runServer(router)
}

// loadConfig memuat konfigurasi dari file atau environment variables
func loadConfig() error {
	viper.SetConfigName("config") // Nama file konfigurasi (tanpa ekstensi)
	viper.SetConfigType("json")   // Tipe file konfigurasi
	viper.AddConfigPath(".")      // Lokasi file konfigurasi
	// viper.AutomaticEnv()        // Aktifkan jika ingin override dengan env

	// Default values
	viper.SetDefault("SERVER_PORT", "8080")
	viper.SetDefault("JWT_SECRET", "defaultsecret")
	viper.SetDefault("DB_DSN", "postgres://root:root@localhost:5432/ticketing_db?sslmode=disable")

	// Tidak error jika file config tidak ada, gunakan env saja
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Config file not found, using environment variables only: %v", err)
	}
	return nil
}

// initDatabaseAndJWT menginisialisasi koneksi database dan utilitas JWT
func initDatabaseAndJWT() (*gorm.DB, *utils.JWTUtil, error) {
	dsn := viper.GetString("DB_DSN")
	if dsn == "" {
		return nil, nil, fmt.Errorf("DSN database tidak disetel")
	}

	jwtSecret := viper.GetString("JWT_SECRET")
	if jwtSecret == "" {
		return nil, nil, fmt.Errorf("JWT secret tidak disetel")
	}

	// Inisialisasi koneksi database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, nil, fmt.Errorf("gagal menghubungkan ke database: %v", err)
	}

	// Inisialisasi JWT util (tambahkan argumen signing method)
	jwtUtil, err := utils.NewJWTUtil(jwtSecret, time.Hour*1, time.Hour*24*7, jwt.SigningMethodHS256)
	if err != nil {
		return nil, nil, fmt.Errorf("gagal menginisialisasi JWT util: %v", err)
	}

	return db, jwtUtil, nil
}

// initApp menginisialisasi router dan dependency aplikasi
func initApp(db *gorm.DB, jwtUtil *utils.JWTUtil) (*gin.Engine, error) {
	// Inisialisasi dependency
	userHandler, authHandler, ticketHandler, reviewHandler, transactionHandler, eventHandler, notificationHandler := initDependencies(db, jwtUtil)

	// Inisialisasi router
	router := gin.Default()
	authMiddleware := middleware.NewAuthMiddleware(jwtUtil)

	// Tambahkan rute
	setupRoutes(router, userHandler, authHandler, ticketHandler, reviewHandler, transactionHandler, eventHandler, notificationHandler, authMiddleware)

	return router, nil
}

// initDependencies menginisialisasi semua repository, service, dan handler
func initDependencies(db *gorm.DB, jwtUtil *utils.JWTUtil) (
	*handlers.UserHandler,
	*handlers.AuthHandler,
	*handlers.TicketHandler,
	*handlers.ReviewHandler,
	*handlers.TransactionHandler,
	*handlers.EventHandler,
	*handlers.NotificationHandler,
) {
	// Inisialisasi repository
	userRepo := repository.NewUserRepository(db)
	ticketRepo := repository.NewTicketRepository(db)
	reviewRepo := repository.NewReviewRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)
	eventRepo := repository.NewEventRepository(db)
	notificationRepo := repository.NewNotificationRepository(db)

	// Inisialisasi service
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userService, jwtUtil)
	ticketService := service.NewTicketService(ticketRepo)
	reviewService := service.NewReviewService(reviewRepo)
	transactionService := service.NewTransactionService(transactionRepo)
	eventService := service.NewEventService(eventRepo)
	notificationService := service.NewNotificationService(notificationRepo)

	// Inisialisasi handler
	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(authService)
	ticketHandler := handlers.NewTicketHandler(ticketService)
	reviewHandler := handlers.NewReviewHandler(reviewService)
	transactionHandler := handlers.NewTransactionHandler(transactionService)
	eventHandler := handlers.NewEventHandler(eventService)
	notificationHandler := handlers.NewNotificationHandler(notificationService)

	return userHandler, authHandler, ticketHandler, reviewHandler, transactionHandler, eventHandler, notificationHandler
}

// setupRoutes menambahkan semua rute ke router
func setupRoutes(
	router *gin.Engine,
	userHandler *handlers.UserHandler,
	authHandler *handlers.AuthHandler,
	ticketHandler *handlers.TicketHandler,
	reviewHandler *handlers.ReviewHandler,
	transactionHandler *handlers.TransactionHandler,
	eventHandler *handlers.EventHandler,
	notificationHandler *handlers.NotificationHandler,
	authMiddleware *middleware.AuthMiddleware,
) {
	// Rute publik (tidak membutuhkan autentikasi)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})

	// Rute yang membutuhkan autentikasi
	protected := router.Group("/api")
	protected.Use(authMiddleware.Middleware())
	{
		// protected.GET("/profile", userHandler.GetProfile)
		// protected.POST("/update", userHandler.UpdateProfile)
	}

	// Routing user (publik)
	userGroup := router.Group("/users")
	{
		userGroup.POST("/register", userHandler.RegisterUser)
	}

	// Routing auth (publik)
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", authHandler.Login)
		authGroup.POST("/refresh", authHandler.RefreshToken)
	}

	// Routing tiket (butuh autentikasi)
	ticketGroup := router.Group("/tickets")
	ticketGroup.Use(authMiddleware.Middleware())
	{
		ticketGroup.POST("/purchase", ticketHandler.PurchaseTicket)
		ticketGroup.GET("/user/:userID", ticketHandler.GetTicketsByUser)
	}

	// Routing review (GET publik, POST butuh autentikasi)
	reviewGroup := router.Group("/reviews")
	{
		reviewGroup.GET("/event/:eventID", reviewHandler.GetReviewsByEvent)
	}

	reviewGroupAuth := router.Group("/reviews")
	reviewGroupAuth.Use(authMiddleware.Middleware())
	{
		reviewGroupAuth.POST("/", reviewHandler.CreateReview)
	}

	// Routing transaksi (butuh autentikasi)
	transactionGroup := router.Group("/transactions")
	transactionGroup.Use(authMiddleware.Middleware())
	{
		transactionGroup.POST("/", transactionHandler.CreateTransaction)
		transactionGroup.GET("/user/:userID", transactionHandler.GetTransactionsByUser)
	}

	// Routing event (butuh autentikasi admin)
	eventGroup := router.Group("/events")
	eventGroup.Use(authMiddleware.Middleware())
	{
		// eventGroup.Use(middleware.RBACMiddleware(middleware.RoleAdmin))
		eventGroup.POST("/", eventHandler.CreateEvent)
		// eventGroup.GET("/", eventHandler.GetAllEvents)
		// eventGroup.PUT("/:id", eventHandler.UpdateEvent)
		// eventGroup.DELETE("/:id", eventHandler.DeleteEvent)
	}

	// Routing notifikasi (butuh autentikasi)
	notificationGroup := router.Group("/notifications")
	notificationGroup.Use(authMiddleware.Middleware())
	{
		notificationGroup.POST("/", notificationHandler.CreateNotification)
		notificationGroup.GET("/user/:userID", notificationHandler.GetNotificationsByUser)
	}
}

// runServer menjalankan server dengan dukungan graceful shutdown
func runServer(router *gin.Engine) {
	port := viper.GetString("SERVER_PORT")
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}

	// Jalankan server di goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Tunggu sinyal untuk shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Shutdown server dengan timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
