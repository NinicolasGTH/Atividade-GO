package models

import (
	"errors"

	"gorm.io/gorm"
)

// Ponteiro global para o banco de dados
var DB *gorm.DB

// Structs para o banco de dados

type Professor struct {
	ID   uint   `gorm:"primaryKey" json:"id_professor"`
	Nome string `json:"nome"`
}

type Aluno struct {
	ID   uint   `gorm:"primaryKey" json:"id_aluno"`
	Nome string `json:"nome"`
}

type Disciplina struct {
	ID          uint        `gorm:"primaryKey" json:"id_disciplina"`
	Nome        string      `json:"nome"`
	Publica     bool        `json:"publica"`
	Professores []Professor `gorm:"many2many:disciplina_professores;" json:"professores"`
	Alunos      []Aluno     `gorm:"many2many:disciplina_alunos;" json:"alunos"`
}

// Erros customizados
var (
	ErrDisciplinaNotFound = errors.New("Disciplina não encontrada")
)

// Listar todos os professores
func ListarProfessores() ([]Professor, error) {
	var professores []Professor
	result := DB.Find(&professores)
	return professores, result.Error
}

// Listar todos os alunos
func ListarAlunos() ([]Aluno, error) {
	var alunos []Aluno
	result := DB.Find(&alunos)
	return alunos, result.Error
}

// Verifica se um professor leciona uma disciplina específica
func Leciona(idProfessor int, idDisciplina int) (bool, error) {
	var disciplina Disciplina
	result := DB.Preload("Professores").First(&disciplina, idDisciplina)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, ErrDisciplinaNotFound
	}
	for _, prof := range disciplina.Professores {
		if int(prof.ID) == idProfessor {
			return true, nil
		}
	}
	return false, nil
}
