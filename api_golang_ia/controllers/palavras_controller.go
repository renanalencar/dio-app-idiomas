package controllers

import (
	"api_golang_ia/models"

	"github.com/gin-gonic/gin"
)

type PalavrasController struct {
}

func (pc PalavrasController) Index(c *gin.Context) {
	c.JSON(200, models.Mensagem{
		Mensagem: "Aqui vai retornar a lista de palavras",
	})
}
