package main

import (
	"log"

	"github.com/MuriloJrMarques/financas-api/internal/db/config"
	"github.com/gin-gonic/gin"
)


func main(){
	postgresDB, err := config.NewPostgresDB()
	if err != nil {
		log.Fatalf("Falha ao conectar com o banco de dados: %v", err)
	}
	defer postgresDB.Db.Close()

	
	server := gin.Default()

	server.Use(gin.Logger())
	server.Use(gin.Recovery())

	v1 := server.Group("/api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	if err := server.Run(":8080"); err != nil {
		log.Fatalf("Falha ao iniciar o servidor: %v", err)
	}
	log.Println("Servidor iniciado na porta 8080")
}