package controller

import (
	"net/http"
	"time"

	usecase "github.com/MuriloJrMarques/financas-api/internal/useCase"
	"github.com/gin-gonic/gin"
)


type TransactionController struct {
	transactionUseCase *usecase.TransactionUseCase
}

func NewTransactionController(transactionUseCase *usecase.TransactionUseCase) *TransactionController {
	return &TransactionController{transactionUseCase: transactionUseCase}
}

type createTransactionRequest struct {
	Title   string  `json:"title" binding:"required"`
	Amount  float64 `json:"amount" binding:"required,gt=0"`
	Type    string  `json:"type" binding:"required,oneof=income expense"`
	DueDate string  `json:"due_date" binding:"required"`
}

func (c *TransactionController) CreateTransaction(ctx *gin.Context) {
	var req createTransactionRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	parsedDate, err := time.Parse("2006-01-02", req.DueDate)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Data inválida. Use o formato YYYY-MM-DD"})
        return
    }

	input := usecase.CreateTransactionInput{
        Title:   req.Title,
        Amount:  req.Amount,
        Type:    req.Type,
        DueDate: parsedDate,
    }

	transaction, err := c.transactionUseCase.Create(ctx.Request.Context(), input)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, transaction)
}