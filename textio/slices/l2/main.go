package l2

import (
	"errors"
)

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	if plan == "pro" {
		return messages[:], nil
	}
	if plan == "free" {
		return messages[:2], nil
	}
	return nil, errors.New("unsupported plan")
}
