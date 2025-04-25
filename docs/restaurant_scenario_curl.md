# Restaurant System Microservices: Curl Scenario

This document outlines a realistic scenario for interacting with the restaurant system’s five microservices (Menu, Order, Kitchen, Payment, Reservation) using `curl` commands. The scenario demonstrates a customer ordering food, paying, having the order processed, and making a reservation, alongside a manager updating the menu. Each service’s CRUD operations are exercised, showcasing inter-service dependencies (e.g., Order Service validating menu items via Menu Service).

## Prerequisites
- The microservices are running locally on:
  - Menu Service: `http://localhost:8081`
  - Order Service: `http://localhost:8082`
  - Kitchen Service: `http://localhost:8083`
  - Payment Service: `http://localhost:8084`
  - Reservation Service: `http://localhost:8085`
- Follow the setup instructions in the `README.md` to start the services (locally with `go run` or via Docker).
- Ensure services are started in order (Menu, Order, Payment, Kitchen, Reservation) due to dependencies.

## Scenario Workflow

### 1. Menu Management
The restaurant manager adds a new menu item (Salad) and lists all menu items to confirm.

#### Create a Menu Item
```bash
curl -X POST http://localhost:8081/menu -H "Content-Type: application/json" -d '{"name":"Salad","price":8.99}'
```
**Expected Response**:
```json
{"id":4,"name":"Salad","price":8.99}
```

#### List All Menu Items
```bash
curl http://localhost:8081/menu
```
**Expected Response**:
```json
[
  {"id":1,"name":"Burger","price":10.99},
  {"id":2,"name":"Pizza","price":12.99},
  {"id":3,"name":"Soda","price":2.99},
  {"id":4,"name":"Salad","price":8.99}
]
```

### 2. Order Placement
The customer places an order for a Burger (ID: 1) and a Soda (ID: 3), then checks the order details. The Order Service validates the menu items by calling the Menu Service.

#### Create an Order
```bash
curl -X POST http://localhost:8082/orders -H "Content-Type: application/json" -d '{"item_ids":[1,3]}'
```
**Expected Response**:
```json
{"id":1,"item_ids":[1,3],"total":13.98,"status":"placed"}
```
**Note**: Total is calculated as $10.99 (Burger) + $2.99 (Soda) = $13.98.

#### Get Order Details
```bash
curl http://localhost:8082/orders/1
```
**Expected Response**:
```json
{"id":1,"item_ids":[1,3],"total":13.98,"status":"placed"}
```

### 3. Payment Processing
The customer pays for the order. The Payment Service validates the order ID and total by calling the Order Service.

#### Create a Payment
```bash
curl -X POST http://localhost:8084/payments -H "Content-Type: application/json" -d '{"order_id":1}'
```
**Expected Response**:
```json
{"id":1,"order_id":1,"amount":13.98,"status":"completed"}
```

#### Get Payment Details
```bash
curl http://localhost:8084/payments/1
```
**Expected Response**:
```json
{"id":1,"order_id":1,"amount":13.98,"status":"completed"}
```

### 4. Kitchen Processing
The kitchen processes the order, marking it as "ready," and later updates the status to "served."

#### Create a Processed Order
```bash
curl -X POST http://localhost:8083/kitchen/orders -H "Content-Type: application/json" -d '{"status":"ready"}'
```
**Expected Response**:
```json
{"id":1,"status":"ready"}
```

#### Update Processed Order Status
```bash
curl -X PUT http://localhost:8083/kitchen/orders/1 -H "Content-Type: application/json" -d '{"status":"served"}'
```
**Expected Response**:
```json
{"id":1,"status":"served"}
```

### 5. Reservation Booking
The customer makes a reservation for a future visit and later cancels it.

#### Create a Reservation
```bash
curl -X POST http://localhost:8085/reservations -H "Content-Type: application/json" -d '{"customer_name":"John Doe","time":"2025-04-26T19:00:00Z","table_number":5}'
```
**Expected Response**:
```json
{"id":1,"customer_name":"John Doe","time":"2025-04-26T19:00:00Z","table_number":5}
```

#### Delete the Reservation
```bash
curl -X DELETE http://localhost:8085/reservations/1
```
**Expected Response**: (204 No Content)

This scenario demonstrates the full functionality of the restaurant system, covering CRUD operations and inter-service interactions.