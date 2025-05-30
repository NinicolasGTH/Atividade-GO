package clients

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const pessoaServiceURL = "http://localhost:5001/pessoas"

type lecionaResponse struct {
	IsOk    bool `json:"isok"`
	Leciona bool `json:"leciona"`
}

func VerificarLeciona(idProfessor, idDisciplina string) bool {
	url := fmt.Sprintf("%s/leciona/%s/%s", pessoaServiceURL, idProfessor, idDisciplina)

	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("Erro ao acessar o pessoa_service: %v\n", err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Erro HTTP ao acessar o pessoa_service: %v\n", resp.Status)
		return false
	}

	var result lecionaResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Printf("Erro ao decodificar resposta do pessoa_service: %v\n", err)
		return false
	}

	if result.IsOk {
		return result.Leciona
	}
	return false
}
