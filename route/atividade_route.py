from flask import Blueprint
from controller.atividade_controller import *

atividade = Blueprint('atividade', __name__)


@atividade.route("/", methods=['GET'])
def listar_atividades():
    return getAtividades()

@atividade.route('/<int:id>', methods=['GET'])
def listar_atividade_por_id(id):
    return getAtividadeid(id)

@atividade.route('/', methods=['POST'])
def criar_atividade():
    return createAtividade()

@atividade.route('/<int:id>', methods=['PUT'])
def atualizar_atividade(id):
    return updateAtividade(id)

@atividade.route('/<int:id>', methods=['DELETE'])
def deletar_atividade(id):
    return deleteAtividade(id)

@atividade.route('/validar_turma/<int:turma_id>', methods=["GET"])
def validar_turma(turma_id):
    return validarTurma(turma_id)