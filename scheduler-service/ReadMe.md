# Scheduler Service

The Scheduler Service invokes the GET endpoints of the restaurant system’s microservices (Menu, Order, Kitchen, Payment, Reservation) to monitor their availability. It executes these calls upon running and exits, with Choreo’s scheduler triggering periodic executions.

## Execution Details
The service calls the following GET endpoints:
- Menu Service: `GET /menu` (e.g., `http://localhost:8081/menu`)
- Order Service: `GET /orders` (e.g., `http://localhost:8082/orders`)
- Kitchen Service: `GET /kitchen/orders` (e.g., `http://localhost:8083/kitchen/orders`)
- Payment Service: `GET /payments` (e.g., `http://localhost:8084/payments`)
- Reservation Service: `GET /reservations` (e.g., `http://localhost:8085/reservations`)

## Local Setup
1. Ensure the Menu, Order, Kitchen, Payment, and Reservation Services are running:
   - Menu Service: `http://localhost:8081`
   - Order Service: `http://localhost:8082`
   - Kitchen Service: `http://localhost:8083`
   - Payment Service: `http://localhost:8084`
   - Reservation Service: `http://localhost:8085`
2. Navigate to the service directory:
   ```bash
   cd scheduler-service
   ```
3. Initialize Go module:
   ```bash
   go mod init scheduler_service
   ```
4. Set the environment variables:
   ```bash
   export MENU_SERVICE_URL=http://localhost:8081
   export ORDER_SERVICE_URL=http://localhost:8082
   export KITCHEN_SERVICE_URL=http://localhost:8083
   export PAYMENT_SERVICE_URL=http://localhost:8084
   export RESERVATION_SERVICE_URL=http://localhost:8085
   ```
5. Run the service:
   ```bash
   go run scheduler_service.go
   ```
   The service executes, calls the GET endpoints, and exits.

## Choreo Deployment
- The `component.yaml` configures the service as a Scheduled Task for Choreo, running every 5 minutes by default.
- Set the following environment variables in Choreo:
  - `MENU_SERVICE_URL` (e.g., `http://menu-service:8081` or Choreo’s assigned URL)
  - `ORDER_SERVICE_URL` (e.g., `http://order-service:8082` or Choreo’s assigned URL)
  - `KITCHEN_SERVICE_URL` (e.g., `http://kitchen-service:8083` or Choreo’s assigned URL)
  - `PAYMENT_SERVICE_URL` (e.g., `http://payment-service:8084` or Choreo’s assigned URL)
  - `RESERVATION_SERVICE_URL` (e.g., `http://reservation-service:8085` or Choreo’s assigned URL)
  - `API_TOKEN` (OAuth token or API key for authentication, if required)
- Deploy via Choreo UI or CLI, connecting to the GitHub repository’s `scheduler-service/` directory.
- The scheduled task runs the service periodically (e.g., every 5 minutes, configurable via `component.yaml`).

## Notes
- Requires Menu, Order, Kitchen, Payment, and Reservation Services for invoking their GET endpoints.
- Logs API call results to the console, resetting on each run.
- Test other services’ endpoints using `curl` or the Postman collection in the root `docs/` directory.