package strategy


func RoundRobin(servers []string) string {
	if len(servers) == 0 {
		return "No available servers"
	}
	// Simple round-robin logic
	server := servers[0]
	servers = append(servers[1:], server)
	return "Request sent to " + server
}