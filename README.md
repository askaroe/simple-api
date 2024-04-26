# Simple API

This is a simple API project written in Go, using the Gin web framework.

## Introduction

This project provides a basic RESTful API for managing user data. It includes endpoints for creating, retrieving, updating, and deleting user records.

## API Endpoints

- **POST /api/v1/users**: Create a new user.
- **GET /api/v1/users**: Get a list of all users.
- **GET /api/v1/users/:id**: Get a user by ID.
- **PUT /api/v1/users/:id**: Update a user's information.
- **DELETE /api/v1/users/:id**: Delete a user.

## Dependencies

- [Gin](https://github.com/gin-gonic/gin): HTTP web framework for Go.
- [GORM](https://gorm.io/): ORM library for Go.
