package constants

const (
	StrategyRoundRobin         = "round_robin"
	StrategyLeastConnections   = "least_connections"
	StrategyRandom             = "random"
	StrategyWeightedRoundRobin = "weighted_round_robin"
	EnvAppPort                 = "APP_PORT"
	EnvServerList              = "SERVER_LIST"
	EnvLoadBalancingStrategy   = "LOAD_BALANCING_STRATEGY"
)
