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
- Docker (para conteinerização)

## Como rodar localmente

1. Instale as dependências:
   ```sh
   go mod tidy


## Inicie cada serviço em terminais separados:

cd pessoa_service
go run .

# Em outro terminal:
cd atividade_service
go run .

# Como rodar com Docker
Certifique-se de ter o Docker e o Docker Compose instalados.
Na raiz do projeto, rode
docker-compose up --build

# Os serviços ficarão disponíveis em:
pessoa_service: http://localhost:8081
atividade_service: http://localhost:8080

# Estrutura de pastas

atividade_service/
  controllers/
  models/
  routes.go
  config.go
  app.go
  Dockerfile

pessoa_service/
  controllers/
  models/
  routes.go
  config.go
  app.go
  Dockerfile

# Exemplos de uso:

# Professor

Listar professores

GET http://localhost:8081/professores

Criar Professor

POST http://localhost:8081/professores
Body: { "nome": "Novo Professor" }

Atualizar Professor

PUT http://localhost:8081/professores/1
Body: { "nome": "Professor Atualizado" }

Deletar Professor

DELETE http://localhost:8081/professores/1

# Alunos 

Listar Alunos 

GET http://localhost:8081/alunos

Criar Aluno

POST http://localhost:8081/alunos
Body: { "nome": "Novo Aluno" }

Obter Aluno por ID

GET http://localhost:8081/alunos/1

Atualizar Aluno

PUT http://localhost:8081/alunos/1
Body: { "nome": "Aluno Atualizado" }

Deletar Aluno

DELETE http://localhost:8081/alunos/1

# Verificação se o professor leciona a disciplina

GET http://localhost:8081/leciona/1/1