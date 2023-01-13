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

## Deployment
This project is deployed to Google Cloud Run. The action can be viewed in the directory `.github/workflows/deployment`. The URL for the deployed application is: https://btc-billionaire-4w6hwgpraa-uw.a.run.app

This repository is using GitHub actions to automate the deployment process, You can manually trigger the deployment process by using the `workflow_dispatch` event in GitHub Actions. To do so, you will need to go to the Actions tab of your repository on GitHub, find the deployment workflow, and then click the "Run workflow" button and select "workflow_dispatch" from the dropdown menu. This will manually start the deployment process and deploy the latest version of the code.
