package main

import (
	"fmt"
)

// custom *no limit*
type Mumber interface {
	int | float64
}

func sumVal[T Mumber](vals []T) T {
	var sum T
	for i := range vals {
		sum += vals[i]
	}
	return sum
}

// any *can't do calculate or compare *
func printAll[T any](vals []T) {
	fmt.Println(vals)
}

// comparable *== !=*
func equal[T comparable](a, b T) bool {
	return a == b
}
func main() {
	//1.18 generics
	fmt.Println(sumVal([]int{1, 2, 3}))
	fmt.Println(sumVal([]float64{1, 2, 3}))
}
