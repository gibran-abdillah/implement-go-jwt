# Go-JWT Implementation

This repository contains a simple program showcasing the implementation of JSON Web Tokens (JWT) for secure authentication and authorization in a web application.

## Features
- User Registration: Register new users with unique usernames and secure passwords.
- User Login: Authenticate users and issue JWT tokens upon successful login.
- User Management: Edit user details, delete users, and retrieve user information.
- Token-based Authentication: Utilize JWT tokens for secure user authentication.


## Technologies Used
- Go (Golang): The backend is developed using Go, a powerful and efficient programming language.
-  Gin Web Framework: The Gin framework is used to build the API endpoints and manage HTTP requests.
- GORM: GORM is employed as an Object Relational Mapping (ORM) library for interacting with the database.
- JSON Web Tokens (JWT): JWTs are used for secure token-based user authentication and authorization.

## Getting Started

- Clone this repository
- Navigate to the project directory
- Install depencidies
```sh
go mod download
```

- Configure your database in ```.env```
```sh
cp .env_examples .env
nano .env
```
- run the server
```
go run main.go
```

## Requests example
Visit ```request_example.md``` for more details


