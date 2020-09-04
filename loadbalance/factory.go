package loadbalance

const (
	Random = iota
	RoundRobin
	Weight
	Hash
)

func LoadBalanceFactory(lbType int) LoadBalance {
	switch lbType {
	case Random:
		return new(RandomBalance)
	case RoundRobin:
		return new(RoundRobinBalance)
	case Weight:
		return new(WeightRoundRobinBalance)
	case Hash:
		return new(HashBalance)
	default:
		return new(RoundRobinBalance)
	}
}
