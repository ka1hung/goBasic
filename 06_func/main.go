package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(add(1, 2))

	year, str := today()
	fmt.Println(year, str)

	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := foo(ii...)
	fmt.Println(n)

	ii2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n2 := bar(ii2)
	fmt.Println(n2)
}
func add(a int, b int) int {
	return a + b
}

func today() (int, string) {
	now := time.Now()
	return now.Year(), now.String()
}

func foo(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func bar(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
