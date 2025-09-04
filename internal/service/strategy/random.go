package strategy

import (
	"math/rand"
)

func Random(servers []string) string {
	if len(servers) == 0 {
		return "No available servers"
	}
	// Simple random selection logic
	// Select a random server from the list
	server := servers[rand.Intn(len(servers))]
	return server
}
