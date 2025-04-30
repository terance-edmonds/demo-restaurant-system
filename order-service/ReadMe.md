# Order Service

The Order Service manages customer orders, validating menu items against the Menu Service. It supports CRUD operations via REST APIs, storing data in memory.

## API Endpoints
- `POST /orders`: Create a new order (e.g., `{"item_ids":[1,3]}`).
- `GET /orders`: List all orders.
- `GET /orders/{id}`: Get an order by ID.
- `PUT /orders/{id}`: Update an order.
- `DELETE /orders/{id}`: Delete an order.

## Local Setup
1. Ensure the Menu Service is running (`http://localhost:8081`).
2. Navigate to the service directory:
   ```bash
   cd order-service
   ```
3. Initialize Go module:
   ```bash
   go mod init order_service
   ```
4. Set the environment variable:
   ```bash
   export MENU_SERVICE_URL=http://localhost:8081
   ```
5. Run the service:
   ```bash
   go run order_service.go
   ```
   The service runs on `http://localhost:8082`.

## Choreo Deployment
- The `component.yaml` configures the service for Choreo, exposing the API on port 8082.
- The `docs/openapi.yaml` defines the API schema.
- Set the `MENU_SERVICE_URL` environment variable in Choreo to the internal Menu Service URL (e.g., `http://menu-service:8081` or Choreo’s assigned URL).
- Deploy via Choreo UI or CLI, connecting to the GitHub repository’s `order-service/` directory.

## Notes
- Requires Menu Service for item validation.
- Data is in-memory and resets on restart.
- Test endpoints using `curl` or the Postman collection in the root `docs/` directory.