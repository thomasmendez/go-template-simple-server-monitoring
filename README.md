# GoLang Gin Starter Project

Simple Gin Starter project with some industry standard tools setup

## Run Locally

### Clone Repository or Use Template

Clone this repository `git clone https://github.com/thomasmendez/go-template-simple-server.git` or click `'Use this template'` in the Github UI to use the current project as a template.

### Install Go

Install the latest version of Go [here](https://go.dev/doc/install)

Verify Go is installed by checking the Go version with `go version`

### Install Makefile

This project uses the `Makefile` in order to easily run sets of commands. The commands can also be run invidually using the commands listed for the makefile command

### Install Dependencies

Run `make tidy` in order to run the Go commands needed to install releated dependencies properly

### Setup Environment Variables

Create a `.env` file from the `.env.example` file and modify the environment variables as needed

## Build

To build the project run `make build`

### Build and Deply with Docker

*Note: Make sure [Docker](https://www.docker.com/) is running on your machine*

The microservice can be build and deployed with Docker with `make docker`

API documentation can be viewed at [http://localhost:8081/api/swagger](http://localhost:8081/api/swagger)

Application can stop running by pressing `ctrl+c` on the terminal

## Tests

### Unit Tests

To run the unit test for the project, run `make test`

## Libraries Used

### Gin

[Gin](https://github.com/gin-gonic/gin) is a high performance web framework in Golang

### Viper

[Viper](https://github.com/spf13/viper) is a highly flexible configuration library

### Zap

[Zap](https://github.com/uber-go/zap) fast, structured, leveled logging in Golang

### Testify

[Testify](https://github.com/stretchr/testify) is an assertion library. Can also be used for mocking and building testing suites

### Prometheus Client

[Prometheus Go Client](github.com/prometheus/client_golang) is the Go client library for Prometheus. It is used to provide application metrics

## Tools Used

### Swagger

[Swagger](https://swagger.io/) is used as API documentation

Can be modified in the `swagger/swagger.yaml` file

### Docker

[Docker](https://www.docker.com/) allows you to build and run this microservice

### Prometheus

[Prometheus](https://prometheus.io/) provides metrics for the application

#### Alertmanager

[Alertmanager](https://prometheus.io/docs/alerting/latest/alertmanager/) handles alerts sent by client applications such as the Prometheus server

### Grafana

[Grafana](https://grafana.com/oss/grafana/) allows you to query, visualize, and alert on and understand your metrics

### Husky

[Husky](https://github.com/automation-co/husky) allows you to run scripts in the git lifecyle

This project is configured to run test before making a successful commit

Modify the `pre-commit` hook with commands you wish to run before committing
