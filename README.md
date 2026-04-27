# test.devices-api

Simple REST API done with Postgres 16, Golang 1.26.2 and Docker Compose.

## Running:

This project is provided with a Makefile that runs all the necessary commands.

To run this API locally, run `make init-db; make run-local`
- This will cleanup the DB, create a new instance with testdata and then start the Backend server.

To run the integration tests: `make test-integration`
- Test data is included in the migration at `migrations/init.sql`

## Available endpoints

### `GET /health` - Basic healthcheck

### `GET /devices` - List all available devices

Success (200 OK)
```json
[
    {
        "id": 1,
        "external_id": "77b85ecb-e767-4f2d-a3cf-be53dae49274",
        "name": "Test device 1",
        "brand": "Brand number one",
        "state": "available",
        "created_at": "2026-04-27T08:06:53Z"
    },
    {
        "id": 2,
        "external_id": "f3b4ea8f-1e69-4736-9cc7-f8775761fefd",
        "name": "Test device 2",
        "brand": "Brand number one",
        "state": "in-use",
        "created_at": "2026-04-27T08:06:53Z"
    }
]
```

### `GET /devices/:external_id` - List a specific device by UUID
Success (200 OK)
```json
{
    "id": 2,
    "external_id": "f3b4ea8f-1e69-4736-9cc7-f8775761fefd",
    "name": "Updated name",
    "brand": "Updated brand",
    "state": "inactive",
    "created_at": "2026-04-27T07:44:28Z",
    "updated_at": "2026-04-27T07:45:23Z"
}
```
Not Found (404 Not Found)
```json
{
    "error": "device with id 3c46c5a3-4fa9-4ce1-8693-12e467f03731 not found"
}
```

### `DELETE /devices/:external_id` - List a specific device by UUID
Success (204 No Content)
```json
// Empty body
```

Device is in use (409 Conflict)
```json
{
    "error": "device uuid d3f32eba-fc67-4440-b054-1cf821042610 is in-use"
}
```

Device doesn't exist (404 Not Found)
```json
{
    "error": "device with id 3c46c5a3-4fa9-4ce1-8693-12e467f03731 not found"
}
```

### `POST /devices` - Create a new device

Example Request
```json
{
  "name": "test name",
  "brand": "test brand",
  "state": "in-use" // in-use, inactive, available are the only valid values
}
```
Success (201 Created / 200 OK)
```json
{
    "id": 2,
    "external_id": "f3b4ea8f-1e69-4736-9cc7-f8775761fefd",
    "name": "Updated name",
    "brand": "Updated brand",
    "state": "inactive",
    "created_at": "2026-04-27T07:44:28Z",
    "updated_at": "2026-04-27T07:45:23Z"
}
```
Device already exists and is soft deleted (409 Conflict)
```json
{
    "error": "device Test device 7 (soft deleted) brand Brand number three uuid 3c46c5a3-4fa9-4ce1-8693-12e467f03730 already exists and it's soft deleted"
}
```
Validation errors (400 Bad Request)
```json
{
    "details": "Key: 'createDeviceRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag\nKey: 'createDeviceRequest.Brand' Error:Field validation for 'Brand' failed on the 'required' tag\nKey: 'createDeviceRequest.State' Error:Field validation for 'State' failed on the 'required' tag",
    "error": "invalid request"
}
```
Invalid state (400 Bad Request)
```json
{
    "error": "invalid state: invalid-state"
}
```

### `PATCH /devices/:external_id` - Update name, brand or state from an existing device

Success (200 OK)
```json
{
    "id": 3,
    "external_id": "c343dabd-0afd-4f5b-a632-42be259df112",
    "name": "Updated name",
    "brand": "Updated brand",
    "state": "in-use",
    "created_at": "2026-04-27T08:06:53Z",
    "updated_at": "2026-04-27T08:07:24Z"
}
```
Device already exists (409 Conflict)
```json
{
    "error": "device with this name and brand already exists"
}
```
Device is in-use and can't be altered (409 Conflict)
```json
{
    "error": "device uuid f3b4ea8f-1e69-4736-9cc7-f8775761fefd is in-use"
}
```
Device not found (404 Not Found)
```json
{
    "error": "device with id f3b4ea8f-1e69-4736-9cc7-f8775761fefa not found"
}
```
Invalid State (400 Bad Request)
```json
{
    "error": "invalid state: invalid-state"
}
```
---

## Architectural decisions

- Code followed a Clean Architecture approach
- Used [Standard Go Project Layout](https://github.com/golang-standards/project-layout) to organize the code.
- Using Docker image with Alpine Linux to speed-up setup.
  - For production, a distroless image is better for increased security.

### Why GORM?
- Specifically due to time constraints.
  - This is the tool i've used in my last 2 companies, mainly due to business requirements.
  - A query builder such as SQLC or SQLX would be much better, as they're more flexible for the complex queries that would be needed to assure a good performance.

### Why Gin?
- It's the framework i used in my last project, the most complex i worked with.
- The whole code architecture was pretty good and easy to work with, so it made life much easier.
  - This is the architecture i partially replicated here, why changes and improvements when necessary.
- Previously worked with Echo, but there was rarely time to improve the arch, so working with it would take more time than i have available.

---

## TO-DO for production (stuff not done due to time constraints)

- 1 Bug/Edge case not addressed:
  - If Device is updated to a name/brand combination, and try to create another device with same name/brand, it is created instead of returning an error.
  - If the opposite is done (create the device then try to update another to same name/brand combination), it correctly returns an error.
- Add some missing test cases / Automate test coverage report
- Use Swagger documentation / Automate with [Swag](https://github.com/swaggo/swag)
- Think about DB connection pool (max open/idle connections)
  - Analyze the predicted read/write ratio and make changes as necessary, for example:
    - If race conditions can happen on Device creation, this needs to be handled, preferably with DB locks and GORM transactions;
    - Same for updating devices, they need to be locked to avoid race conditions and data consistency problems;
- Improve the test runner, current integration tests don't have a very good maintenability;
