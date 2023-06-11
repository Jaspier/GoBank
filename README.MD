# GO BANK

A simple banking application back-end built with Golang and PostgreSQL.

Learning about how to design, develop, and deploy a complete backend system from scratch using Golang, PostgreSQL, and Docker.

Built with Test-Driven Development (TDD).

> Project is currently put on hold.

## Go bank service

This service is for a simple banking app. It will provide APIs for the frontend to do following things:

1. Create and manage bank accounts, which are composed of owner’s name, balance, and currency.
2. Record all balance changes to each of the account. So every time some money is added to or subtracted from the account, an account entry record will be created.
3. Perform a money transfer between 2 accounts. This should happen within a transaction, so that either both accounts’ balance are updated successfully or none of them are.

## Setup local development

### Install tools

- [Docker desktop](https://www.docker.com/products/docker-desktop)
- [TablePlus](https://tableplus.com/)
- [Golang](https://golang.org/)
- [Homebrew](https://brew.sh/)
- [Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
  ```bash
  brew install golang-migrate
  ```
- [Sqlc](https://github.com/kyleconroy/sqlc#installation)
  ```bash
  brew install sqlc
  ```
- [Gomock](https://github.com/golang/mock)

  ```bash
  go install github.com/golang/mock/mockgen@v1.6.0
  ```

### Setup infrastructure

- Start postgres container:
  ```bash
  make postgres
  ```
- Create go_bank database:
  ```bash
  make createdb
  ```
- Run db migration:
  ```bash
  make migrateup
  ```

## How to generate code

### How to generate code

- Generate SQL CRUD with sqlc:

  ```bash
  make sqlc
  ```

- Generate DB mock with gomock:

  ````bash
  ```bash
  make mock
  ```
  ````

### How to run

- Run server:
  ```bash
  make server
  ```
- Run test:
  ```bash
  make test
  ```
