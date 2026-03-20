package main

import (
	_ "gre-api/docs" // Substitua "gre-api" pelo nome do seu módulo
	"gre-api/internal/handlers"

	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API de Geração de Boletos
// @version 1.0
// @description API que recebe dados e gera um boleto em HTML.
// @host localhost:5005
// @BasePath /
func main() {
	r := gin.Default()

	r.POST("/boleto", handlers.RenderBoleto)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("Servidor rodando em http://localhost:6000")
	if err := r.Run(":5005"); err != nil {
		log.Fatal(err)
	}
}
