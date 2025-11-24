package loopl6

func countConnections(groupSize int) int {
	// ?
	total := 0
	for groupSize > 1 {
		groupSize--
		total += groupSize
	}
	return total
}
