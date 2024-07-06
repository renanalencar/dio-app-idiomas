package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Olá, mundo!",
		})
	})

	r.Run(":8888") // por padrão, o servidor é executado na porta 8080
}
