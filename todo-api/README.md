## Estrutura do Projeto


```bash
todo-api/
│
├── main.go
├── domain/
│   └── todo.go
├── repository/
│   └── todo_repository.go
├── handler/
│   └── todo_handler.go
└── postgres/
    ├── Dockerfile
    └── init.sql
```

## Configuração do Projeto

Instale as dependências necessárias:

```bash
go get -u github.com/gorilla/mux
go get -u github.com/lib/pq
```

Para iniciar o banco de dados, execute no root da aplicação:

```bash
docker build -t todo-postgres ./postgres
docker run -d --name todo-db -p 5432:5432 todo-postgres
```

## Testando a API

1. Certifique-se de que o banco de dados está rodando:

   ```bash
   docker start todo-db
   ```

2. Execute a API:

   ```bash
   go run main.go
   ```

Use uma ferramenta como cURL ou Postman para testar os endpoints:

**Criar um Todo**:

```bash
curl -X POST http://localhost:8080/todos -H "Content-Type: application/json" -d '{"title":"Learn Go","description":"Study Go programming"}'
```

**Obter todos os Todos**:

```bash
curl http://localhost:8080/todos
```

**Obter um Todo específico**:

```bash
curl http://localhost:8080/todos/1
```

**Atualizar um Todo**:

```bash
curl -X PUT http://localhost:8080/todos/1 -H "Content-Type: application/json" -d '{"title":"Learn Go","description":"Study Go programming","completed":true}'
```

**Deletar um Todo**:

```bash
curl -X DELETE http://localhost:8080/todos/1
```
