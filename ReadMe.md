# Restaurant System Microservices

This repository contains a restaurant system implemented as five microservices in Go. Each service supports full CRUD operations, uses in-memory data storage (no database), and communicates via HTTP/REST. The system is organized into Finance, Management, Customer Management, and Operations categories, making it a great example for learning microservices architecture with Go.

## System Overview

The system consists of five microservices:
1. **Menu Service** (Operations): Manages menu items (e.g., Burger, Pizza).
   - Endpoints: `POST /menu`, `GET /menu`, `GET /menu/{id}`, `PUT /menu/{id}`, `DELETE /menu/{id}`
2. **Order Service** (Customer Management): Manages customer orders, validates menu items.
   - Endpoints: `POST /orders`, `GET /orders`, `GET /orders/{id}`, `PUT /orders/{id}`, `DELETE /orders/{id}`
3. **Kitchen Service** (Operations): Manages processed orders (e.g., marking as "ready").
   - Endpoints: `POST /kitchen/orders`, `GET /kitchen/orders`, `GET /kitchen/orders/{id}`, `PUT /kitchen/orders/{id}`, `DELETE /kitchen/orders/{id}`
4. **Payment Service** (Finance): Manages order payments, validates orders.
   - Endpoints: `POST /payments`, `GET /payments`, `GET /payments/{id}`, `PUT /payments/{id}`, `DELETE /payments/{id}`
5. **Reservation Service** (Management): Manages table reservations.
   - Endpoints: `POST /reservations`, `GET /reservations`, `GET /reservations/{id}`, `PUT /reservations/{id}`, `DELETE /reservations/{id}`

## Prerequisites

- Go 1.21 or later
- Docker (optional, for containerized deployment)
- Git

## Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/<your-username>/restaurant-system.git
cd restaurant-system
```

### 2. Run Services Locally

Each service is a standalone Go application. Navigate to each service directory and run it.

#### Menu Service
```bash
cd menu-service
go mod init menu_service
go run menu_service.go
```
Runs on `http://localhost:8081`.

#### Order Service
```bash
cd order-service
go mod init order_service
go run order_service.go
```

Update environment variables

```env
MENU_SERVICE_URL=http://menu-service:8081
```

Runs on `http://localhost:8082`. Requires Menu Service running.

#### Kitchen Service
```bash
cd kitchen-service
go mod init kitchen_service
go run kitchen_service.go
```
Runs on `http://localhost:8083`.

#### Payment Service
```bash
cd payment-service
go mod init payment_service
go run payment_service.go
```

Update environment variables

```env
ORDER_SERVICE_URL=http://localhost:8082
```

Runs on `http://localhost:8084`. Requires Order Service running.

#### Reservation Service
```bash
cd reservation-service
go mod init reservation_service
go run reservation_service.go
```
Runs on `http://localhost:8085`.

**Note**: Start services in order (Menu, Order, Payment, Kitchen, Reservation) due to dependencies (Order Service calls Menu Service, Payment Service calls Order Service).

### 3. Run with Docker

Each service includes a Dockerfile for containerized deployment.

1. **Build Docker Images**:
   For each service directory:
   ```bash
   cd <service-directory>
   docker build -t <service-name>:latest .
   cd ..
   ```
   Example for Menu Service:
   ```bash
   cd menu-service
   docker build -t menu-service:latest .
   cd ..
   ```

2. **Run Containers**:
   Run each service in a Docker network to enable communication:
   ```bash
   docker network create restaurant-net
   docker run -d --name menu-service --network restaurant-net -p 8081:8081 menu-service:latest
   docker run -d --name order-service --network restaurant-net -p 8082:8082 -e MENU_SERVICE_URL=http://menu-service:8081 order-service:latest
   docker run -d --name kitchen-service --network restaurant-net -p 8083:8083 kitchen-service:latest
   docker run -d --name payment-service --network restaurant-net -p 8084:8084 -e ORDER_SERVICE_URL=http://order-service:8082 payment-service:latest
   docker run -d --name reservation-service --network restaurant-net -p 8085:8085 reservation-service:latest
   ```

### 4. Test the System

Use `curl`, Postman, or any HTTP client to test the APIs. Example for Menu Service:

1. **Create Menu Item**:
   ```bash
   curl -X POST http://localhost:8081/menu -H "Content-Type: application/json" -d '{"name":"Salad","price":8.99}'
   ```
   Response: `{"id":4,"name":"Salad","price":8.99}`

2. **List Menu Items**:
   ```bash
   curl http://localhost:8081/menu
   ```
   Response: `[{"id":1,"name":"Burger","price":10.99},...,{"id":4,"name":"Salad","price":8.99}]`

3. **Get Menu Item**:
   ```bash
   curl http://localhost:8081/menu/4
   ```
   Response: `{"id":4,"name":"Salad","price":8.99}`

4. **Update Menu Item**:
   ```bash
   curl -X PUT http://localhost:8081/menu/4 -H "Content-Type: application/json" -d '{"name":"Caesar Salad","price":9.99}'
   ```
   Response: `{"id":4,"name":"Caesar Salad","price":9.99}`

5. **Delete Menu Item**:
   ```bash
   curl -X DELETE http://localhost:8081/menu/4
   ```
   Response: (204 No Content)

Repeat for other services (Order, Kitchen, Payment, Reservation) with their respective endpoints and ports.

## Dependencies

- **Order Service**: Calls Menu Service (`http://menu-service:8081/menu`) to validate menu items.
- **Payment Service**: Calls Order Service (`http://order-service:8082/orders`) to validate orders.
- Ensure services are running and accessible (use Docker network or localhost).

## Notes

- **In-Memory Data**: Data is stored in memory and resets on service restart.
- **Go Modules**: Each service initializes a Go module (`go mod init`). For production, run `go mod tidy` to clean dependencies.
- **Scalability**: The system is a demo. For production, consider adding a database, logging, and authentication.
- **Testing**: Use the provided `curl` commands or write automated tests with Go’s `testing` package.

Checkout [Restaurant System Microservices: Curl Scenario](./docs/restaurant_scenario_curl.md) for an example restaurant scenario.