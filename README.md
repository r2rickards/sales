<h1 align="center">💫 Go MongoDB API</h1>

## 🎯 What is that

API made in Go using MongoDB and JWT.

## 🔧 Features

- Not many libs & clean code
- JWT & MongoDB
- 3 routes (`GET` /token, `GET` (With token protection) /articles, `POST` (With token protection) /articles)

## 🚀 Run the project

1. Update the configuration in config.yml

```yml
server:
  host: 127.0.0.1
  port: 8080

database:
  host: 127.0.0.1
  port: 27017
  name: test
  user: test
  password: test
  ssl: false

jwt:
  secret: JwtSecret1234!
```

2. Run

```bash
go run main.go
```

3. Enjoy

By default, the API is available at the following address: http://127.0.0.1:8080 or http://localhost:8080
