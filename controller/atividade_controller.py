from flask import jsonify
from model.atividade_model import *
from cliente.cliente import turma_existe


def getAtividades():
    atividades = ListarAtividades()
    return jsonify([a.dici() for a in atividades])


def getAtividadeid(id):
    atividade = ListarAtividadePorId(id)
    if not atividade:
        return jsonify({"erro": "Atividade não encontrada"}), 404
    return jsonify(atividade.dici())


def createAtividade():
    dados = request.get_json()
    nova, erro = CriarAtividade(dados)
    if erro:
        return jsonify({"erro": erro}), 400
    return jsonify(nova.dici()), 201


def updateAtividade(id):
    dados = request.get_json()
    atualizada, erro = AtualizarAtividade(id, dados)
    if erro:
        return jsonify({"erro": erro}), 404
    return jsonify(atualizada.dici())


def deleteAtividade(id):
    sucesso, erro = DeletarAtividade(id)
    if not sucesso:
        return jsonify({"erro": erro}), 404
    return jsonify({"mensagem": "Atividade deletada com sucesso"})


def validarTurma(turma_id):
    existe = turma_existe(turma_id)
    if not existe:
        return jsonify({"valida": existe}), 404
    return jsonify({"valida": existe})