import requests

url = "http://api_turma:5000"

def turma_existe(turma_id):
    try:
        resp = requests.get(f"{url}/turmas/{turma_id}")
        return resp.status_code == 200
    except requests.RequestException as e:
        print(f"Erro ao conectar com serviço de turmas: {e}")
        return False
    
def aluno_existe(aluno_id):
    try:
        resp = requests.get(f"{url}/alunos/{aluno_id}")
        return resp.status_code == 200
    except requests.RequestException as e:
        print(f"Erro ao conectar com serviço de turmas: {e}")
        return False