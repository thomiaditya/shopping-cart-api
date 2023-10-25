# Shopping Cart API

This is a RESTful API for a shopping cart application built with Golang, GORM, PostgreSQL, and Echo. It allows users to create and manage customers, products, carts, orders, and payments.

## Features

- CRUD operations for customers, products, carts, orders, and payments
- Authentication and authorization using JWT
- Database migrations and seeding using GORM
- Validation and error handling using Echo
- Configuration using environment variables
- Makefile for common tasks

## Installation

To run this project, you need to have Go, PostgreSQL, and Make installed on your system. You also need to set up the following environment variables in a `.env` file:

```
DB_HOST=your database host
DB_PORT=your database port
DB_USER=your database user
DB_PASSWORD=your database password
DB_NAME=your database name
JWT_SECRET=your JWT secret key
```

Then, you can clone this repository and run the following commands in the project root directory:

```
# Run database migrations
make migrate

# Run database seeding
make seed

# Run the server
make run
```

## Usage

The API exposes the following endpoints:

| Method | Path                    | Description                          |
| ------ | ----------------------- | ------------------------------------ |
| GET    | /customers              | Get all customers                    |
| GET    | /customers/:id          | Get a customer by ID                 |
| POST   | /customers              | Create a new customer                |
| PUT    | /customers/:id          | Update a customer by ID              |
| DELETE | /customers/:id          | Delete a customer by ID              |
| GET    | /products               | Get all products                     |
| GET    | /products/:id           | Get a product by ID                  |
| POST   | /products               | Create a new product                 |
| PUT    | /products/:id           | Update a product by ID               |
| DELETE | /products/:id           | Delete a product by ID               |
| GET    | /carts                  | Get all carts                        |
| GET    | /carts/:id              | Get a cart by ID                     |
| POST   | /carts                  | Create a new cart                    |
| PUT    | /carts/:id              | Update a cart by ID                  |
| DELETE | /carts/:id              | Delete a cart by ID                  |
| GET    | /orders                 | Get all orders                       |
| GET    | /orders/:id             | Get an order by ID                   |
| POST   | /orders                 | Create a new order                   |
| PUT    | /orders/:id             | Update an order by ID                |
| DELETE | /orders/:id             | Delete an order by ID                |
| GET    | /payments               | Get all payments                     |
| GET    | /payments/:id           | Get a payment by ID                  |
| POST   | /payments               | Create a new payment                 |
| PUT    | /payments/:id           | Update a payment by ID               |
| DELETE | /payments/:id           | Delete a payment by ID               |

You can use any HTTP client tool such as Postman or curl to test the API. For example, to get all customers, you can send a GET request to `http://localhost:8080/customers`.

## License

This project is licensed under the MIT License - see the [LICENSE] file for details.