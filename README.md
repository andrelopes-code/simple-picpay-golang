
#  Desafio Backend PicPay


esta é uma implementação em golang do desafio [Desafio Back-end PicPay](https://github.com/PicPay/picpay-desafio-backend).

### Local
Para rodar o app localmente é necessario um arquivo `.env` com as variaveis:
```bash
DB_PATH=./data  # Diretorio do banco de dados sqlite
DB_NAME=data.db  # Nome do banco de dados sqlite
NOTIFICATION_URL=https://run.mocky.io/v3/3157459b-602d-4ceb-bdaa-6e3ea5115ee6
AUTH_URL=https://util.devi.tools/api/v2/authorize
```
```bash
# Utiliza o makefile para rodar o app com as variaveis do .env
make run
```

### Docker
Ou então rodar o app utilizando o docker com o docker compose:
```bash
docker compose up --build
```

O banco tem 5 usuários inicializados ao subir a aplicação, são eles:
```
 ID | NAME | EMAIL   | USER TYPE     | BALANCE
 1  | John | j@a.com | common user   | 530
 2  | Jane | j@b.com | merchant user | 5000
 3  | Mary | m@c.com | merchant user | 10000
 4  | Mark | m@d.com | common user   | 2800
 5  | Mike | m@e.com | common user   | 10
```

Com a api rodando, acesse o swagger em: `http://localhost:8080/swagger/index.html`