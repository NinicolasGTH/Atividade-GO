package main

import (
	"log"
	"net/http"

	"pessoa_service/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("banco_pessoa.db"), &gorm.Config{})
	if err != nil {
		panic("Erro ao conectar ao banco")
	}
	// Migrar o model Pessoa
	database.AutoMigrate(&models.Pessoa{})
	models.DB = database
}

func StartServer() {
	ConnectDatabase()
	router := CreateRouter()
	log.Println("Pessoa Service rodando na porta 8081...")
	err := http.ListenAndServe(":8081", router)
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
