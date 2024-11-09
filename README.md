## Desafio #03 - Clean Architecture

O desafio consiste em adicionar o usecase de listagem de pedidos (orders).

#### 🖥️ Detalhes Gerais:

O novo usecase deverá ser utilizado para expor as informações, na medida em que forem requisitadas:
- Em endpoint (GET /order)
- Em um novo service ListOrders com gRPC
- Em uma nova query ListOrders com GraphQL

> Como adicional, será também possível consultar um pedido específico, através de um endpoint (GET /order/:id), um service `ListOrderById` com gRPC e uma query `listOrder` com GraphQL.

#### 🗂️ Estrutura do Projeto
    .
    ├── cmd                # Entrypoints da aplicação
    │    └── ordersystem   
    │           ├── main.go       ### Entrypoint principal
    │           ├── wire.go       ### Injeção de dependências
    │           └── .env          ### Arquivo de parametrizações globais
    ├── configs            # helpers para configuração da aplicação (viper)
    ├── internal
    │    ├── domain        # Core da aplicação
    │    │      ├── repository    ### Interfaces de repositório
    │    │      └── entity        ### Entidades de domínio
    │    ├── application   # Implementações de casos de uso e utilitários
    │    │      └── usecase       ### Casos de uso da aplicação
    │    ├── infra         # Implementações de repositórios e conexões com serviços externos
    │    │      ├── database      ### Implementações de repositório
    │    │      ├── graph         ### Implementações e códigos gerados para a API GraphQL
    │    │      ├── grpc          ### Implementações e códigos gerados para a API gRPC
    │    │      └── web           ### Implementações e códigos gerados para a API Rest
    │    └── event         # Implementações de eventos e listeners
    ├── pkg                # Pacotes reutilizáveis utilizados na aplicação
    ├── init_db.sql        # Script de inicialização do banco de dados
    └── README.md

#### 🧭 Parametrização
A aplicação servidor possui um arquivo de configuração `cmd/ordersystem/.env` onde é possível definir os parâmetros de timeout e URL's das API's para busca das informações do endereço.

```
DB_DRIVER=mysql                 # Database driver
DB_HOST=localhost               # Database host (More database details in docker-compose)
DB_PORT=3306                    # Database port
DB_USER=root                    # Database user
DB_PASSWORD=root                # Database password
DB_NAME=fc_challenge            # Database name
WEB_SERVER_PORT=:8000           # Web server port
GRPC_SERVER_PORT=50051          # gRPC server port
GRAPHQL_SERVER_PORT=8080        # GraphQL server port
```

> 💡 Os recursos externos MySQL e RabbitMQ são executador por meio de imagens Docker. Caso necessário alterar, poderá ser necessário revisar as variáveis de ambiente no arquivo `.env`.

#### 🚀 Execução:
Antes de iniciar, é necessário instalar as dependências do projeto. Para isso, execute o comando abaixo:
```bash
$ go mod tidy
```

Para executar a aplicação, existem duas opções:

#### 1. Utilizando o `makefile`:
Para facilitar a execução da aplicação, todas as etapas necessárias foram adicionadas ao makefile. Para executar a aplicação, basta executar o comando abaixo:
```bash
$ make run
```

#### 2. Executando manualmente:
Caso a opção anterior falhe, é possível executar a aplicação manualmente, seguindo os passos abaixo:
```bash
$ docker-compose up -d
$ cd cmd/ordersystem
$ go run ./main.go ./wire_gen.go
```

> 💡 Os comandos acima poderão falhar caso alguma das portas utilizadas estejam em uso. Caso isso ocorra, será necessário alterar as portas no arquivo `.env` ou encerrar os processos que estão utilizando as portas (8000, 8080, 50051, 3306, 5672 e 15672).

