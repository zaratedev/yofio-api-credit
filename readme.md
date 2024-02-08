# YoFio API Credit Asigment

Prueba técnica para asignación de creditos para YoFio

Descripción del problema:

Tenemos 3 montos de créditos que damos a nuestros clientes ($300, $500 y $700). Cuando llega el dinero de inversión, queremos determinar cuántos créditos de cada monto podríamos asignar con ese dinero, sin que nos sobre 1 peso. Tu trabajo para este ejercicio es ayudarnos a calcular las posibles cantidades de créditos de $300, $500 y $700 que podemos otorgar con el total de la inversión. Si existe más de una opción podrías seleccionar cualquiera de ellas.

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

Ejemplo Request

```curl
curl  -X POST \
  'http://localhost:8000/credit-assigment' \
  --header 'Accept: */*' \
  --header 'Content-Type: application/json' \
  --data-raw '{"investment": 3000}'
```

Respuesta


```bash
POST - http://{{host}}/statistics
```

Ejemplo Request

```curl
curl  -X POST \
  'http://localhost:8000/statistics' \
  --header 'Accept: */*' \
  --header 'Content-Type: application/json' \
```

Respuesta

## Deploy

PROD: Se usa [Serverless Framework](https://www.serverless.com/framework/docs/getting-started) para el deploy de la API, se require tener una cuenta de AWS.

```bash
make deploy
```

### Unit Test
```bash
go test ./... -v
```