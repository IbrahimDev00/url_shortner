
# Go URL Shortener

## Overview

This is a URL shortener application built using Go and Docker. It allows users to shorten URLs, edit and customize shortened URLs, and add tags to their URLs for better organization and management.

## Features

- **Shorten URLs**: Convert long URLs into shorter, more manageable links.
- **Edit and Customize URLs**: Modify existing shortened URLs and customize their appearance.
- **Add Tags**: Organize and categorize your URLs by adding tags.

## Prerequisites

- [Go](https://golang.org/doc/install) (version 1.17 or later)
- [Docker](https://docs.docker.com/get-docker/)

## Getting Started

### Clone the Repository

```sh
git clone https://github.com/yourusername/your-repo.git
cd your-repo
```

### Build and Run with Docker

1. **Build and run the application using Docker Compose:**

    ```sh
    docker compose up --build
    ```

2. **Access the application:**
   - The application will be accessible at `http://localhost:8000`.

## Project Structure

```plaintext
.
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
├── handlers/
│   └── url_handler.go
├── models/
│   └── url.go
├── storage/
│   └── redis.go
└── README.md
```

## Dockerfile

The `Dockerfile` sets up the Go environment, installs dependencies, and builds the application.

```Dockerfile
# Use the official Golang image as the base image
FROM golang:1.17

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download the dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o /project/go-docker/build/myapp .

# Command to run the application
CMD ["/project/go-docker/build/myapp"]
```

## Docker Compose

The `docker-compose.yml` file defines the services for the application and the database.

```yaml
version: '3'

services:
  api:
    build:
      context: .
      dockerfile: Dockerfile
    ports:
      - "8000:8080"
    depends_on:
      - db

  db:
    build: ./db
    ports:
      - "6379:6379"
    volumes:
      - .data:/data
```

## Usage

- **Shorten a URL**: Use the `/shorten` endpoint to shorten a URL.
- **Edit a URL**: Use the `/edit` endpoint to edit an existing shortened URL.
- **Add Tags**: Use the `/add-tag` endpoint to add tags to a URL.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any changes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

If you have any questions or feedback, please contact me at chikani.ibrahim6@gmail.com
