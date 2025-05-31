from flask import jsonify
from model.resposta_model import *
from cliente.cliente import aluno_existe


def getRespostas_por_aluno(id):
    resposta = ListarRespostasPorAluno(id)
    if not resposta:
        return jsonify({"erro": "Respostas não encontradas"}), 404
    return jsonify([r.dici() for r in resposta])


def createResposta():
    dados = request.get_json()
    nova, erro = AdicionarRespostas(dados)
    if erro:
        return jsonify({"erro": erro}), 400
    return jsonify(nova.dici()), 201


def updateResposta(id):
    dados = request.get_json()
    atualizada, erro = AtualizarResposta(id, dados)
    if erro:
        return jsonify({"erro": erro}), 404
    return jsonify(atualizada.dici())


def deleteResposta(id):
    sucesso, erro = DeletarResposta(id)
    if not sucesso:
        return jsonify({"erro": erro}), 404
    return jsonify({"mensagem": "Resposta deletada com sucesso"})


def validarAluno(aluno_id):
    existe = aluno_existe(aluno_id)
    if not existe:
        return jsonify({"valida": existe}), 404
    return jsonify({"valida": existe})