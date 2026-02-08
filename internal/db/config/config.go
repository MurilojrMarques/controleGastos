package config

import (
	"context"
	"fmt"
	"log"
	"os" 

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresDB struct {
	Db *pgxpool.Pool
}

func NewPostgresDB() (*PostgresDB, error) {
	dbUrl := os.Getenv("DATABASE_URL")
    

    if dbUrl == "" {
        return nil, fmt.Errorf("a variável DATABASE_URL é obrigatória")
    }

	dbPool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar pool de conexões: %w", err)
	}

	if err := dbPool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("erro ao pingar o banco: %w", err)
	}

	log.Println("Conectado no banco com sucesso!")
	return &PostgresDB{Db: dbPool}, nil
}