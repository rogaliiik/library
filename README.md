## Library REST API
#### An API with ability to operate books which is built by Clean Architecture principles.

### Stack
- Go 1.21
- PostgreSQL 16
- Docker-Compose

### Features
- Following the Clean Architecture Principles
- HTTP routing with <a href=" https://github.com/go-chi/chi">go-chi/chi</a> framework
- Work with PostgreSQL. Migration files generation. SQL queries.
- Registration and authentication. Working with JWT. Middleware.
- Graceful Shutdown
- Running project in Docker-Compose
- Linting project with <a href="https://github.com/golangci/golangci-lint">golangci-lint</a>
- Generated Swagger docs with <a href="https://github.com/swaggo/http-swagger">swaggo/http-swagger</a>

### Endpoints 
#### API provides authentication endpoints:

- `/v1/auth/sign-up` with "POST" method to create new user
- `/v1/auth/sign-in` with "POST" method to login user and create JWT

#### Endpoints working with JWT:
- `/v1/book` with "POST" method to create new book
- `/v1/book/{bookId}` with "GET" method to get book by ID
- `/v1/book` with "GET" method to get all books by for current user
- `/v1/book/{bookId}` with "PUT" method to update book by ID
- `/v1/book/{bookId}` with "DELETE" method to delete book by ID

### Deploy
#### Clone repo:
```
git clone https://github.com/rogaliiik/library.git
```

#### Run app in docker-compose:
```
make run
```

#### To get Swagger open page in browser: 
```
localhost:8080/swagger/
```
