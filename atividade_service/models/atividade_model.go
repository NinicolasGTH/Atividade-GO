package models

import (
	"errors"

	"gorm.io/gorm"
)

// Ponteiro global para o banco de dados
var DB *gorm.DB

// Structs para o banco de dados
type Resposta struct {
	ID          uint   `gorm:"primaryKey"`
	IdAluno     uint   `json:"id_aluno"`
	Resposta    string `json:"resposta"`
	Nota        *int   `json:"nota"` // Nota pode ser nula
	AtividadeID uint
}

type Atividade struct {
	ID           uint       `gorm:"primaryKey" json:"id_atividade"`
	IdDisciplina uint       `json:"id_disciplina"`
	Enunciado    string     `json:"enunciado"`
	Respostas    []Resposta `json:"respostas" gorm:"foreignKey:AtividadeID"`
}

// Erro customizado
var ErrAtividadeNotFound = errors.New("Atividade não encontrada")

// Listar todas as atividades
func ListarAtividades() ([]Atividade, error) {
	var atividades []Atividade
	result := DB.Preload("Respostas").Find(&atividades)
	return atividades, result.Error
}

// Obter uma atividade por ID
func ObterAtividade(id uint) (Atividade, error) {
	var atividade Atividade
	result := DB.Preload("Respostas").First(&atividade, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return atividade, ErrAtividadeNotFound
	}
	return atividade, result.Error
}
