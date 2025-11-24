package loopl1

func bulkSend(numMessages int) float64 {
	total := 0.0
	for i := 0; i < numMessages; i++ {
		total += 1.0 + (0.01 * float64(i))
	}
	return total
}
