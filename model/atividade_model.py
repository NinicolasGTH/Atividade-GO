from flask import request
from config import db
from cliente.cliente import turma_existe

class Atividade (db.Model):
    __tablename__ = "atividade"

    id = db.Column(db.Integer, primary_key=True)
    id_turma = db.Column(db.Integer, nullable=False)
    enunciado = db.Column(db.String(255), nullable=False)
    respostas = db.relationship('Resposta', backref='atividade', cascade='all, delete-orphan')

    def __init__(self, id_turma, enunciado, respostas=None):
        self.id_turma = id_turma
        self.enunciado = enunciado
        self.respostas = respostas or []

    def dici(self):
        return {
            "id": self.id,
            "id_turma": self.id_turma,
            "enunciado": self.enunciado,
            "respostas": [r.dici() for r in self.respostas]
        }
    
def ListarAtividades():
    return Atividade.query.all()

def ListarAtividadePorId(idAtividade):
    return Atividade.query.get(idAtividade)

def CriarAtividade(dados):
    id_turma = dados.get("id_turma")
    enunciado = dados.get("enunciado")
    if not id_turma or not enunciado:
        return None, "ID da Turma é obrigatório"    
    if not turma_existe(id_turma):
        return None, "Turma não encontrada"
    if not enunciado:
        return None, "Enunciado é obrigatório"

    nova_atividade = Atividade(id_turma=id_turma, enunciado=enunciado)
    db.session.add(nova_atividade)
    db.session.commit()
    return nova_atividade, None

def AtualizarAtividade(idAtividade, dados):
    atividade = Atividade.query.get(idAtividade)
    if not atividade:
        return None, "Atividade não encontrada"

    id_turma = dados.get("id_turma")
    if id_turma is not None:
        if not turma_existe(id_turma):
            return None, "Turma não encontrada"
        atividade.id_turma = id_turma

    enunciado = dados.get("enunciado")
    if enunciado is not None:
        if not enunciado:
            return None, "Enunciado não pode ser vazio"
        atividade.enunciado = enunciado

    db.session.commit()
    return atividade, None

def DeletarAtividade(idAtividade):
    atividade = Atividade.query.get(idAtividade)
    if not atividade:
        return False, "Atividade não encontrada"

    db.session.delete(atividade)
    db.session.commit()
    return True, None