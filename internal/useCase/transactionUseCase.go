package usecase

import (
	"context"
	"errors"
	"fmt" 
	"time"

	"github.com/MuriloJrMarques/financas-api/internal/db"
	"github.com/MuriloJrMarques/financas-api/internal/repository"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateTransactionInput struct {
	Title   string
	Amount  float64
	Type    string
	DueDate time.Time
}

type TransactionUseCase struct {
	repo *repository.TransactionRepository
}

func NewTransactionUseCase(repo *repository.TransactionRepository) *TransactionUseCase {
	return &TransactionUseCase{repo: repo}
}

func (uc *TransactionUseCase) Create(ctx context.Context, input CreateTransactionInput) (db.Transaction, error) {
	if input.Amount <= 0 {
		return db.Transaction{}, errors.New("o valor deve ser maior que zero")
	}

	if input.Type != "income" && input.Type != "expense" && input.Type != "INCOME" && input.Type != "EXPENSE" {
		return db.Transaction{}, errors.New("o tipo deve ser 'income' ou 'expense'")
	}

	
	var amountNumeric pgtype.Numeric
    amountStr := fmt.Sprintf("%.2f", input.Amount)
    
    err := amountNumeric.Scan(amountStr)
    if err != nil {
        return db.Transaction{}, fmt.Errorf("erro ao converter valor: %v", err)
    }

	args := db.CreateTransactionParams{
		Title:   input.Title,
		Amount:  amountNumeric, 
		Type:    input.Type,
		DueDate: pgtype.Date{Time: input.DueDate, Valid: true},
	}

	transaction, err := uc.repo.Create(ctx, args)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			return db.Transaction{}, errors.New("erro de banco de dados: " + pgErr.Message)
		}
		return db.Transaction{}, err
	}

	return transaction, nil
}