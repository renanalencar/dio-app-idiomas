package controllers

import (
	"api_golang_ia/services"

	"github.com/gin-gonic/gin"
)

type PalavrasController struct {
}

func (pc PalavrasController) Index(c *gin.Context) {
	servico := services.IAService{}

	c.JSON(200, servico.BuscaPalavras())
}
