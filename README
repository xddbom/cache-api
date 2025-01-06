# Redis Cache API
---
Minimalistic API based on Gin v1.10.0 and Redis v7.4.0, designed for simple user data caching via HTTP requests.

## Configuration
---
- Default Redis `PORT` (`:6379`)
- Default HTTP `PORT` (`:8080`)
- The `User` information you will work with is located in the file `./user/models.go`

## Routes
---
- `/` - Native root
  > "Welcome to Cache API!"

- `/health` - Redis connection check
  > "Redis is reachable!"

- `POST /user` - Creates a new user with a TTL of 60 seconds
  > "User created successfully!"

- `GET /user/:id` - Returns the specified user and extends their TTL by 60 seconds
  > *You will get user info or a relevant error.*

- `DELETE /user/:id` - Deletes the selected user
  > "User deleted successfully!"
  > *You will receive information about the user you just deleted.*

### Example of Requests
- POST
```bash
curl -X POST http://localhost:8080/user \
-H "Content-Type: application/json" \
-d '{"id": "123", "name": "Alice", "email": "alice@example.com", "age": 25}'
```

- GET
```bash
curl -X GET http://localhost:8080/user/123
```

- DELETE
```bash
curl -X DELETE http://localhost:8080/user/123
```

# Running the Application

## Installation
---
- Clone the repository:
  ```bash
  git clone https://github.com/xddbom/echo-server.git
  ```
*Note: Docker support is not available yet :(*

## Dependencies
---
- [Redis](https://redis.io/downloads/)

## Start Process
---
1. Start the Redis server using:
   ```bash
   redis-server
   ```
2. Start the HTTP server using:
   ```bash
   go run cmd/main.go
   ```

If everything is set up correctly, you should see logs like:
```bash
Connection successful!
...
[GIN-debug] Listening and serving HTTP on :8080
```
