package strategy

import (
	"sync"
)

var (
	rrIndex int
	rrMutex sync.Mutex
)

func RoundRobin(servers []string) string {
	if len(servers) == 0 {
		return "No available servers"
	}
	rrMutex.Lock()
	defer rrMutex.Unlock()
	server := servers[rrIndex%len(servers)]
	rrIndex++
	return server
}
