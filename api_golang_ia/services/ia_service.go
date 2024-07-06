package services

import "api_golang_ia/models"

type IAService struct {
}

func (ias *IAService) BuscaPalavras() []models.Palavra {
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

	return palavras

}
