# Menu Service

The Menu Service manages menu items (e.g., Burger, Pizza) for the restaurant system. It supports CRUD operations via REST APIs, storing data in memory.

## API Endpoints
- `POST /menu`: Create a new menu item (e.g., `{"name":"Salad","price":8.99}`).
- `GET /menu`: List all menu items.
- `GET /menu/{id}`: Get a menu item by ID.
- `PUT /menu/{id}`: Update a menu item.
- `DELETE /menu/{id}`: Delete a menu item.

## Local Setup
1. Navigate to the service directory:
   ```bash
   cd menu-service
   ```
2. Initialize Go module:
   ```bash
   go mod init menu_service
   ```
3. Run the service:
   ```bash
   go run menu_service.go
   ```
   The service runs on `http://localhost:8081`.

## Choreo Deployment
- The `component.yaml` configures the service for Choreo, exposing the API on port 8081.
- The `docs/openapi.yaml` defines the API schema.
- Deploy via Choreo UI or CLI, connecting to the GitHub repository’s `menu-service/` directory.
- No environment variables are required, as this service doesn’t make internal calls.

## Notes
- Data is in-memory and resets on restart.
- Test endpoints using `curl` or the Postman collection in the root `docs/` directory.