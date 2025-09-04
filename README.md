# Go Gin Load Balancer

A high-performance, pluggable HTTP load balancer implemented in Go using the Gin framework. Easily route incoming requests across multiple backend microservices using configurable load balancing strategies, with custom logic for forwarding requests to a root service.

---

## Features

- 🚦 **Multiple Load Balancing Strategies**: Round Robin, Least Connections, and Random
- 🔄 **Dynamic Configuration**: via environment variables or `.env`
- ⚡ **Fast Request Handling**: Gin-based for high concurrency and low latency
- 🛠️ **Modular Architecture**: Clean, extendable codebase
- 📝 **Request Logging & Monitoring**: Built-in request and memory usage logs
- 🏷️ **Centralized Constants**: Easy management of config, error strings, and strategy names
- 🔗 **Root Service Routing Logic**: Requests are forwarded to a designated root microservice; if not possible, returns the server name for possible load transfer

---

## Root Service Forwarding Logic (in Handler)

For each incoming request, the handler:

1. **Tries to send the request to the configured "root microservice"**.
2. **If the root service responds**, its response is returned to the client.
3. **If the root service cannot process the request** (for example, unavailable, error, or any specific application logic),
   - The handler returns the name of the backend server to which the load should be transferred (so an orchestrator or client can reroute or retry).

**This is not a health check or automatic failover mechanism:**  
It is a forwarding logic that always tries the root service first and only returns a server name when forwarding isn't possible.

---

## Example Flow

1. **Client sends a request** to the load balancer.
2. **Handler forwards the request to the root microservice**.
3. - **If root microservice responds:**  
     - The response is returned to the client.
   - **If not:**  
     - The handler returns the backend server name for manual load transfer.

---

## Getting Started

### Prerequisites

- Go 1.18+
- [Gin Web Framework](https://github.com/gin-gonic/gin) (installed via `go mod`)

### Installation

```bash
git clone https://github.com/Shagun-Rajput/Go-Gin-LoadBalancer.git
cd Go-Gin-LoadBalancer
go mod tidy
```

### Configuration

Set the following environment variables or use a `.env` file:

- `APP_PORT`: Port for load balancer (e.g., `8080`)
- `SERVER_LIST`: Comma-separated backend URLs (`http://localhost:9001,http://localhost:9002`)
- `ROOT_SERVICE_URL`: URL of the root microservice to which requests are forwarded
- `LOAD_BALANCING_STRATEGY`: `round_robin`, `least_connections`, or `random`

Example `.env`:

```
APP_PORT=8080
SERVER_LIST=http://localhost:9001,http://localhost:9002
ROOT_SERVICE_URL=http://localhost:9000
LOAD_BALANCING_STRATEGY=round_robin
```

---

## Usage

Start the load balancer:

```bash
go run cmd/main.go
```

Send a request:

```bash
curl http://localhost:8080/your-path
```

The handler will attempt to forward the request to the root microservice. If it can't, it will return a backend server name for manual transfer.

---

## Code Structure

```
.
├── cmd/
│   └── main.go
├── internal/
│   ├── config/
│   │   └── serverconfig.go
│   ├── constants/
│   │   └── constants.go
│   ├── handler/
│   │   └── handler.go  # Root service forwarding logic is here!
│   ├── router/
│   │   └── router.go
│   ├── service/
│   │   ├── service.go
│   │   └── strategy/
│   │       ├── roundrobin.go
│   │       ├── leastconnection.go
│   │       └── random.go
```

- **cmd/main.go**: Entrypoint. Initializes server, loads config, and starts monitoring.
- **internal/config/serverconfig.go**: Loads config from environment.
- **internal/constants/constants.go**: Defines constants for configuration and logic.
- **internal/handler/handler.go**: Handles HTTP requests, implements root service forwarding logic.
- **internal/router/router.go**: Sets up routes and attaches handlers.
- **internal/service/service.go**: Implements strategy selection and request distribution.
- **internal/service/strategy/**: Contains algorithm implementations for each load balancing strategy.

---

## Extending

- **Add new constants:** Update `internal/constants/constants.go`.
- **Add new strategies:** Implement in `internal/service/strategy/` and register in `service.go`.
- **Change forwarding logic:** Edit `internal/handler/handler.go` as needed for your root service routing.

---

## License

MIT

---

Happy load balancing! 🚀
