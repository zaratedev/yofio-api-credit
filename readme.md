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

PROD

```bash
make deploy
```

### Unit Test
```bash
go test ./tests -v
```