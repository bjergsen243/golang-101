package l11

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
	res := make([]float64, 0)
	for _, v := range costs {
		if day == v.day {
			res = append(res, v.value)
		}
	}
	return res
}
