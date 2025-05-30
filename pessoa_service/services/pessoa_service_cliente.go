package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const pessoaServiceURL = "http://localhost:8081/leciona"

type LecionaResponse struct {
	Leciona bool `json:"leciona"`
}

func VerificaLeciona(idProfessor, idDisciplina int) (bool, string) {
	url := fmt.Sprintf("%s/%d/%d", pessoaServiceURL, idProfessor, idDisciplina)

	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return false, fmt.Sprintf("Erro na comunicação com pessoa_service: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return false, "Disciplina não encontrada"
	}

	var result LecionaResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, fmt.Sprintf("Erro ao decodificar resposta: %v", err)
	}

	return result.Leciona, ""
}
