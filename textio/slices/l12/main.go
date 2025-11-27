package l12

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i, msg := range msg {
		for _, bad := range badWords {
			if msg == bad {
				return i
			}
		}
	}
	return -1
}
