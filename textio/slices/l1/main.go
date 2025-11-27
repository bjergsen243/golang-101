package sicesl1

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	strArr := [3]string{primary, secondary, tertiary}
	intArr := [3]int{len(primary), len(primary) + len(secondary), len(primary) + len(secondary) + len((tertiary))}
	return strArr, intArr
}
