package repository

import (
    "context"
    "github.com/MuriloJrMarques/financas-api/internal/db"
)

type TransactionRepository struct {
    q *db.Queries 
}

func NewTransactionRepository(q *db.Queries) *TransactionRepository {
    return &TransactionRepository{q: q}
}

func (r *TransactionRepository) Create(ctx context.Context, params db.CreateTransactionParams) (db.Transaction, error) {
    return r.q.CreateTransaction(ctx, params)
}