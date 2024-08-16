### Para iniciar o banco de dados, execute:

```bash
// todo-api/postgres
docker build -t todo-postgres ./postgres
docker run -d --name todo-db -p 5432:5432 todo-postgres
```
