![badge](https://img.shields.io/badge/Status-InProgress-yellow.svg)

# 🧠 Go CRUD API Beginner Project

This repository is a hands-on learning project for developers transitioning from .NET (C#) to **Go (Golang)**.  
It contains basic exercises that demonstrate Go syntax, struct, slice, function, goroutine, channel, and an in-memory REST API.

---

## 📌 Project Goals

- Learn basic Go syntax and data structures
- Practice writing Go functions, structs, and methods
- Understand Go's concurrency model (goroutines & channels)
- Build a simple REST API without using external libraries
- Become comfortable transitioning from OOP (.NET) to minimalistic Go

---

## 📦 Project Structure
```text
go-exercise
├── main.go # Entry point
├── exercise1_hello.go # Print "Hello, Go!"
├── exercise2_variable.go # Variable declaration and output
├── exercise3_add.go # Function for addition
├── exercise4_if_switch.go # Conditionals and switch
├── exercise5_loop.go # Looping 1 to 10
├── exercise6_struct.go # Person struct & slice
├── exercise7_method.go # Method on struct
├── exercise8_crud.go # In-memory CRUD
├── exercise9_goroutine.go # Basic goroutines
├── exercise10_channel.go # Channels for communication
├── webserver.go # Simple REST API with net/http
└── README.md # You are here
```

---

## 🧪 Exercises Overview

| Exercise | Description |
|----------|-------------|
| **1** | Print `"Hello, Go!"` to the console |
| **2** | Declare and print variables |
| **3** | Create function to add two numbers |
| **4** | Use `if`, `else`, and `switch` conditions |
| **5** | Use `for` loop to print numbers |
| **6** | Create a `Person` struct and store in a slice |
| **7** | Add method to `Person` struct (e.g. `.Greet()`) |
| **8** | Implement basic CRUD operations in slice |
| **9** | Run concurrent goroutines to print messages |
| **10** | Use channels to send messages between goroutines |
| **11** | Serve a web API using `net/http` |
| **12** | Implement REST API: `GET`, `POST`, `DELETE` with in-memory data |

---

## 🌐 REST API Endpoints

In `webserver.go`

| Method | Endpoint       | Description              |
|--------|----------------|--------------------------|
| GET    | `/books`       | List all books           |
| POST   | `/books`       | Add new book             |
| GET    | `/books/{id}`  | Get book by ID           |
| DELETE | `/books/{id}`  | Delete book by ID        |

---

## 🚀 How to Run

> Prerequisite: [Install Go](https://go.dev/dl)

### Run specific exercise:

```bash
go run exercise3_add.go
