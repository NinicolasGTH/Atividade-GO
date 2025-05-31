from flask import request
from config import db
from cliente.cliente import aluno_existe

class Resposta (db.Model):
    __tablename__ = "resposta"

    id = db.Column(db.Integer, primary_key=True)
    id_atividade = db.Column(db.Integer, db.ForeignKey('atividade.id'), nullable=False)
    id_aluno = db.Column(db.Integer, nullable=False)
    resposta = db.Column(db.String(255), nullable=False)
    nota = db.Column(db.Float, nullable=True)

    def __init__(self, id_atividade, id_aluno, resposta, nota=None):
        self.id_atividade = id_atividade
        self.id_aluno = id_aluno
        self.resposta = resposta
        self.nota = nota

    def dici(self):
        return {
            "id": self.id,
            "id_aluno": self.id_aluno,
            "resposta": self.resposta,
            "nota": self.nota
        }
    
def ListarRespostasPorAluno(idAluno):
    return Resposta.query.filter_by(id_aluno=idAluno).all()
    
def AdicionarRespostas(dados):
    id_atividade = dados.get("id_atividade")
    id_aluno = dados.get("id_aluno")
    resposta = dados.get("resposta")
    nota = dados.get("nota", "")

    if not id_atividade:
        return None, "Atividade é obrigatório"
    if not id_aluno:
        return None, "Aluno é obrigatório"
    if not aluno_existe(id_aluno):
        return None, "Aluno não encontrado"
    if not resposta:
        return None, "Respostas são obrigatórias"

    nova_resposta = Resposta(id_atividade, id_aluno, resposta, nota)
    db.session.add(nova_resposta)
    db.session.commit()
    return nova_resposta, None

def AtualizarResposta(idResposta, dados):
    resposta = Resposta.query.get(idResposta)
    if not resposta:
        return None, "Resposta não encontrada"
    
    id_aluno = dados.get("id_aluno")
    if id_aluno is not None:
        if not aluno_existe(id_aluno):
            return None, "Aluno não encontrado"
        resposta.id_aluno = id_aluno

    id_atividade = dados.get("id_atividade")
    if id_atividade is not None:
        resposta.id_atividade = id_atividade

    resposta_texto = dados.get("resposta")
    if resposta_texto is not None:
        if not resposta_texto:
            return None, "Resposta não pode ser vazia"
        resposta.resposta = resposta_texto

    if "nota" in dados:
        resposta.nota = dados.get("nota")

    db.session.commit()
    return resposta, None

def DeletarResposta(idResposta):
    resposta = Resposta.query.get(idResposta)
    if not resposta:
        return False, "Resposta não encontrada"

    db.session.delete(resposta)
    db.session.commit()
    return True, None