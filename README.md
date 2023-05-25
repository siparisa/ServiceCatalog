# Service - Services

## Introduction

`Services`  is responsible for providing the domain of *Organization's Services*.

## Contents

* [API Documentation](#API-Documentation)
    * [Installation]()
    * [Endpoints]()
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

## API Documentation

### Installation
  to start the service check `Prerequisites` of this file then run `make watch` from the root directory of the project and run below endpoints.

example : `http://localhost:8080/services/2` 

### Endpoints

all Endpoints are located in server.go to call related handlers.
this is the router group :`/services`

**1.Create a Service**
Endpoint: `POST /services` 

Request Body:
```
{
  "data": {
    "name": "my name",
    "description": "My Servic",
    "version": "1.0"
  }
}
```
Please note that the request body is in JSON format and includes the following parameters:

**name** (string): The name of the service.
**description** (string): The description of the service.
**version** (string): The version of the service. 

**2.Get Services**
Endpoint: `GET /services`

Description: Retrieve a list of services.

Query Parameters:

**name** (string, optional): Filter services by name.
**page** (integer, optional): The page number for pagination.
**limit** (integer, optional): The maximum number of services to return per page.
**description** (string, optional): Filter services by description.
Example Request:
```
GET /services?name=name1&page=1&limit=10&description=service1
```
Please note that all query parameters are optional. You can include any combination of the parameters in the request URL to filter and paginate the results.

**3.Get Service by ID**
Endpoint: `GET /services/:serviceID` 

Description: Retrieve a specific service by its ID.

Path Parameters:

serviceID (integer): The ID of the service to retrieve. 

Example Request:
```
GET /services/2
```
In this example, the serviceID is provided as part of the URL path to specify the ID of the desired service.

**4.Update Service by ID**
Endpoint: `PUT /services/:serviceID`

Description: Update a specific service by its ID.

Path Parameters:

serviceID (integer): The ID of the service to update. 

Body Parameters:

**data** (object, optional): The data to update for the service.
**name** (string, optional): The new name of the service.
**description** (string, optional): The new description of the service.
version (string, optional): The new version of the service.
Example Request:
_PUT /services/2_
```
{
  "data": {
    "name": "name2",
    "description": "My service2",
    "version": "2.0"
  }
}
```
In this example, the serviceID is provided as part of the URL path to specify the ID of the service to update. The body contains optional parameters (name, description, version) to update the service's information.

**5.Delete Service by ID**
Endpoint: `DELETE /services/:serviceID` 

Description: Delete a specific service by its ID.

Path Parameters:

**serviceID** (integer): The ID of the service to delete.
Example Request:
```
DELETE /services/2
```

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

## Deployment
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

