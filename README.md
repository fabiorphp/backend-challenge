# Basket
Basket is a small app for managing checkout baskets.

## Requirements
App requires Golang 1.11 or later.

## Installation
- Install [Golang](https://golang.org/doc/install)

## Build
For building binaries please use make, look at the commands bellow:

```
// Build the binary in your environment.
$ make build

// Build with another OS. Default Linux
$ make OS=darwin build

// Build with custom version.
$ make APP_VERSION=0.1.0 build

// Build with custom app name.
$ make APP_NAME=basket build

// Passing all flags.
$ make OS=darwin APP_NAME=basket APP_VERSION=0.1.0 build

// Clean Up.
$ make clean

// Create the required folders.
$ make configure
```

## Develpoment
```
// Running tests
$ make test

// Running tests with coverage. Output coverage file: coverage.html
$ make cover

// Run the application without build
$ go run ./cmd/basket/main.go
```

## Running
After build the application the should follow the steps bellow for running.
```sh
$ make
$ ./build/basket agent # Running the agent server

# In another terminal

# Execute the app binary to create a basket
$ ./build/basket create

# Add product into basket. (Products code: VOUCHER, TSHIRT and MUG)
$ ./build/basket add <BASKET_ID> VOUCHER

# Get the amount
$ ./build/basket amount <BASKET_ID>

# Delete basket
$ ./build/basket delete <BASKET_ID>
```

## Challenge instructions
Besides providing exceptional transportation services, Cabify also runs a physical store which sells 3 products:

```
Code         | Name                |  Price
-------------------------------------------------
VOUCHER      | Cabify Voucher      |   5.00€
TSHIRT       | Cabify T-Shirt      |  20.00€
MUG          | Cafify Coffee Mug   |   7.50€
```

Various departments have insisted on the following discounts:

 * The marketing department thinks a buy 2 get 1 free promotion will work best (buy two of the same product, get one free), and would like this to only apply to `VOUCHER` items.

 * The CFO insists that the best way to increase sales is with discounts on bulk purchases (buying x or more of a product, the price of that product is reduced), and requests that if you buy 3 or more `TSHIRT` items, the price per unit should be 19.00€.

This set of rules to apply may change quite frequently in the future.

Your task is to implement a checkout system for this store.

The system should have differentiated client and server components that communicate over the network.

The server should expose the following independent operations:

- Create a new checkout basket
- Add a product to a basket
- Get the total amount in a basket
- Remove the basket

The server must support concurrent invocations of those operations: any of them may be invoked at any time, while other operations are still being performed, even for the same basket.

The client must connect user input with those operations via the protocol exposed by the server.

We don't have any DBAs at Cabify, so the service shouldn't use any external databases of any kind.

Using Go, implement a checkout service and its client that fulfils these requirements.

Examples:

    Items: VOUCHER, TSHIRT, MUG
    Total: 32.50€

    Items: VOUCHER, TSHIRT, VOUCHER
    Total: 25.00€

    Items: TSHIRT, TSHIRT, TSHIRT, VOUCHER, TSHIRT
    Total: 81.00€

    Items: VOUCHER, TSHIRT, VOUCHER, VOUCHER, MUG, TSHIRT, TSHIRT
    Total: 74.50€

**The code should:**
- Build and execute in a Unix operating system.
- Be written as production-ready code. You will write production code.
- Be easy to grow and easy to add new functionality.
- Have notes attached, explaining the solution and why certain things are included and others are left out.
- If submitted as a compressed package, its size must not exceed 20MB.
