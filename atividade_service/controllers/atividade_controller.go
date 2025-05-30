package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"atividade_service/clients"
	"atividade_service/models"

	"github.com/gorilla/mux"
)

// GET /atividades
func ListarAtividades(w http.ResponseWriter, r *http.Request) {
	atividades, err := models.ListarAtividades()
	if err != nil {
		http.Error(w, `{"erro":"Erro ao listar atividades"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(atividades)
}

// GET /atividades/{id}
func ObterAtividade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"erro":"ID inválido"}`, http.StatusBadRequest)
		return
	}
	atividade, err := models.ObterAtividade(uint(id))
	if err != nil {
		http.Error(w, `{"erro":"Atividade não encontrada"}`, http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(atividade)
}

// POST /atividades
func CriarAtividade(w http.ResponseWriter, r *http.Request) {
	var atividade models.Atividade
	if err := json.NewDecoder(r.Body).Decode(&atividade); err != nil {
		http.Error(w, `{"erro":"JSON inválido"}`, http.StatusBadRequest)
		return
	}
	if err := models.DB.Create(&atividade).Error; err != nil {
		http.Error(w, `{"erro":"Erro ao criar atividade"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(atividade)
}

// PUT /atividades/{id}
func AtualizarAtividade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"erro":"ID inválido"}`, http.StatusBadRequest)
		return
	}
	var atividade models.Atividade
	if err := json.NewDecoder(r.Body).Decode(&atividade); err != nil {
		http.Error(w, `{"erro":"JSON inválido"}`, http.StatusBadRequest)
		return
	}
	atividade.ID = uint(id)
	if err := models.DB.Save(&atividade).Error; err != nil {
		http.Error(w, `{"erro":"Erro ao atualizar atividade"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(atividade)
}

// GET /atividades/{id_atividade}/professor/{id_professor}
func ObterAtividadeParaProfessor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idAtividadeStr := vars["id_atividade"]
	idProfessorStr := vars["id_professor"]

	idAtividade, err := strconv.Atoi(idAtividadeStr)
	if err != nil {
		http.Error(w, `{"erro":"ID de atividade inválido"}`, http.StatusBadRequest)
		return
	}

	atividade, err := models.ObterAtividade(uint(idAtividade))
	if err != nil {
		http.Error(w, `{"erro":"Atividade não encontrada"}`, http.StatusNotFound)
		return
	}

	if !clients.VerificarLeciona(idProfessorStr, strconv.FormatUint(uint64(atividade.IdDisciplina), 10)) {
		atividade.Respostas = nil
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(atividade)
}

// Função para registrar as rotas no router
func RegisterAtividadeRoutes(r *mux.Router) {
	r.HandleFunc("/atividades", ListarAtividades).Methods("GET")
	r.HandleFunc("/atividades", CriarAtividade).Methods("POST")
	r.HandleFunc("/atividades/{id}", ObterAtividade).Methods("GET")
	r.HandleFunc("/atividades/{id}", AtualizarAtividade).Methods("PUT")
	r.HandleFunc("/atividades/{id_atividade}/professor/{id_professor}", ObterAtividadeParaProfessor).Methods("GET")
}
