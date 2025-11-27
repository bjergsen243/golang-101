package l6

func getMessageCosts(messages []string) []float64 {
	cost := make([]float64, len(messages))
	for i, msg := range messages {
		cost[i] = float64(len(msg)) * 0.01
	}
	return cost
}
