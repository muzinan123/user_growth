# User Growth Management System

This repository demonstrates a User Growth Management System developed in Go using gRPC and Gin. The system is designed to manage user data and track user growth metrics.

## Prerequisites

- Go installed (version >= 1.16)
- Docker installed (for running the database)
- Make installed (for running make commands)
- Protobuf compiler (protoc) installed

## Project Structure
![18](https://github.com/user-attachments/assets/cde416ac-51be-4ac5-840a-fdcbe5e26fec)

```
user_growth/
├── comm/
├── conf/
├── dao/
├── database/
├── dbhelper/
├── grpc-go-master/
├── mainclient/
├── mainserver/
├── models/
├── pb/
├── service/
├── ugserver/
├── README.md
├── go.mod
└── go.sum
```

## Features

- User data management
- gRPC server for efficient communication
- Gin-based HTTP server for RESTful API
- Database integration (likely using a SQL database)
- Protobuf definitions for data serialization

## Getting Started

1. Clone the repository:
   ```
   git clone https://github.com/muzinan123/user_growth.git
   cd user_growth
   ```

2. Install dependencies:
   ```
   go mod download
   ```

3. Set up the database (instructions would be provided here)

4. Run the gRPC server:
   ```
   go run mainserver/main.go
   ```

5. Run the HTTP server:
   ```
   go run ugserver/main.go
   ```

## Testing

To run the tests:
```
go test ./...
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
