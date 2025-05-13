package handlers

import (
    "net/http"
    "ticketing-konser/internal/models"
    "ticketing-konser/internal/service"

    "github.com/gin-gonic/gin"
)

type UserHandler struct {
    UserService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
    return &UserHandler{UserService: userService}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    createdUser, err := h.UserService.RegisterUser(&user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "User berhasil didaftarkan",
        "data":    createdUser,
    })
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
    id := c.Param("id")
    user, err := h.UserService.GetUserByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
        return
    }

    c.JSON(http.StatusOK, user)
}