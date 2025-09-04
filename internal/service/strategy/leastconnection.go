package strategy

import (
	"sync"
)

var (
	connCount = make(map[string]int)
	connMutex sync.Mutex
)

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

	// Simulate incrementing connection count
	connCount[minServer]++

	return minServer
}
