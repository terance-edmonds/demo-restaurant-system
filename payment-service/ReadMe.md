# Payment Service

The Payment Service manages order payments, validating orders against the Order Service. It supports CRUD operations via REST APIs, storing data in memory.

## API Endpoints
- `POST /payments`: Create a new payment (e.g., `{"order_id":1}`).
- `GET /payments`: List all payments.
- `GET /payments/{id}`: Get a payment by ID.
- `PUT /payments/{id}`: Update a payment.
- `DELETE /payments/{id}`: Delete a payment.

## Local Setup
1. Ensure the Order Service is running (`http://localhost:8082`).
2. Navigate to the service directory:
   ```bash
   cd payment-service
   ```
3. Initialize Go module:
   ```bash
   go mod init payment_service
   ```
4. Set the environment variable:
   ```bash
   export ORDER_SERVICE_URL=http://localhost:8082
   ```
5. Run the service:
   ```bash
   go run payment_service.go
   ```
   The service runs on `http://localhost:8084`.

## Choreo Deployment
- The `component.yaml` configures the service for Choreo, exposing the API on port 8084.
- The `docs/openapi.yaml` defines the API schema.
- Set the `ORDER_SERVICE_URL` environment variable in Choreo to the internal Order Service URL (e.g., `http://order-service:8082` or Choreo’s assigned URL).
- Deploy via Choreo UI or CLI, connecting to the GitHub repository’s `payment-service/` directory.

## Notes
- Requires Order Service for order validation.
- Data is in-memory and resets on restart.
- Test endpoints using `curl` or the Postman collection in the root `docs/` directory.