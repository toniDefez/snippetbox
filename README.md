# 📘 Snippetbox

A simple web application for creating and viewing snippets, built with Go — following the structure and concepts from the book [Let’s Go](https://lets-go.alexedwards.net/).

## 🚀 Features

- Built with the Go programming language
- Create and view text snippets via HTTP
- MySQL database for persistent storage
- Clean separation of concerns (handlers, models, templates)
- Organized project structure suitable for testing and scaling
- Docker support for database containerization

## 📁 Project Structure

```
.
├── cmd/web             # Main web application entry point
│   ├── main.go         # Application startup logic
│   ├── handlers.go     # HTTP handlers
│   └── routes.go       # Route definitions
├── internal/models     # Database models
│   └── snippets.go     # SnippetModel with DB methods
├── ui                  # HTML templates
├── docker              # Docker-related files
│   └── mysql/init.sql  # MySQL init script
├── go.mod              # Go module definition
└── README.md           # Project documentation
```

## 🐳 Running the App with Docker

This project uses Docker to spin up a MySQL instance.

### 1. Start MySQL with Docker

```bash
docker-compose up -d
```

> MySQL will be accessible at `localhost:3306` with the credentials set in `docker-compose.yml`.

### 2. Run the Go application

Make sure your DSN is correctly configured in `main.go` or passed via flags:

```go
myuser:mypassword@tcp(127.0.0.1:3306)/snippetboxDB?parseTime=true
```

Then run:

```bash
go run ./cmd/web
```

The server will start on [http://localhost:4000](http://localhost:4000)

## 📬 Example Request

```bash
curl.exe -i -L -X POST -d " " http://localhost:4000/snippet/create
```

You should receive a redirect to `/snippet/view/{id}` if everything is working.

## 🧪 Development Notes

- SQL queries are handled in the `SnippetModel` methods.
- Routes are wired in `routes.go` and tied to handler methods on the `application` struct.
- The `internal` directory is used for non-application-specific code and packages.

## 📌 Requirements

- Go 1.21+
- Docker (for MySQL setup)
- MySQL 8.0+ or compatible

## ✅ To Do

- [ ] Add form-based snippet creation
- [x] Render snippets with HTML templates
- [ ] Add validation
- [ ] Add sessions and authentication

## 📝 License

This project is for learning purposes, inspired by the *Let’s Go* book by [Alex Edwards](https://alexedwards.net/).
