package handlers

import (
    "net/http"
    "strconv"
    "ticketing-konser/internal/models"
    "ticketing-konser/internal/service"

    "github.com/gin-gonic/gin"
)

type EventHandler struct {
    EventService *service.EventService
}

func NewEventHandler(eventService *service.EventService) *EventHandler {
    return &EventHandler{EventService: eventService}
}

func (h *EventHandler) CreateEvent(c *gin.Context) {
    var event models.Event
    if err := c.ShouldBindJSON(&event); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    createdEvent, err := h.EventService.CreateEvent(&event)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, createdEvent)
}

func (h *EventHandler) GetEventByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
        return
    }

    event, err := h.EventService.GetEventByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Event tidak ditemukan"})
        return
    }

    c.JSON(http.StatusOK, event)
}