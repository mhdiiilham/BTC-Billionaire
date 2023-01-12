# BTC-Billionaire

Welcome to BTC-Billionaire, a Go-based web server for handling POST and GET requests for Bitcoin transactions. This repository contains all the necessary code and instructions for setting up and running your own BTC-Billionaire server.

## Setting up Migrations
1. Install `golang-migrate` by running the following command:
```shell
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

2. Export an environment variable for convenience:
```shell
export POSTGRESQL_URL='postgres://postgres:password@localhost:5432/example?sslmode=disable'
```

3. Run the migrations based on what you need:

- To run migration up: `make migrate-up`
- To create a new migration file: `make migrate-create`
