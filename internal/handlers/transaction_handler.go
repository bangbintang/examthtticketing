package handlers

import (
	"net/http"
	"strconv"
	"ticketing-konser/internal/models"
	"ticketing-konser/internal/service"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	TransactionService service.ITransactionService
}

func NewTransactionHandler(transactionService service.ITransactionService) *TransactionHandler {
	return &TransactionHandler{TransactionService: transactionService}
}

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTransaction, err := h.TransactionService.CreateTransaction(&transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdTransaction)
}

func (h *TransactionHandler) GetTransactionsByUser(c *gin.Context) {
	userID := c.Param("userID")
	transactions, err := h.TransactionService.GetTransactionsByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func (h *TransactionHandler) GetTransactionByID(c *gin.Context) {
	idStr := c.Param("transactionID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "transactionID harus berupa angka"})
		return
	}

	transaction, err := h.TransactionService.GetTransactionByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func (h *TransactionHandler) GetTransactionsByEvent(c *gin.Context) {
	eventIDStr := c.Param("eventID")
	eventID, err := strconv.Atoi(eventIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "eventID harus berupa angka"})
		return
	}

	transactions, err := h.TransactionService.GetTransactionsByEvent(eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}
