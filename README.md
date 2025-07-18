# 💬 GoFr Chat

GoFr Chat is a modular and scalable chat backend service built with Go. It follows a clean three-layer architecture to ensure separation of concerns and ease of maintenance. This project includes RESTful APIs, Swagger documentation, and a clear folder structure — making it ideal for extension and production use.

---

## 📁 Project Architecture

```
/cmd/server           → Application entry point  
/configs              → Configuration files (e.g., YAML)  
/internal/handler     → HTTP handlers and route controllers  
/internal/service     → Business logic layer  
/internal/store       → Data persistence and access layer  
/internal/modules     → Shared models and types  
/pkg/utils            → Utility/helper functions  
/docs                 → Auto-generated Swagger documentation  
```

---

## ✨ Getting Started

### ✅ Prerequisites

* Go 1.18+
* `swag` CLI installed
  Install via:

  ```bash
  go install github.com/swaggo/swag/cmd/swag@latest
  ```

---

### ⚙️ Setup

```bash
# Clone the repo
git clone https://github.com/talanayush/gofr-chat.git
cd gofr-chat

# Install dependencies
go mod tidy

# Generate Swagger docs
swag init

# Run the server
go run cmd/server/main.go
```

> The server will start on `http://localhost:8080`

---

## 🔗 Swagger Documentation

We use [Swaggo](https://github.com/swaggo/swag) for API documentation.

### 📖 Access the Swagger UI:

```
http://localhost:8080/swagger/index.html
```

### 🔄 To regenerate docs:

```bash
swag init
```

---

## 📌 API Endpoints

| Method | Endpoint     | Description              |
| ------ | ------------ | ------------------------ |
| GET    | `/health`    | Server health check      |
| POST   | `/chat/send` | (Planned) Send a message |
| GET    | `/chat/{id}` | (Planned) Get chat logs  |

---

## 🛠 Tech Stack

* **Language:** Go
* **Router:** Gorilla Mux
* **Docs:** Swagger (swaggo)
* **Architecture:** 3-Layer Modular Design

---

## 🤝 Contribution

This is the first phase of GoFr Chat setup. Contributions are welcome!

* Please review the architecture
* Raise issues for improvement
* Submit PRs with modular changes

