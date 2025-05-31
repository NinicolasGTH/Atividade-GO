package main

import (
	"atividade_service/controllers"

	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/atividades", controllers.ListarAtividades).Methods("GET")
	r.HandleFunc("/atividades", controllers.CriarAtividade).Methods("POST")
	r.HandleFunc("/atividades/{id}", controllers.ObterAtividade).Methods("GET")
	r.HandleFunc("/atividades/{id}", controllers.AtualizarAtividade).Methods("PUT")
	r.HandleFunc("/atividades/{id}", controllers.DeletarAtividade).Methods("DELETE") // <-- Adicione esta linha
	r.HandleFunc("/atividades/{id_atividade}/professor/{id_professor}", controllers.ObterAtividadeParaProfessor).Methods("GET")
	return r
}
