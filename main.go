package main

import (
	_ "gre-api/docs" // Substitua "gre-api" pelo nome do seu módulo
	"gre-api/internal/handlers"
	"time"

	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API de Geração de Boletos
// @version 1.0
// @description API que recebe dados e gera um boleto em HTML.
// @host
// @BasePath /
func main() {
	r := gin.Default()

	// r.Use(cors.Default())
	//cors para liberar específico
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{
			"Content-Length",
			"Content-Disposition",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/boleto", handlers.RenderBoleto)
	r.POST("/retorno/read", handlers.ReadRetorno)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("Servidor rodando em http://localhost:5005")
	if err := r.Run(":5005"); err != nil {
		log.Fatal(err)
	}
}
