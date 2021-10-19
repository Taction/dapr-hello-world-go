# introduction 

This is a go demo ready to convert business functions into wasm.

# Hello World

This tutorial will demonstrate how to instrument your application with Dapr, and run it locally on your machine.
You will deploying `order` applications with the flow identical to [Hello World](https://github.com/dapr/quickstarts/tree/master/hello-world).

The application invokes Dapr API via Dapr client, which, in turn, calls Dapr runtime.

The following architecture diagram illustrates the components that make up this quickstart:

![Architecture Diagram](./img/arch-diag1.png)

## Prerequisites
This quickstart requires you to have the following installed on your machine:
- [Docker](https://docs.docker.com/)
- [Go](https://golang.org/)

## Step 1 - Setup Dapr

Follow [instructions](https://docs.dapr.io/getting-started/install-dapr/) to download and install the Dapr CLI and initialize Dapr.

## Step 2 - Understand the code

> This demo has been modified. The code contained in `order.go` can be regarded as the business code will be complied to wasm, and the code contained in `main.go` is the code of wasm's host.

The [order.go](./order.go) is a simple application code, that has one function `serve` making different responses according to different routes:
* `put` sends an order with configurable order ID.
* `get` return the current order number.
* `del` deletes the order.


It interacts with the "host" through 3 functions:

Persist the state:
```go
    putState(storeName, orderID string) (errStr string)
```
Retrieve the state:
```go
    getState(storeName string) (value string, metadata map[string]string, errStr string)
```
Delete the state:
```go
    delState(storeName string) (errStr string)
```

## Step 3 - Run the app with Dapr

1. Build the app

<!-- STEP 
name: Build the app
-->

```bash
go mod vendor
go build -o order
```

<!-- END_STEP -->

2. Run the app

There are two ways to launch Dapr applications. You can pass the app executable to the Dapr runtime:

```bash
dapr run --app-id order-app --log-level error ./order
```
