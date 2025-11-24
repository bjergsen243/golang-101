package loopl2

func maxMessages(thresh int) int {
	total := 100
	count := 0
	for i := 1; total <= thresh; i++ {
		total += 100 + 1*i
		count++
	}
	return int(count)
}
