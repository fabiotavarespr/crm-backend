# CRM Backend - Udacity Golang Nanodegree Project

This project implements a basic CRUD in a REST-API.

Your context would be for a customer in a CRM project scenario

# Index
- [CRM Backend - Udacity Golang Nanodegree Project](#crm-backend---udacity-golang-nanodegree-project)
- [Index](#index)
- [Stack](#stack)
- [Execution](#execution)
  - [Start infrastructure](#start-infrastructure)
    - [Makefile](#makefile)
    - [Docker Compose](#docker-compose)
  - [Stop infrastructure](#stop-infrastructure)
    - [Makefile](#makefile-1)
    - [Docker Compose](#docker-compose-1)
  - [Start project](#start-project)
    - [Makefile](#makefile-2)
    - [Go command](#go-command)
- [API](#api)
- [Endpoint](#endpoint)
  - [Getting all customers](#getting-all-customers)
    - [Request example](#request-example)
    - [Body example](#body-example)
    - [Response Example](#response-example)
  - [Getting a single customer](#getting-a-single-customer)
    - [Request example](#request-example-1)
    - [Body example](#body-example-1)
    - [Response Example](#response-example-1)
  - [Creating a customer](#creating-a-customer)
    - [Request example](#request-example-2)
    - [Body example](#body-example-2)
    - [Response Example](#response-example-2)
  - [Updating a customer](#updating-a-customer)
    - [Request example](#request-example-3)
    - [Body example](#body-example-3)
    - [Response Example](#response-example-3)
  - [Deleting a customer](#deleting-a-customer)
    - [Request example](#request-example-4)
    - [Body example](#body-example-4)
    - [Response Example](#response-example-4)


# Stack

- [Golang](https://go.dev/)
- [Gin Web Framework](https://gin-gonic.com/)
- [MongoDB](https://www.mongodb.com/)

# Execution

First, you need to start your docker infrastructure with the following command:

## Start infrastructure

### Makefile
```sh
make docker-compose-up
```

### Docker Compose
```sh
docker-compose -f resources/docker-compose/docker-compose.yaml up -d
```

## Stop infrastructure

### Makefile
```sh
make docker-compose-down
```

### Docker Compose
```sh
docker-compose -f resources/docker-compose/docker-compose.yaml down --remove-orphans
```

## Start project
Once your infrastructure is up, start the API with the following command:

### Makefile
```sh
make run-api
```

### Go command
```sh
go run ./cmd/main.go
```


# API

| Method | Resource         |
|:------:|:-----------------|
| GET    | /health          |
| GET    | /customers       |
| GET    | /customers/{id}  |
| POST   | /customers       |
| PUT    | /customers/{id}  |
| DELETE | /customers/{id}  |

Import the [postman collection](/resources/postman/CRM-Backend.postman_collection.json) for more details.


# Endpoint

## Getting all customers

 - GET - /customers

### Request example

```sh
curl --location --request GET 'http://localhost:3000/customers'
```

### Body example
```sh
No body
```

### Response Example

```json
{
    "customers": [
        {
            "id": "636fec6fad4733dfb2c87aef",
            "name": "Marty Zemlak",
            "email": "martyzemlak@qgx.info",
            "role": "onwer",
            "phone": "+1 (870) 940-4987",
            "contacted": false,
            "created_at": "2022-11-12T18:56:47.017Z",
            "updated_at": "2022-11-12T18:56:47.017Z"
        }
    ]
}
```

## Getting a single customer

 - GET - /customers/{id}

### Request example

```sh
curl --location --request GET 'http://localhost:3000/customers/636fec6fad4733dfb2c87aef'
```

### Body example
```sh
No body
```

### Response Example

```json
{
    "id": "636fec6fad4733dfb2c87aef",
    "name": "Marty Zemlak",
    "email": "martyzemlak@qgx.info",
    "role": "onwer",
    "phone": "+1 (870) 940-4987",
    "contacted": false,
    "created_at": "2022-11-12T18:56:47.017Z",
    "updated_at": "2022-11-12T18:56:47.017Z"
}
```

## Creating a customer

 - POST - /customers/

### Request example

```sh
curl --location --request POST 'http://localhost:3000/customers' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Marty Zemlak",
    "role": "onwer",
    "email": "martyzemlak@qgx.info",
    "phone": "+1 (870) 940-4987",
    "contacted": false
}'
```

### Body example
```sh
{
    "name": "Marty Zemlak",
    "role": "onwer",
    "email": "martyzemlak@qgx.info",
    "phone": "+1 (870) 940-4987",
    "contacted": false
}
```

### Response Example

```json
{
    "id": "636fec6fad4733dfb2c87aef",
    "name": "Marty Zemlak",
    "email": "martyzemlak@qgx.info",
    "role": "onwer",
    "phone": "+1 (870) 940-4987",
    "contacted": false,
    "created_at": "2022-11-12T18:56:47.017Z",
    "updated_at": "2022-11-12T18:56:47.017Z"
}
```

## Updating a customer

 - PUT - /customers/{id}

### Request example

```sh
curl --location --request PUT 'http://localhost:3000/customers/636fec6fad4733dfb2c87aef' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Mr. Walker Moore VI",
    "role": "manager",
    "email": "walker.moore@ikb.biz",
    "phone": "(660) 365-7314",
    "contacted": true
}'
```

### Body example
```sh
{
    "name": "Mr. Walker Moore VI",
    "role": "manager",
    "email": "walker.moore@ikb.biz",
    "phone": "(660) 365-7314",
    "contacted": true
}
```

### Response Example

```json
{
    "id": "636fec6fad4733dfb2c87aef",
    "name": "Mr. Walker Moore VI",
    "email": "walker.moore@ikb.biz",
    "role": "manager",
    "phone": "(660) 365-7314",
    "contacted": true,
    "created_at": "2022-11-12T18:56:47.017Z",
    "updated_at": "2022-11-12T18:59:16.437Z"
}
```

## Deleting a customer

 - DELETE - /customers/{id}

### Request example

```sh
curl --location --request DELETE 'http://localhost:3000/customers/636fec6fad4733dfb2c87aef'
```

### Body example
```sh
No body
```

### Response Example

```json
No Content
```