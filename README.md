# Order Book

![GitHub](https://img.shields.io/github/license/markraiter/order-book)  ![GitHub top language](https://img.shields.io/github/languages/top/markraiter/order-book)  [![Go Report Card](https://goreportcard.com/badge/github.com/markraiter/order-book)](https://goreportcard.com/report/github.com/markraiter/order-book)

This is a simple implementation of an order book in Go. An order book is a list of buy and sell orders for a specific security or financial instrument, organized by price level.

## Features

- Insert orders into the order book
- Match orders and calculate balance changes
- Print balance changes to the console

## Efficiency

The efficiency of this solution comes from its use of sorting and binary search, which are both efficient algorithms for dealing with ordered data.

When an order is inserted into the order book, it is added to either the `BuyOrders` or `SellOrders` slice depending on its side. These slices are kept sorted by price, which allows us to efficiently find matching orders using binary search. This operation is O(log n) in the worst case, where n is the number of orders in the slice.

The `MatchOrders` function iterates over the buy and sell orders and matches them where possible. This operation is O(n), where n is the total number of orders. However, because the orders are sorted, we can use a two-pointer technique to match orders in a single pass, which is more efficient than a naive approach that would require nested loops.

The overall time complexity of the solution is therefore O(n log n) due to the sorting of orders. The space complexity is O(n), where n is the total number of orders, as all orders are stored in memory.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go (version 1.22.0)

### Installing

A step by step series of examples that tell you how to get a development environment running.

1. Clone the repository.
2. Run the app by `go run main.go`.
3. Provide all the nessesary information and recieve feedback.

## Running the tests

1. Run the tests with `go test ./...`

## Built With

- [Go](https://golang.org/) - The programming language used.