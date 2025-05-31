package main

import (
	"pessoa_service/controllers"

	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	r := mux.NewRouter()
	// Professores
	r.HandleFunc("/professores", controllers.ListarProfessores).Methods("GET")
	r.HandleFunc("/professores", controllers.CriarProfessor).Methods("POST")
	r.HandleFunc("/professores/{id}", controllers.ObterProfessor).Methods("GET")
	r.HandleFunc("/professores/{id}", controllers.AtualizarProfessor).Methods("PUT")
	r.HandleFunc("/professores/{id}", controllers.DeletarProfessor).Methods("DELETE")

	// Alunos
	r.HandleFunc("/alunos", controllers.ListarAlunos).Methods("GET")
	r.HandleFunc("/alunos", controllers.CriarAluno).Methods("POST")
	r.HandleFunc("/alunos/{id}", controllers.ObterAluno).Methods("GET")
	r.HandleFunc("/alunos/{id}", controllers.AtualizarAluno).Methods("PUT")
	r.HandleFunc("/alunos/{id}", controllers.DeletarAluno).Methods("DELETE")

	// Disciplinas (exemplo de CRUD, se quiser implementar)
	// r.HandleFunc("/disciplinas", controllers.ListarDisciplinas).Methods("GET")
	// r.HandleFunc("/disciplinas", controllers.CriarDisciplina).Methods("POST")
	// r.HandleFunc("/disciplinas/{id}", controllers.ObterDisciplina).Methods("GET")
	// r.HandleFunc("/disciplinas/{id}", controllers.AtualizarDisciplina).Methods("PUT")
	// r.HandleFunc("/disciplinas/{id}", controllers.DeletarDisciplina).Methods("DELETE")

	// Leciona
	r.HandleFunc("/leciona/{id_professor}/{id_disciplina}", controllers.VerificarLeciona).Methods("GET")
	return r
}
