from flask import Blueprint
from controller.resposta_controller import *

resposta = Blueprint('resposta', __name__)

@resposta.route('/<int:id>', methods=["GET"])
def get_respostas_por_aluno(id):
    return getRespostas_por_aluno(id)

@resposta.route('/', methods=["POST"])
def create_Resposta():
    return createResposta()

@resposta.route('/<int:id>', methods=["PUT"])
def update_Resposta(id):
    return updateResposta(id)

@resposta.route('/<int:id>', methods=["DELETE"])
def delete_Resposta(id):
    return deleteResposta(id)

@resposta.route('/validar_aluno/<int:aluno_id>', methods=["GET"])
def validar_Aluno(aluno_id):
    return validarAluno(aluno_id)