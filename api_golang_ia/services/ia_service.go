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
	openAIModel := os.Getenv("OPEN_AI_MODEL")

	if openAIKey == "" {
		log.Fatal("A variável OPEN_AI_KEY não foi definida")
		return nil
	}

	messages := []map[string]string{
		{
			"role":    "system",
			"content": "Sua missão é retornar para mim uma lista de palavras em inglês com quatro alternativas sendo uma correta. O formato da lista retornada será assim: `[{\"palavra\": \"Hello\", \"traducao\": \"Olá\", \"opcoes\": [\"Boa\", \"Ok\", \"Olá\", \"Bacana\"]}]`. Somente o JSON deve ser retornado com nenhum outro texto. Não utilizar formatação markdown.",
		},
		{
			"role":    "user",
			"content": "Traga para mim o campo 'palavra' em inglês, o campo 'traducao' em português e o campo 'opcoes', uma lista de quatro itens em português.",
		},
	}

	requestBody, err := json.Marshal(map[string]interface{}{
		"model":    openAIModel,
		"messages": messages,
	})

	if err != nil {
		log.Fatal("Erro ao montar a requisição: ", err)
	}

	req, _ := http.NewRequest("POST", openAIUrl, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+openAIKey)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal("Erro ao enviar requisição: ", err)
	}

	defer resp.Body.Close()

	reponseBody, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal("Erro ao ler a resposta: ", err)
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(reponseBody), &result)

	var palavras []models.Palavra

	if choices, ok := result["choices"].([]interface{}); ok && len(choices) > 0 {
		if firstChoice, ok := choices[0].(map[string]interface{}); ok {
			if message, ok := firstChoice["message"].(map[string]interface{}); ok {
				resposta := message["content"].(string)
				fmt.Println("Resposta: ", resposta)

				err := json.Unmarshal([]byte(resposta), &palavras)
				if err != nil {
					log.Fatal("Erro ao desserializar resposta: ", err)
					return nil
				}

				return palavras
			}
		}
	}

	return nil
}
