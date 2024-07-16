# Go API Service Dockerized

This repository contains a sample Go application that serves as an HTTP API service for managing messages. The application is containerized using Docker for easy deployment and scalability.

## Features

- **HTTP API**: Provides endpoints for managing messages with sender and receiver details.
- **Database**: Uses PostgreSQL for data storage.
- **Dockerized**: Ready for deployment using Docker, ensuring consistency across different environments.

## Prerequisites

- Docker: Ensure Docker is installed on your system. [Docker Installation Guide](https://docs.docker.com/get-docker/)
- Go (Optional for local development): If you plan to develop locally without Docker, install Go. [Go Installation Guide](https://golang.org/doc/install)

## Getting Started

Follow these steps to get the project up and running locally or in a Docker environment:

### 1. Clone the Repository

```bash
git clone <repository-url>
cd <repository-directory>
```

### 2. Build and Run with Docker

```bash
docker build -t go-api-service .
docker run -p 8080:8080 go-api-service
```

### 3. Access the API

Once the Docker container is running, you can access the API endpoints:

- Base URL: `http://localhost:8080`

#### Endpoints:

- **GET `/get/messages/<account_id>`**: Returns all messages for the specified account ID.
- **POST `/create`**: Creates a new message with sender and receiver details.
- **GET `/search?message_id=<id>`**: Searches messages by message ID.
- **GET `/search?sender_number=<number>`**: Searches messages by sender number.
- **GET `/search?receiver_number=<number>`**: Searches messages by receiver number.

### 4. Configuration

For database configuration, set environment variables for connection details (`DB_HOST`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`) either in your Docker environment or a `.env` file. 

### 5. Development

For local development without Docker:

```bash
# Install dependencies
go mod download

# Run the application
go run main.go
```

### 6. Testing

Ensure proper test coverage for your application. Run tests using:

```bash
go test ./...
```

## Contributing

Contributions are welcome! Please fork the repository and submit pull requests or open issues for any bugs, feature requests, or improvements.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---