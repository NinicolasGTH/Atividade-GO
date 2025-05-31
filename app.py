from config import app, db
from route.atividade_route import atividade
from route.resposta_route import resposta


app.register_blueprint(atividade, url_prefix="/atividade")
app.register_blueprint(resposta, url_prefix="/resposta")

with app.app_context():
    db.create_all()

if __name__ == '__main__':    
    app.run(
        host=app.config['HOST'],
        port=app.config['PORT'],
        debug=app.config['DEBUG']
    )