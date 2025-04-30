# Kitchen Service

The Kitchen Service manages processed orders (e.g., marking as "ready" or "served"). It supports CRUD operations via REST APIs, storing data in memory.

## API Endpoints
- `POST /kitchen/orders`: Create a new processed order (e.g., `{"status":"ready"}`).
- `GET /kitchen/orders`: List all processed orders.
- `GET /kitchen/orders/{id}`: Get a processed order by ID.
- `PUT /kitchen/orders/{id}`: Update a processed order.
- `DELETE /kitchen/orders/{id}`: Delete a processed order.

## Local Setup
1. Navigate to the service directory:
   ```bash
   cd kitchen-service
   ```
2. Initialize Go module:
   ```bash
   go mod init kitchen_service
   ```
3. Run the service:
   ```bash
   go run kitchen_service.go
   ```
   The service runs on `http://localhost:8083`.

## Choreo Deployment
- The `component.yaml` configures the service for Choreo, exposing the API on port 8083.
- The `docs/openapi.yaml` defines the API schema.
- Deploy via Choreo UI or CLI, connecting to the GitHub repository’s `kitchen-service/` directory.
- No environment variables are required, as this service doesn’t make internal calls.

## Notes
- Data is in-memory and resets on restart.
- Test endpoints using `curl` or the Postman collection in the root `docs/` directory.