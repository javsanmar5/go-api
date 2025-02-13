# HTTP JSON Server in Go (No Libraries)

A simple HTTP JSON server implemented with a minimalist approach, built for learning purposes. The server handles HTTP requests and responses manually but uses the gorilla/mux package for routing to simplify request handling. In the future, I may try to implement routing and other aspects from scratch to further my learning, but that’s not my current goal.

## Features

- Handles HTTP requests and responses manually
- Serves basic API endpoints

## Installation

Ensure you have Go installed, then clone the repository:

```sh
git clone https://github.com/javsanmar5/go-api.git
cd go-api
```

## Usage

To build and run the server we use Makefile. You can run:
```bash
make run
```

Also you can run it manually without building it with 
```bash
go run main.go
```
