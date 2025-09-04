# Go Gin Load Balancer

A high-performance, pluggable HTTP load balancer implemented in Go using the Gin framework. Easily route incoming requests across multiple backend servers using configurable load balancing strategies.

## Features

- 🚦 Supports Round Robin, Least Connections, and Random strategies
- 🔄 Dynamic configuration via environment variables
- ⚡ Fast, concurrent request handling with [Gin](https://github.com/gin-gonic/gin)
- 🛠️ Modular codebase for easy extension
- 📝 Request logging and memory usage monitoring

## Getting Started

### Prerequisites

- Go 1.18+
- [Gin Web Framework](https://github.com/gin-gonic/gin) (installed via `go mod`)

### Installation

Clone the repository:

```bash
git clone https://github.com/Shagun-Rajput/Go-Gin-LoadBalancer.git
cd Go-Gin-LoadBalancer
go mod tidy
```

### Configuration

Set the following environment variables (directly or via `.env`):

- `APP_PORT`: The port for the load balancer to listen on (e.g., `8080`)
- `SERVER_LIST`: Comma-separated list of backend server URLs (e.g., `http://localhost:9001,http://localhost:9002`)
- `LOAD_BALANCING_STRATEGY`: Strategy to use (`round_robin`, `least_connections`, `random`)

Example `.env`:

```
APP_PORT=8080
SERVER_LIST=http://localhost:9001,http://localhost:9002
LOAD_BALANCING_STRATEGY=round_robin
```

### Usage

Start the load balancer:

```bash
go run cmd/main.go
```

Send a request:

```bash
curl http://localhost:8080/your-path
```

The request will be routed to one of the backend servers, according to your chosen strategy.

## Code Structure

```
.
├── cmd/
│   └── main.go
├── internal/
│   ├── config/
│   │   └── serverconfig.go
│   ├── handler/
│   │   └── handler.go
│   ├── router/
│   │   └── router.go
│   └── service/
│       ├── service.go
│       └── strategy/
│           ├── roundrobin.go
│           ├── leastconnection.go
│           └── random.go
```

- **cmd/main.go**: Entrypoint, initializes server, loads config, and starts memory monitoring.
- **internal/config/serverconfig.go**: Loads config from environment.
- **internal/router/router.go**: Sets up request routes.
- **internal/handler/handler.go**: Handles HTTP requests and proxies them.
- **internal/service/service.go**: Implements strategy selection logic.
- **internal/service/strategy/**: Contains load balancing algorithms.

## Extending

Add new strategies by implementing them in `internal/service/strategy/`. Update the switch in `service.go` to incorporate your strategy.

## License

MIT

---

Happy load balancing! 🚀
