package main

import "fmt"

func sumInt(vals []int) int {
	var sum int
	for i := range vals {
		sum += vals[i]
	}
	return sum
}
func sumFloat64(vals []float64) float64 {
	var sum float64
	for i := range vals {
		sum += vals[i]
	}
	return sum
}

func sumNum[T int | float64](vals []T) T {
	var sum T
	for i := range vals {
		sum += vals[i]
	}
	return sum
}

func add[T int | int32](a, b T) T {
	return a + b
}

func addUnderlyingType[T ~int | int32](a, b T) T {
	return a + b
}

type INT int

func main() {
	// before 1.18
	fmt.Println(sumInt([]int{1, 2, 3}))
	fmt.Println(sumFloat64([]float64{1, 2, 3}))

	// after 1.18
	fmt.Println(sumNum([]int{1, 2, 3}))
	fmt.Println(sumNum([]float64{1.1, 2, 3}))

	// ~ UnderlyingType
	var i INT = 1
	fmt.Println(i)
	// fmt.Println(add(i, 1)) //error
	fmt.Println(addUnderlyingType(i, i))
}
