# YoFio API Credit Asigment

Prueba técnica para asignación de creditos para YoFio

## Installation

```bash
go mod download
```

## Usage
Ejecutar el siguiente comando para levantar la API en local

```bash
make run
```

## Endpoints

```bash
POST - http://{{host}}/credit-assigment
```
## Deploy

PROD: Se usa [Serverless Framework](https://www.serverless.com/framework/docs/getting-started) para el deploy de la API, se require tener una cuenta de AWS.

```bash
make deploy
```

### Unit Test
```bash
go test ./... -v
```