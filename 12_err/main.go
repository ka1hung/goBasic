package main

import (
	"errors"
	"fmt"
)

func divider(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("can't divided by zero")
	}
	return a / b, nil
}
func main() {
	n, err := divider(1.1, 222)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n)
	}
}
