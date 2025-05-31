package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"pessoa_service/models"

	"github.com/gorilla/mux"
)

// PROFESSORES

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

// POST /professores
func CriarProfessor(w http.ResponseWriter, r *http.Request) {
	var professor models.Professor
	if err := json.NewDecoder(r.Body).Decode(&professor); err != nil {
		http.Error(w, `{"erro":"JSON inválido"}`, http.StatusBadRequest)
		return
	}
	if err := models.DB.Create(&professor).Error; err != nil {
		http.Error(w, `{"erro":"Erro ao criar professor"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(professor)
}

// GET /professores/{id}
func ObterProfessor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"erro":"ID inválido"}`, http.StatusBadRequest)
		return
	}
	var professor models.Professor
	if err := models.DB.First(&professor, id).Error; err != nil {
		http.Error(w, `{"erro":"Professor não encontrado"}`, http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(professor)
}

// PUT /professores/{id}
func AtualizarProfessor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"erro":"ID inválido"}`, http.StatusBadRequest)
		return
	}
	var professor models.Professor
	if err := json.NewDecoder(r.Body).Decode(&professor); err != nil {
		http.Error(w, `{"erro":"JSON inválido"}`, http.StatusBadRequest)
		return
	}
	professor.ID = uint(id)
	if err := models.DB.Save(&professor).Error; err != nil {
		http.Error(w, `{"erro":"Erro ao atualizar professor"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(professor)
}

// DELETE /professores/{id}
func DeletarProfessor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"erro":"ID inválido"}`, http.StatusBadRequest)
		return
	}
	if err := models.DB.Delete(&models.Professor{}, id).Error; err != nil {
		http.Error(w, `{"erro":"Erro ao deletar professor"}`, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// ALUNOS

// GET /alunos
func ListarAlunos(w http.ResponseWriter, r *http.Request) {
	alunos, err := models.ListarAlunos()
	if err != nil {
		fmt.Println("Erro ao listar alunos:", err)
		http.Error(w, `{"erro":"Erro ao listar alunos"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(alunos)
}

// POST /alunos
func CriarAluno(w http.ResponseWriter, r *http.Request) {
	var aluno models.Aluno
	if err := json.NewDecoder(r.Body).Decode(&aluno); err != nil {
		http.Error(w, `{"erro":"JSON inválido"}`, http.StatusBadRequest)
		return
	}
	if err := models.DB.Create(&aluno).Error; err != nil {
		http.Error(w, `{"erro":"Erro ao criar aluno"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(aluno)
}

// GET /alunos/{id}
func ObterAluno(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"erro":"ID inválido"}`, http.StatusBadRequest)
		return
	}
	var aluno models.Aluno
	if err := models.DB.First(&aluno, id).Error; err != nil {
		http.Error(w, `{"erro":"Aluno não encontrado"}`, http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(aluno)
}

// PUT /alunos/{id}
func AtualizarAluno(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"erro":"ID inválido"}`, http.StatusBadRequest)
		return
	}
	var aluno models.Aluno
	if err := json.NewDecoder(r.Body).Decode(&aluno); err != nil {
		http.Error(w, `{"erro":"JSON inválido"}`, http.StatusBadRequest)
		return
	}
	aluno.ID = uint(id)
	if err := models.DB.Save(&aluno).Error; err != nil {
		http.Error(w, `{"erro":"Erro ao atualizar aluno"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(aluno)
}

// DELETE /alunos/{id}
func DeletarAluno(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"erro":"ID inválido"}`, http.StatusBadRequest)
		return
	}
	if err := models.DB.Delete(&models.Aluno{}, id).Error; err != nil {
		http.Error(w, `{"erro":"Erro ao deletar aluno"}`, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// LECCIONA

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
