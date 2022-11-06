# GoLang Gin Starter Project

# React Starter Project

Simple Gin Starter project with some industry standard tools setup

## Run Locally

### Clone Repository or Use Template

Clone this repository `git clone https://github.com/thomasmendez/go-template-simple-server.git` or click `'Use this template'` in the Github UI to use the current project as a template.

### Install Go

Install the latest version of Go [here](https://go.dev/doc/install)

Verify Go is installed by checking the Go version with `go version`

### Install Makefile

This project uses the `Makefile` in order to easily run sets of commands. The commands can also be 
runned invidually using the commands listed for the makefile command. 

## Install Dependencies

Run `make tidy` in order to run the Go commands needed to install releated dependencies properly.

## Build

To build the project for a production environment run `make build`

### Build and Deply with Docker

The microservice can be build and deployed with Docker with `make docker`
Stop the application with `ctrl+c` on the terminal

## Tests

### Unit Tests

To run the unit test for the project, run `make test`

## Libraries Used

### Gin

High performance web framework in Golang.

### Viper

Highly flexible configuration library.

### Testify

Assertion library. Can also be used for mocking and building testing suites.

## Tools Used

### Docker

[Docker](https://www.docker.com/) allows you to build and run this microservice

### Husky

[Husky](https://github.com/automation-co/husky) allows you to run scripts in the git lifecyle

This project is configured to run test before making a successful commit

Modify the `pre-commit` hook with commands you wish to run before committing
