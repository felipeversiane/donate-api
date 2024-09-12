# Donation Manager API
API responsible for creating and managing donations.

## Introduction

The API itself is for study purposes, and aims to be a API for donations managment, using Docker, AWS S3 and Efí Gateway to managment transactions.

## 🛠 Main libraries used

- **gin/gonic** - Gin is a fast, lightweight web framework for Go, ideal for building APIs. It offers routing, middleware, and error handling with high performance and simplicity.

- **pgxpool** - A package from the pgx library, pgxpool provides a connection pool for PostgreSQL databases. It allows efficient management of database connections and helps in handling multiple concurrent requests to the database with reduced latency and resource consumption.

- **github.com/google/uuid** - This library provides utilities for generating and working with UUIDs (Universally Unique Identifiers). It supports the creation of various types of UUIDs, including UUIDv1, UUIDv4, and UUIDv5, and is used to ensure unique identifiers in distributed systems and databases.

- **github.com/aws/aws-sdk-go** - This is a Go client library for interacting with AWS services, including S3, EC2, Lambda, and more. It provides a convenient and efficient way for Go applications to communicate with AWS resources by offering APIs for managing cloud services, storage, computing, and other AWS features.

For a complete list of dependencies see the file [go.mod](https://github.com/felipeversiane/donate-api/blob/main/go.mod).


## 🚀 Development build

To be able to run the project as a developer, follow the steps:

 - Clone the repository to your machine.
 - Remember to checkout the branch you are developing.

From there you can run the following command to run the project and test the application

```bash
  docker-compose up --build
```

That's it, the API is running, be happy!

## Suport

For support, please email me [felipeversiane09@gmail.com](mailto:felipeversiane09@gmail.com)