package controllers

import (
	"github.com/gin-gonic/gin"
)

type HomeController struct {
}

func (hc HomeController) Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Esta é uma API construída na Live Code The Future Evento DIO",
	})
}
