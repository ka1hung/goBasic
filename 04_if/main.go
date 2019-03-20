package main

import "fmt"

func main() {
	//if else
	condition := false
	if condition {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	//if else-if else
	n := 199
	if n > 100 {
		fmt.Println("n>100")
	} else if n < 100 {
		fmt.Println("n<100")
	} else {
		fmt.Println("n=100")
	}

	//switch-case go會自行煞車，如果不想煞車可以加fallthrough
	switch n {
	case 100:
		fmt.Println("n=100")
		// fallthrough
	case 199:
		fmt.Println("n=199")
	default:
		fmt.Printf("No match number")
	}

	//前面先跑點運算 最後再來判斷也是可以的
	if a, b := 99, 99; a == b {
		fmt.Println("a==b")
	}
}
