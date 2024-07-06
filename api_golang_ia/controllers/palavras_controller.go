package controllers

import (
	"api_golang_ia/models"

	"github.com/gin-gonic/gin"
)

type PalavrasController struct {
}

func (pc PalavrasController) Index(c *gin.Context) {
	palavras := []models.Palavra{
		{
			Palavra:  "Cachorro",
			Traducao: "Dog",
			Opcoes:   []string{"Cat", "Dog", "Elephant"},
		},
		{
			Palavra:  "Gato",
			Traducao: "Cat",
			Opcoes:   []string{"Dog", "Cat", "Elephant"},
		},
		{
			Palavra:  "Elefante",
			Traducao: "Elephant",
			Opcoes:   []string{"Dog", "Cat", "Elephant"},
		},
	}

	c.JSON(200, palavras)
}
