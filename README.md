# To-do List API

This is a simple to-do list API built with Go and the Gin framework.

## Features

*   Get all to-do items.
*   Create a new to-do item.
*   Update a to-do item.
*   Delete a to-do item.

## Getting Started

### Prerequisites

*   [Go](https://golang.org/) installed on your machine.

### Installation and Running

1.  Clone the repository:
    ```bash
    git clone https://github.com/kidus-yoseph1/To-do-list-api.git
    ```
2.  Navigate to the project directory:
    ```bash
    cd To-do-list-api
    ```
3.  Run the server:
    ```bash
    go run backend/main.go
    ```
The server will start on `localhost:8080`.

## API Endpoints

| Method  | URL                | Description              |
| ------- | ------------------ | ------------------------ |
| `GET`   | `/api/lists`       | Get all to-do items.     |
| `POST`  | `/api/lists`       | Create a new to-do item. |
| `PATCH` | `/api/update/:id`  | Update a to-do item.     |
| `DELETE`| `/api/tasks/:id`   | Delete a to-do item.     |
