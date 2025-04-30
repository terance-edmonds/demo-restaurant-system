# Scheduler Service

The Scheduler Service invokes the GET endpoints of the restaurant system’s microservices (Menu, Order, Kitchen, Payment, Reservation) to monitor their availability. It supports triggering these calls via a REST API, with Choreo’s scheduler initiating periodic executions.

## API Endpoints
- `POST /trigger`: Trigger calls to all GET endpoints (e.g., Menu’s `/menu`, Order’s `/orders`).
- `GET /health`: Check the service’s health status.

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
   The service runs on `http://localhost:8086`.

## Choreo Deployment
- The `component.yaml` configures the service for Choreo, exposing the API on port 8086.
- The `docs/openapi.yaml` defines the API schema.
- Set the following environment variables in Choreo:
  - `MENU_SERVICE_URL` (e.g., `http://menu-service:8081` or Choreo’s assigned URL)
  - `ORDER_SERVICE_URL` (e.g., `http://order-service:8082` or Choreo’s assigned URL)
  - `KITCHEN_SERVICE_URL` (e.g., `http://kitchen-service:8083` or Choreo’s assigned URL)
  - `PAYMENT_SERVICE_URL` (e.g., `http://payment-service:8084` or Choreo’s assigned URL)
  - `RESERVATION_SERVICE_URL` (e.g., `http://reservation-service:8085` or Choreo’s assigned URL)
  - `API_TOKEN` (OAuth token or API key for authentication, if required)
- Deploy via Choreo UI or CLI, connecting to the GitHub repository’s `scheduler-service/` directory.
- Configure a scheduled task in Choreo to invoke `POST /trigger` periodically (e.g., every 5 minutes using cron `*/5 * * * *`).

## Notes
- Requires Menu, Order, Kitchen, Payment, and Reservation Services for invoking their GET endpoints.
- Logs API call results in memory, resetting on restart.
- Test endpoints using `curl` or the Postman collection in the root `docs/` directory.