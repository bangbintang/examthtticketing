package handlers

import (
	"log"
	"net/http"
	"ticketing-konser/internal/service"
	"ticketing-konser/internal/utils"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	// Gunakan instance logger dari utils
	logger := utils.DefaultLoggerInstance

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Input tidak valid", err.Error(), logger)
		return
	}

	token, err := h.AuthService.Login(input.Email, input.Password)
	if err != nil {
		log.Printf("Login gagal untuk email %s: %v", input.Email, err)
		utils.RespondError(c, http.StatusUnauthorized, "Login gagal", err.Error(), logger)
		return
	}

	log.Printf("Login berhasil untuk email %s", input.Email)
	utils.RespondSuccess(c, http.StatusOK, "Login berhasil", gin.H{
		"token": token,
	})
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var input struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	// Gunakan instance logger dari utils
	logger := utils.DefaultLoggerInstance

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Input tidak valid", err.Error(), logger)
		return
	}

	token, err := h.AuthService.RefreshToken(input.RefreshToken)
	if err != nil {
		log.Printf("Refresh token gagal: %v", err)
		utils.RespondError(c, http.StatusUnauthorized, "Refresh token gagal", err.Error(), logger)
		return
	}

	log.Println("Refresh token berhasil")
	utils.RespondSuccess(c, http.StatusOK, "Token berhasil diperbarui", gin.H{
		"token": token,
	})
}
