package services

import (
	"api_golang_ia/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type IAService struct {
}

func (ias *IAService) BuscaPalavras() []models.Palavra {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	openAIKey := os.Getenv("OPEN_AI_KEY")
	openAIUrl := os.Getenv("OPEN_AI_URL")

	if openAIKey == "" {
		log.Fatal("A variável OPEN_AI_KEY não foi definida")
		return nil
	}

	prompt := "Traga para mim um conjunto de palavras em inglês"
	data := map[string]interface{}{
		"model":      "gpt-3.5-turbo-instruct",
		"prompt":     prompt,
		"max_tokens": 100,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal("Erro ao serializar dados: ", err)
		return nil
	}

	req, err := http.NewRequest("POST", openAIUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal("Erro ao criar a requisição: ", err)
		return nil
	}

	req.Header.Set("Authorization", "Bearer "+openAIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Erro ao fazer a requisição: ", err)
		return nil
	}
	defer resp.Body.Close()

	var response map[string]interface{}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Erro ao ler a resposta: ", err)
		return nil
	}

	fmt.Printf("Resposta: %s\n", body)

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal("Erro ao desserializar a resposta: ", err)
		return nil
	}

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
