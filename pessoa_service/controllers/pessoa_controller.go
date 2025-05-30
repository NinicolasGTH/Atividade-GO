package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"pessoa_service/models"

	"github.com/gorilla/mux"
)

// GET /professores
func ListarProfessores(w http.ResponseWriter, r *http.Request) {
	professores, err := models.ListarProfessores()
	if err != nil {
		http.Error(w, `{"erro":"Erro ao listar professores"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(professores)
}

// GET /alunos
func ListarAlunos(w http.ResponseWriter, r *http.Request) {
	alunos, err := models.ListarAlunos()
	if err != nil {
		http.Error(w, `{"erro":"Erro ao listar alunos"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(alunos)
}

// GET /leciona/{id_professor}/{id_disciplina}
func VerificarLeciona(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idProfessorStr := vars["id_professor"]
	idDisciplinaStr := vars["id_disciplina"]

	idProfessor, err := strconv.Atoi(idProfessorStr)
	if err != nil {
		http.Error(w, `{"erro":"ID de professor inválido"}`, http.StatusBadRequest)
		return
	}
	idDisciplina, err := strconv.Atoi(idDisciplinaStr)
	if err != nil {
		http.Error(w, `{"erro":"ID de disciplina inválido"}`, http.StatusBadRequest)
		return
	}

	leciona, err := models.Leciona(idProfessor, idDisciplina)
	if err != nil {
		http.Error(w, `{"erro":"Disciplina não encontrada"}`, http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"leciona": leciona})
}

// Função para registrar as rotas no router
func RegisterPessoaRoutes(r *mux.Router) {
	r.HandleFunc("/professores", ListarProfessores).Methods("GET")
	r.HandleFunc("/alunos", ListarAlunos).Methods("GET")
	r.HandleFunc("/leciona/{id_professor}/{id_disciplina}", VerificarLeciona).Methods("GET")
}
