# Service - Services

## Introduction

Services service is responsible for providing the domain of *Organization Services*.

## Contents

* [Development](#development)
    * [Prerequisites](#prerequisites)
    * [Make Commands](#make-commands)
    * [Running the Service Locally](#running-the-service-locally)
    * [Project Layout](#project-layout)
        * [Layers and Folder Structure](#layers-and-folder-structure)
* [Deployment](#deployment)
    * [Dependencies](#dependencies)
        * [Infrastructure Dependencies](#infrastructure-dependencies)
        * [Service Dependencies](#service-dependencies)
        * [Environment Variables](#environment-variables)

    
## Development

### Prerequisites

The following table lists _hard_ dependencies you will need to use this project.

| Name                                                       | Version  | Notes                                    |
|------------------------------------------------------------|----------|------------------------------------------|
| [Go](https://golang.org/doc/)                              | 1.17+    | Required to build and spin up service    |
| [docker](https://www.docker.com/products/docker-desktop)   | 18.02.0+ | Required to build and spin up service    |
| [docker-compose](https://docs.docker.com/compose/install/) | 1.20.0+  | Required to build and spin up service    |


### Make Commands

| Command  | Description                                                                                                                       |
|----------|-----------------------------------------------------------------------------------------------------------------------------------|
| watch    | Starts the service and all necessary dependencies in the foreground including pulling postgress docker image and doing migrations |
| run      | Starts the service                                                                                                                |
| migrate  | Starts migrations  and creates tables                                                                                             |
| rollback | Starts rolling back the migration and drops tables.                                                                               |
| test     | starts running unit tests                                                                                                         |



### Running the Service Locally

1. Run ```make watch``` from project root it starts the service on port 8080.
2. Run ```make test```  from project root it starts the unit tests.


### Project Layout

This project roughly followed the layout of Go projects as described at
[https://github.com/golang-standards/project-layout](https://github.com/golang-standards/project-layout).

| Directory     | Description                                                                                    |
|---------------|------------------------------------------------------------------------------------------------|
| `cmd/`        | This Go package is where `main` is used for the executables of the project                     |
| `internal/`   | Application specific Go packages, e.g., they cannot be shared and are specific to this service |
| `migrations/` | Any files relating to migration                                                                |
| `tests/`      | tests for the service are located in here.                                                     |

#### Layers and Folder Structure

There are 3 main layers in this repo, including `Controller`, `ServiceHandler`, and `Repository`. The only way for these layers
to interact with each other should be through their interfaces. The lower layers do not have any knowledge about
the upper layers.

The `entity` is the entities that represents the model in the database.

`internal/controller/helper` contains the models for every request and response.
The `helper` should not be used in `Service` or `Repository`.


#### Infrastructure Dependencies

Services Service depends on Postgres, it is connected to Postgres Database.

### Environment Variables

The following environment variables are [defined in Services Service](./.env), and can be used to
influence behaviour.

| Name                                    | Description                                                                                                                                                    |
|-----------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `DB_HOST`                               | Data Base host                                                                                                                                                 |
| `DB_PORT`                               | data base port that application connects to                                                                                                                    |
| `DB_USER`                               | user name that connect to db and makes connection to it                                                                                                        |
| `DB_PASSWORD`                           | password to connect to the DB                                                                                                                                  |
| `DB_NAME`                               | database name                                                                                                                                                  |

