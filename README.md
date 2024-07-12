# Golang Command Execution Server

This project is a simple HTTP server written in Go that executes shell commands received via a POST request and returns the output as JSON.

## Features

- Executes shell commands from HTTP requests
- Returns command output and error messages in JSON format

## Prerequisites

- [Go](https://golang.org/dl/) (version 1.20.5 or later)

## Installation

1. **Clone the repository:**

   ```sh
   git clone https://github.com/YOUR_USERNAME/golang-cmd-server.git
   cd golang-cmd-server

    Install Go if you haven't already:
        Download and install Go from the official website.

Usage

    Run the server:

    sh
    go run main.go


The server will start on port 8080.

Send a POST request to execute a command:

Use curl or Postman to send a POST request to http://localhost:8080/api/cmd with a JSON body containing the command you want to execute.
```sh
Example with curl:

curl -X POST -H "Content-Type: application/json" -d "{\"command\":\"dir\"}" http://localhost:8080/api/cmd

On Windows PowerShell:

sh

curl -X POST -H "Content-Type: application/json" -d '{\"command\":\"dir\"}' http://localhost:8080/api/cmd

On Linux/macOS:

sh

    curl -X POST -H "Content-Type: application/json" -d '{"command":"ls -l"}' http://localhost:8080/api/cmd

    Response:

    The server will return the command output and any error messages in JSON format.

Project Structure

css

golang-cmd-server/
  ├── main.go
  └── README.md

Contributing

Contributions are welcome! Please open an issue or submit a pull request.
