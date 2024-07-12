package main

import (
	"api_golang_ia/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.Default())

	homeController := controllers.HomeController{}
	r.GET("/", homeController.Index)

	palavraController := controllers.PalavrasController{}
	r.GET("/palavras", palavraController.Index)

	r.Run(":8888") // por padrão, o servidor é executado na porta 8080
}
