# Atividades-go

Projeto de microserviços em Go para gerenciamento de atividades, professores, alunos e disciplinas.

## Estrutura

- **atividade**: Serviço para CRUD de atividades.
- **resposta**: Serviço para CRUD de respostas.

## Tecnologias

- Python
- Flask
- SQLite (banco de dados)
- Docker (para conteinerização)

## Como rodar localmente

1. Instale as dependências:
pip install flask
pip install flask_sqlalchemy
pip install requests


## Inicie cada API em terminais separados:

API school system:
docker-compose up --build

# Em outro terminal a API atividade:
docker-compose up --build


# As APIs ficarão disponíveis em:
School System: http://api_turma:8081
atividade: http://atividade:8080

# Estrutura de pastas

atividade/
  cliente/
  controller/
  model/
  route/
  config.py
  app.py
  Dockerfile
  docker-compose.yml
  requirements.txt



# Exemplos de uso:

# Atividade

POST -  /atividade

json:


{

  "id_turma": 1,

  "enunciado": "faça um hello world",

}



GET -  /atividade Lista todas as atividades:

{
  "enunciado": "faça um hello world",

  "id": 1,

  "id_turma": 1,

  "respostas": [
   
  ]
}



POST -  /resposta/ Cria uma nova resposta.

json:


{

"id_atividade": 1,

"id_aluno": 1,

"resposta": "helloworld.py",

"nota": 7.0

}



GET -  /resposta/<aluno_id> Lista todas as respostas por aluno: 

{

  "enunciado": "faça um hello world",

  "id": 1,

  "id_turma": 1,

  "respostas": 
[
    {

      "id": 1,

      "id_aluno": 1,

      "nota": 7.0,

      "resposta": "helloworld.py"

    }
  ]
}


GET -  /atividade/validar_turma/<turma_id> Verifica se uma turma existe (via outra API).

Resposta:

{
  "valida": true
}


GET -  /resposta/validar_aluno/<aluno_id> Verifica se um aluno existe (via outra API).

Resposta:

{
  "valida": true
}
