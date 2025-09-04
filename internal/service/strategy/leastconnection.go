package strategy

import (
	"sync"
)

var (
	connCount = make(map[string]int)
	connMutex sync.Mutex
)

// Increment connection count for a server
func IncConnection(server string) {
	connMutex.Lock()
	defer connMutex.Unlock()
	connCount[server]++
}

// Decrement connection count for a server
func DecConnection(server string) {
	connMutex.Lock()
	defer connMutex.Unlock()
	if connCount[server] > 0 {
		connCount[server]--
	}
}

func LeastConnections(servers []string) string {
	if len(servers) == 0 {
		return "No available servers"
	}
	connMutex.Lock()
	defer connMutex.Unlock()

	// Initialize connection count for new servers
	for _, s := range servers {
		if _, exists := connCount[s]; !exists {
			connCount[s] = 0
		}
	}

	// Find server with least connections
	minServer := servers[0]
	minConn := connCount[minServer]
	for _, s := range servers[1:] {
		if connCount[s] < minConn {
			minServer = s
			minConn = connCount[s]
		}
	}

	// Do NOT increment here; increment should be done when request is assigned
	return minServer
}
