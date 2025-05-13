package handlers

import (
	"net/http"
	"strconv"
	"ticketing-konser/internal/models"
	"ticketing-konser/internal/service"

	"github.com/gin-gonic/gin"
)

type TicketHandler struct {
	TicketService *service.TicketService
}

func NewTicketHandler(ticketService *service.TicketService) *TicketHandler {
	return &TicketHandler{TicketService: ticketService}
}

func (h *TicketHandler) PurchaseTicket(c *gin.Context) {
	var ticket models.Ticket
	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTicket, err := h.TicketService.PurchaseTicket(&ticket)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdTicket)
}

func (h *TicketHandler) GetTicketsByUser(c *gin.Context) {
	userID := c.Param("userID")
	tickets, err := h.TicketService.GetTicketsByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tickets)
}

func (h *TicketHandler) GetTicketsByEvent(c *gin.Context) {
	eventIDStr := c.Param("eventID")
	eventID, err := strconv.Atoi(eventIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "eventID harus berupa angka"})
		return
	}

	tickets, err := h.TicketService.GetTicketsByEvent(eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tickets)
}

func (h *TicketHandler) CancelTicket(c *gin.Context) {
	ticketIDStr := c.Param("ticketID")
	ticketID, err := strconv.Atoi(ticketIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ticketID harus berupa angka"})
		return
	}

	if err := h.TicketService.CancelTicket(ticketID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tiket berhasil dibatalkan"})
}
