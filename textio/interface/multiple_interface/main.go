package multipleinterface

import "fmt"

func (e email) cost() int {
	eachChar := 5
	if e.isSubscribed {
		eachChar = 2
	}
	return eachChar * len(e.body)
}

func (e email) format() string {
	subText := "Not Subscribed"
	if e.isSubscribed {
		subText = "Subscribed"
	}

	return fmt.Sprintf("'%s' | %s", e.body, subText)
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}
