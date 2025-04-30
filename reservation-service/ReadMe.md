# Reservation Service

The Reservation Service manages table reservations for the restaurant. It supports CRUD operations via REST APIs, storing data in memory.

## API Endpoints
- `POST /reservations`: Create a new reservation (e.g., `{"customer_name":"John Doe","time":"2025-04-26T19:00:00Z","table_number":5}`).
- `GET /reservations`: List all reservations.
- `GET /reservations/{id}`: Get a reservation by ID.
- `PUT /reservations/{id}`: Update a reservation.
- `DELETE /reservations/{id}`: Delete a reservation.

## Local Setup
1. Navigate to the service directory:
   ```bash
   cd reservation-service
   ```
2. Initialize Go module:
   ```bash
   go mod init reservation_service
   ```
3. Run the service:
   ```bash
   go run reservation_service.go
   ```
   The service runs on `http://localhost:8085`.

## Choreo Deployment
- The `component.yaml` configures the service for Choreo, exposing the API on port 8085.
- The `docs/openapi.yaml` defines the API schema.
- Deploy via Choreo UI or CLI, connecting to the GitHub repository’s `reservation-service/` directory.
- No environment variables are required, as this service doesn’t make internal calls.

## Notes
- Data is in-memory and resets on restart.
- Test endpoints using `curl` or the Postman collection in the root `docs/` directory.