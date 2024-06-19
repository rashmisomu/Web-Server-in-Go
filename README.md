# Web Server in Go

This project implements a concurrent web server and client in Go. The server handles multiple endpoints (/hello, /form, /time) and uses worker goroutines for background processing. The client sends concurrent requests to these endpoints and displays the server's responses.

## Table of Contents

- [Overview](#overview)
- [Server Features](#server-features)
- [Client Features](#client-features)
- [Project Structure](#project-structure)
- [Server Endpoints](#server-endpoints)
- [Server Worker Functionality](#server-worker-functionality)
- [Client Setup](#client-setup)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Server Setup and Usage](#server-setup-and-usage)
  - [Running the Server](#running-the-server)
  - [Server Endpoint Details](#server-endpoint-details)
- [Client Usage](#client-usage)
  - [Running the Client](#running-the-client)

## Overview

This project consists of a Go web server and client implemented in separate files. The server handles incoming requests at /hello, /form, and /time endpoints and uses worker goroutines for concurrent processing. The client sends multiple concurrent requests to these endpoints and displays the server's responses.

## Server Features

- **HTTP Server**:
  - Listens on port 8089.
  - Supports endpoints for /hello, /form, and /time.
  - Utilizes worker goroutines for background processing.

## Client Features

- **HTTP Client**:
  - Sends concurrent GET and POST requests to the server endpoints.
  - Logs responses received from the server.

## Project Structure

The project structure is organized as follows:

- **server.go**: Contains the server implementation.
- **client.go**: Contains the client implementation.
- **README.md**: This documentation file.

## Server Endpoints

### /hello

- **Method**: GET
- **Response**: "Accepted"
- **Description**: Responds with a simple "Accepted" message and logs the request timestamp.

### /form

- **Method**: POST
- **Request Parameters**:
  - `name`: Name of the person.
  - `address`: Address of the person.
- **Response**: "POST request successful\nName = {name}\nAddress = {address}"
- **Description**: Parses form data, responds with the parsed data, and logs the request timestamp.

### /time

- **Method**: GET
- **Response**: "Current time is: {current_time}"
- **Description**: Responds with the current server time and logs the request timestamp.

## Server Worker Functionality

- **Channel**: ch1 (Buffered channel of size 2)
- **Number of Workers**: 10
- **Worker Task**: Each worker listens on the channel for request timestamps and processes them by simulating a workload (sleeping for 5 seconds).

## Client Setup

### Prerequisites

- Go 1.16 or higher installed on your machine.

### Installation

 Clone the repository:

   ```sh
   git clone <repository_url>
   cd <repository_directory>
   ```
## Server Setup and Usage

### Running the Server

To start the server:

```sh
go run server.go
```
- The server will start listening on http://localhost:8089.

## Client Usage
### Running the Client
To run the client and send requests to the server:

```sh

go run client.go
```
- The client will send concurrent requests to the server endpoints and display the responses received.