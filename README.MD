# Atividades-go

Projeto de microserviços em Go para gerenciamento de atividades, professores, alunos e disciplinas.

## Estrutura

- **atividade_service**: Serviço para CRUD de atividades.
- **pessoa_service**: Serviço para gerenciamento de professores, alunos e disciplinas.

## Tecnologias

- Go (Golang)
- GORM (ORM)
- SQLite (banco de dados)
- Gorilla Mux (roteador HTTP)
- Docker (opcional para conteinerização)

## Como rodar

1. Instale as dependências:
   ```sh
   go mod tidy
# Inicie cada serviço
   cd atividade_service
go run .
# Em outro terminal:
cd pessoa_service
go run .




## Estrutura de pastas
atividade_service/
  controllers/
  models/
  routes.go
  config.go
  app.go

pessoa_service/
  controllers/
  models/
  routes.go
  config.go
  app.go

  ## Observações 

Cada serviço possui seu próprio banco SQLite.
Comunicação entre serviços via HTTP.
Pronto para ser conteinerizado com Docker.