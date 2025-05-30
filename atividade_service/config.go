package main

import (
	"atividade_service/models"
	"log"
	"net/http"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// Model de exemplo (pode remover se não for usar)
type Sala struct {
	ID   uint   `gorm:"primaryKey"`
	Nome string `json:"nome"`
}

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("banco.db"), &gorm.Config{})
	if err != nil {
		panic("Erro ao conectar ao banco")
	}
	// Migrar os models principais
	database.AutoMigrate(&models.Atividade{}, &models.Resposta{})
	// Atribuir o ponteiro do banco ao pacote models
	models.DB = database
}

func CreateApp() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/exemplo", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Rota exemplo funcionando!"))
	})
	return mux
}

func StartServer() {
	ConnectDatabase() // Conecta ao banco antes de iniciar o servidor
	app := CreateRouter()
	log.Println("Servidor rodando em modo debug na porta 8080...")
	err := http.ListenAndServe(":8080", app)
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
