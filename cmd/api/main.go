package main

import (
	"log"

	"github.com/MuriloJrMarques/financas-api/internal/controller"
	"github.com/MuriloJrMarques/financas-api/internal/db"
	"github.com/MuriloJrMarques/financas-api/internal/db/config"
	"github.com/MuriloJrMarques/financas-api/internal/repository"
	usecase "github.com/MuriloJrMarques/financas-api/internal/useCase"
	"github.com/gin-gonic/gin"
)


func main(){
	postgresDB, err := config.NewPostgresDB()
	if err != nil {
		log.Fatalf("Falha ao conectar com o banco de dados: %v", err)
	}
	defer postgresDB.Db.Close()

	queries := db.New(postgresDB.Db)

	transactionRepo := repository.NewTransactionRepository(queries)
	
	transactionUseCase := usecase.NewTransactionUseCase(transactionRepo)

	transactionController := controller.NewTransactionController(transactionUseCase)
	server := gin.Default()

	server.Use(gin.Logger())
	server.Use(gin.Recovery())

	v1 := server.Group("/api/v1")
	{
		v1.POST("/transactions", transactionController.CreateTransaction)
	}

	if err := server.Run(":8080"); err != nil {
		log.Fatalf("Falha ao iniciar o servidor: %v", err)
	}
	log.Println("Servidor iniciado na porta 8080")
}