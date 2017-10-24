# Openshop

Online version of this web is hosted on https://openshop-service.herokuapp.com

## Requirements

- Go >= 1.9
- PostgreSQL server

## Installation

First, you have to download glide on your local.

```
curl https://glide.sh/get | sh
```
Now, install all dependencies using glide.
```
glide install
```
Finally, you can run the server:
```
go run main.go
```

## Configuration

This app receive configurations via environment variables. Supported configuration flags are:

| Variable        | Description                                | Available Options                               |
| --------------- |:------------------------------------------ | ----------------------------------------------- |
| PORT            | Application port binding                   |                                                 |
| SSL_MODE        | Database SSL connection (default: disable) | required \| verify-full \| verify-ca \| disable |
| DB_HOST         | Database host url (default: localhost)     |                                                 |
| DB_PORT         | Database connection port (default: 5432)   |                                                 |
| DB_DB           | Database name (default: openshop)          |                                                 |
| DB_USER         | Database login name (default: postgres)    |                                                 |
| DB_PASS         | Database login password                    |                                                 |

## Architectural Layer

```
| ----------------- |
|  Controller Layer |
| ----------------- |
|   Service Layer   |
| ----------------- |
|     DAO Layer     |
| ----------------- |
|      Database     |
| ----------------- |
```

Roughly, this app divides into 3 layers: controller, service, and DAO layer. Controller layer is responsible for handling request and response. This layer also doing fields type validation while parsing. Then, service layer is responsible for processing main domain validation and bussiness process. Finally, DAO layer is responsible for wrapping database access query and DML.