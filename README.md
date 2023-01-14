# BTC-Billionaire

<p align="center">
    <a href="https://goreportcard.com/report/github.com/mhdiiilham/BTC-Billionaire" target="_blank"><img src="https://goreportcard.com/badge/github.com/mhdiiilham/BTC-Billionaire" /></a>
</p>

<p align="center">
    <img src="https://github.com/mhdiiilham/BTC-Billionaire/actions/workflows/ci.yml/badge.svg" />
    <img src="https://github.com/mhdiiilham/BTC-Billionaire/actions/workflows/migrations.yml/badge.svg" />
    <img src="https://github.com/mhdiiilham/BTC-Billionaire/actions/workflows/deployment.yaml/badge.svg" />
</p>

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

## Running on Local Machine
1. Create `app.env` based on the `sample.env` file. This file should contain all the necessary environment variables for the application to run properly.

2. Run the unit tests by executing the command `make test`. This will ensure that all the code is functioning as expected before running the server.

3. Spin up the database using Docker based on the `docker-compose.yml` file in the scripts directory. This file contains all the necessary configurations to run a local instance of the database.
```shell
make dependencies
# or
docker-compose -f scripts/docker-compose.yml up
```
4. Run the server using the command `make run`. This will start the server and make it available on your specified port on localhost.

<em>It is important to note that, you may need to modify the docker-compose.yml file to match with the database credentials and endpoint you are using. Also, make sure you have docker and docker-compose installed on your machine.</em>

## API Documentation
The BTC-Billionaire server exposes a simple RESTful API for creating and retrieving Bitcoin transactions.

You can find a Postman exported collection on the project root directory with the name `BTC Billionaire.postman_collection.json` You can import this collection into Postman to easily test the API endpoints and see examples of the expected request and response formats.

You can also visit this [Postman](https://documenter.getpostman.com/view/9584176/2s8ZDSbk2B) URL to view the API Documentation.

## Deployment
This project is deployed to Google Cloud Run. The action can be viewed in the directory `.github/workflows/deployment`. The URL for the deployed application is: https://btc-billionaire-4w6hwgpraa-uw.a.run.app

This repository is using GitHub actions to automate the deployment process, You can manually trigger the deployment process by using the `workflow_dispatch` event in GitHub Actions. To do so, you will need to go to the Actions tab of your repository on GitHub, find the deployment workflow, and then click the "Run workflow" button and select "workflow_dispatch" from the dropdown menu. This will manually start the deployment process and deploy the latest version of the code.
