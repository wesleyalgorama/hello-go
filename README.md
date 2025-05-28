# Go Hello Server

A minimal HTTP server built with Go that responds with a simple greeting message.

## About

This project implements a basic web server using Go’s standard `net/http` package. It listens on port `8080` and returns a plain text response to any request made to the root endpoint (`/`).

## Table of Contents

* [About](#about)
* [Getting Started](#getting-started)

  * [Running Locally](#running-locally)
  * [Running with Docker](#running-with-docker)
* [Usage](#usage)
* [Contributing](#contributing)
* [Licensing](#licensing)

## Getting Started

### Running Locally

**1 - Prerequisites**

* [Go](https://golang.org/dl/) (v3.0 or later)

**2 - Clone the repository**

```bash
git clone https://github.com/wesleyalgorama/go-hello-server.git
cd go-hello-server
```

**3 - Run the server**

```bash
go run main.go
```

The server will start on [http://localhost:8080](http://localhost:8080).

### Running with Docker

**1 - Build the Docker image**

```bash
docker build -t go-hello-server .
```

**2 - Run the container**

```bash
docker run -d -p 8080:8080 --name go-hello-server go-hello-server
```

The server will now be available at [http://localhost:8080](http://localhost:8080).

To stop and remove the container:

```bash
docker stop go-hello-server
docker rm go-hello-server
```

## Usage

Send a request to the root endpoint:

```bash
curl http://localhost:8080/
```

**Response:**

```
✅ Hello from Go!
```

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests with improvements or new features.

## Licensing

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
