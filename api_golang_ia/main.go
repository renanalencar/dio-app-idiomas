package main

import (
	"api_golang_ia/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	controllers := controllers.HomeController{}
	r.GET("/", controllers.Index)

	r.Run(":8888") // por padrão, o servidor é executado na porta 8080
}
