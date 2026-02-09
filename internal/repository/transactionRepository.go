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

func (r *TransactionRepository) GetByID(ctx context.Context, id int32) (db.Transaction, error) {
	return r.q.GetTransaction(ctx, id)
}

func (r *TransactionRepository) Delete(ctx context.Context, id int32) error {
	return r.q.DeleteTransaction(ctx, id)
}

func (r *TransactionRepository) List(ctx context.Context) ([]db.Transaction, error) {
	return r.q.ListTransactions(ctx)
}

func (r *TransactionRepository) Update(ctx context.Context, params db.UpdateTransactionParams) (db.Transaction, error) {
	return r.q.UpdateTransaction(ctx, params)
}