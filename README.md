# gocommerce

Simple Go backend for the svelte ecommerce frontend

## Installation

Run the code directly if you have go installed on the machine

i. Clone Repository and change directory into the folder

```bash
$ git clone https://github.com/wesleymutwiri/authvice.git && cd authvice
```

ii. Build the application into a single binary via:

```bash
$ go build -o main .
```

iii. Create a sample .env file with your postgres configuration(Dockerfile coming soon to save you the trouble)

```bash
$ cp .env.example .env
```

iv. Run the executable go file by:

```bash
$ ./main
```

# REST API

The REST API to the gocommerce app is described below.

## Get list of Users

### Request

`GET /users`

    curl -i -H 'Accept: application/json' http://localhost:8080/users/

### Response

    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    Content-Length: 2

    []

## Create a new User

### Request

`POST /users`

    curl -i -H 'Accept: application/json' --data '{"username": "aa","email": "aa@gmail.com","password": "password"}' http://localhost:8080/users

### Response

    HTTP/1.1 201 Created
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 201 Created
    Connection: close
    Content-Type: application/json
    Location: /thing/1
    Content-Length: 36

    {
        "id":1,"username":"aa",
        "email":"aa@gmail.com",
        "created_at":"2020-11-10T20:57:31.763849+03:00",
        "updated_at":"2020-11-10T20:57:31.763849+03:00"
    }

## Get a specific User

### Request

`GET /users/id`

    curl -i -H 'Accept: application/json' http://localhost:8080/users/1

### Response

    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    Content-Length: 36

    {
        "id":1,
        "username":"wes",
        "email":"wes@gmail.com",
        "created_at":"2020-11-10T20:57:31.763849+03:00",
        "updated_at":"2020-11-10T20:57:31.763849+03:00"
    }

## Get a non-existent User

### Request

`GET /users/id`

    curl -i -H 'Accept: application/json' http://localhost:8080/users/9999

### Response

    HTTP/1.1 404 Not Found
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 404 Not Found
    Connection: close
    Content-Type: application/json
    Content-Length: 35

    {"error":"record not found"}

## Get list of Products

### Request

`GET /products`

    curl -i -H 'Accept: application/json' http://localhost:8080/products/

### Response

    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    Content-Length: 2

    []

## Create a new Product

### Request

`POST /products`

    curl -i -H 'Accept: application/json' --data '{"name": "socks","description": "Some socks over here that do well","regular_price": 100,"discount_price": 30"quantity": 2,"number_items": 10}' http://localhost:8080/products

### Response

    HTTP/1.1 201 Created
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 201 Created
    Connection: close
    Content-Type: application/json
    Location: /thing/1
    Content-Length: 36


       {
        "id":1,
        "name":"socks",
        "description":"Some socks over here that do well",
        "regular_price":100,
        "discount_price":30,
        "quantity":2,
        "number_items":10,
        "created_at":"2021-01-03T22:45:29.737201+03:00",
        "updated_at":"2021-01-03T22:45:29.737201+03:00"
        }

## Get a specific Product

### Request

`GET /products/id`

    curl -i -H 'Accept: application/json' http://localhost:8080/products/1

### Response

    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    Content-Length: 36

    {
        "id":1,
        "name":"socks",
        "description":"Some socks over here that do well",
        "regular_price":100,
        "discount_price":30,
        "quantity":2,
        "number_items":10,
        "created_at":"2021-01-03T22:45:29.737201+03:00",
        "updated_at":"2021-01-03T22:45:29.737201+03:00"
    }

## Get a non-existent Product

### Request

`GET /products/id`

    curl -i -H 'Accept: application/json' http://localhost:8080/products/9999

### Response

    HTTP/1.1 404 Not Found
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 404 Not Found
    Connection: close
    Content-Type: application/json
    Content-Length: 35

    {"error":"record not found"}
