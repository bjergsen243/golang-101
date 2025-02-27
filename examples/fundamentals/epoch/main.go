package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("now:", now)

	fmt.Println("now in unix:", now.Unix())
	fmt.Println("now ms:", now.UnixMilli())
	fmt.Println("now ns:", now.UnixNano())

	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.Unix()))
}
