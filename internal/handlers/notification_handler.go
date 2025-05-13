package handlers

import (
	"net/http"
	"ticketing-konser/internal/models"
	"ticketing-konser/internal/service"

	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	NotificationService *service.NotificationService
}

func NewNotificationHandler(notificationService *service.NotificationService) *NotificationHandler {
	return &NotificationHandler{NotificationService: notificationService}
}

func (h *NotificationHandler) CreateNotification(c *gin.Context) {
	var notification models.Notification
	if err := c.ShouldBindJSON(&notification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdNotification, err := h.NotificationService.CreateNotification(&notification)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdNotification)
}

func (h *NotificationHandler) GetNotificationsByUser(c *gin.Context) {
	userID := c.Param("userID")
	notifications, err := h.NotificationService.GetNotificationsByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, notifications)
}
