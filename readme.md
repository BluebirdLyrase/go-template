# Golang Boilerplate

Production-ready backend template built with Go and Gin.  
Designed to accelerate development of scalable REST APIs and microservices using clean architecture principles.

---

## 🚀 Tech Stack

- Go
- Gin (HTTP framework)
- GORM (ORM)
- Clean Architecture
- Environment-based configuration
- Middleware & standardized error handling

---

## 📁 Project Structure

/cmd # Application entrypoints
/internal
/modules
    /[module-name]
        /repository # Data access layer
        /service # Business logic
        /handler # HTTP handlers
        /model # Domain & DTO models
        module.go # Module initialization / wiring
        routes.go # Route registration
/shared
    /database # DB initialization
    /config # Environment config
    /utils # Global/helper functions

## 🛠 Installation

Install nodemon globally (for development auto-reload):

npm install -g nodemon


---

## ▶️ Run Project

nodemon


---

## 🧪 Run Tests

go test -v ./internal/modules/...


---


## 📌 Purpose

This boilerplate provides a scalable and maintainable foundation for building backend services, APIs, and microservices with Go.
