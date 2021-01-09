<div align="center">
<img src="./assets/gopher-grpc.jpg">
</div>

# GO GRPC TUTORIAL

This is a project with a basic example of grpc in golang for study purposes.

* [Installation](#installation)
* [Running](#running)
* [References](#references)

## Installation

In root folder of project, run:

```
go get -u ./...
```

## Running

### Server

First run the grpc server using the following command:

```
go run ./server/cmd/main.go 
```

The following message should appear on the console: `INFO[0000] GRPC server is running on port: 8000`.

### Client

Then run the client with the following command:

```
go run ./client/cmd/main.go 
```

The client will call the server's `AddBook` function simulating the insertion of a book and showing an id on the console.

Example: `INFO[0000] a new was book added with id: 1fc6d548-356a-42cc-92ce-d03fb4758117`

## References

* [GO GRPC Github](https://github.com/grpc/grpc-go)
* [GRPC Quickstart GO](https://grpc.github.io/docs/quickstart/go.html)
