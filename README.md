# About The Project

Sample REST API application, written as part of an interview challenge.
The app implements a single end-point (`/products`) which serves a list of avilable products,
with optional filtering.

# Getting Started

## Prerequisites

- make
- Docker
- docker-compose

## Usage

### Starting

Run the following command to start the application.

```
make up
```

Once the command has completed, you should be able to reach the web application running on `localhost:8080`.
```
$ curl 'http://localhost:8080/products' | jq

{
  "products": []
}
```

### Importing test data

Initially the database is empty and therefore all the requests to `/products` are going to return empty result (see example above). Use the command below to load pre-defined set of test data into the database.

```
make fixtures
```

If you repeat the `curl` command now, you should be able to see some results being returned.

```
$ curl 'http://localhost:8080/products?category=boots' | jq

{
  "products": [
    {
      "sku": "000001",
      "name": "BV Lean leather ankle boots",
      "category": "boots",
      "price": {
        "original": 89000,
        "final": 80100,
        "discount_percentage": "10%",
        "currency": "EUR"
      }
    },
    {
      "sku": "000004",
      "name": "Newer BV Lean leather ankle boots",
      "category": "boots",
      "price": {
        "original": 99000,
        "final": 99000,
        "discount_percentage": "0%",
        "currency": "EUR"
      }
    }
  ]
}
```

### Stopping

In order to stop the application, run the following command.
```
make down
```

This stops Docker containers w/o performing any clean-up.

### Cleaning up

In order to perform full clean-up, including removal of images & volumes, run:

```
make clean
```

### Running tests

A test suite, coming with the application can be run using the following command:
```
make test
```

This runs the tests inside a Docker container, to avoid dependency to the `Go` compiler.