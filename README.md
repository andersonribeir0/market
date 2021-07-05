# Market
Aplicação responsável pela consolidação e gestão dos dados da feiras em São Paulo.

# Tecnologias usadas
- docker
- golang
- dynamodb

# Requisitos para executar a aplicação
- docker
- docker-compose
- make
- cURL (opcional)

# Executando a aplicação localmente
1. Na raiz do repositório, execute o comando `make build` para iniciar a geração da imagem da aplicação.
2. Após a geração da imagem "market_app", agora podemos criar os containers executando `make run`.
3. Agora deve ser possível executar um simples comando cURL para verificar se o container está rodando normalmente:
`curl -X GET "localhost:8080/health" | json_pp -json_opt pretty,canonical`
4. Agora é preciso criar a tabela e popular a base de dados com o arquivo csv que se encontra em seu respectivo pacote.
Para isso, execute: `make createTable`. Esse comando inicia a criação da tabela e pode levar alguns poucos segundos para ser executado.
5. Uma vez criada, a tabela deve ser preenchida com os dados do arquivo csv. Execute o comando `make importCsv`.
6. Feita a importação podemos novamente executar um comando cURL para validar a importação feita com sucesso: `curl -X GET "localhost:8080/v1/market/1" | json_pp -json_opt pretty,canonical`

**Caso não possua cURL, você pode também utilizar um outro cliente como o Postman.*


# Endpoints

### Get
- `/health`
- `/v1/market/:id`
- `/v1/market?codDist=:codDist`

### Delete
- `/v1/market/:id`

### Put
- `/v1/market` 

# Testes
Para executar os testes unitários e de integração. Basta executar o comando `make test`. Ele irá gerar um relatório em dois aquivos, um html `coverage.html` e outro `coverage.out` com a cobertura de testes por pacote.
A cobertura de testes é realizada de maneira independente pacote a pacote.

# Logs
Todas os logs das chamadas http são enriquecidas com informação de requestId com o objetivo de melhorar o tracing.
Logs estão estruturados e possuem duas saídas: stdout (/dev/stdout) e para o arquivo `app_logs.log` localizado dentro do projeto. Como o volume do container da aplicação está mapeado para raiz do projeto, um arquivo de log será criado localmente a medida que a aplicação recebe requisições.

# Comandos

- `make clear` exclui todos os containers
- `make stop` derruba todos os containers
- `make build` executa o processo de construção da imagem da aplicação
- `make dynamodb` levanta um container baseado na imagem existente do dynamo no dockerhub
- `make createTable` cria tabela no dynamodb
- `make deleteTable` exclui a tabela no dynamodb
- `make importCsv` popula o dynamodb com os dados do csv
- `make app` levanta um container da aplicação baseada na imagem gerada pelo comando `make build`
- `make test` executa os testar unitários e de integração
