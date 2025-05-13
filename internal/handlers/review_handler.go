package handlers

import (
    "net/http"
    "strconv"
    "ticketing-konser/internal/models"
    "ticketing-konser/internal/service"

    "github.com/gin-gonic/gin"
)

type ReviewHandler struct {
    ReviewService *service.ReviewService
}

func NewReviewHandler(reviewService *service.ReviewService) *ReviewHandler {
    return &ReviewHandler{ReviewService: reviewService}
}

func (h *ReviewHandler) CreateReview(c *gin.Context) {
    var review models.Review
    if err := c.ShouldBindJSON(&review); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    createdReview, err := h.ReviewService.CreateReview(&review)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, createdReview)
}

func (h *ReviewHandler) GetReviewsByEvent(c *gin.Context) {
    eventID, err := strconv.Atoi(c.Param("eventID"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID event tidak valid"})
        return
    }

    reviews, err := h.ReviewService.GetReviewsByEvent(eventID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, reviews)
}