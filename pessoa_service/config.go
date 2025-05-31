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
	database.AutoMigrate(&models.Professor{}, &models.Aluno{}, &models.Disciplina{})
	models.DB = database

	// Popular o banco apenas se estiver vazio
	var count int64
	database.Model(&models.Professor{}).Count(&count)
	if count == 0 {
		// Professores
		prof1 := models.Professor{Nome: "joao"}
		prof2 := models.Professor{Nome: "jose"}
		prof3 := models.Professor{Nome: "maria"}
		database.Create(&prof1)
		database.Create(&prof2)
		database.Create(&prof3)

		// Alunos
		aluno1 := models.Aluno{Nome: "alexandre"}
		aluno2 := models.Aluno{Nome: "miguel"}
		aluno3 := models.Aluno{Nome: "janaina"}
		aluno4 := models.Aluno{Nome: "cicero"}
		aluno5 := models.Aluno{Nome: "dilan"}
		database.Create(&aluno1)
		database.Create(&aluno2)
		database.Create(&aluno3)
		database.Create(&aluno4)
		database.Create(&aluno5)

		// Disciplinas
		disc1 := models.Disciplina{
			Nome:        "apis e microservicos",
			Publica:     false,
			Professores: []models.Professor{prof1},
			Alunos:      []models.Aluno{aluno1, aluno2, aluno3, aluno4},
		}
		disc2 := models.Disciplina{
			Nome:        "matematica financeira",
			Publica:     true,
			Professores: []models.Professor{prof3},
			Alunos:      []models.Aluno{aluno2},
		}
		disc3 := models.Disciplina{
			Nome:        "matematica basica",
			Publica:     false,
			Professores: []models.Professor{prof2, prof3},
			Alunos:      []models.Aluno{aluno1, aluno2},
		}
		database.Create(&disc1)
		database.Create(&disc2)
		database.Create(&disc3)
	}
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
