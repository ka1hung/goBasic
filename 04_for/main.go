package main

import (
	"fmt"
	"time"
)

func main() {
	ns := []int{99, 22, 44}
	fmt.Println(ns)
	//tranditional for
	for i := 0; i < len(ns); i++ {
		fmt.Printf("ns[%d]=%d\n", i, ns[i])
	}

	//foreach
	for n, i := range ns {
		fmt.Println("n:", n, "; i:", i)
	}
	for n := range ns {
		fmt.Println("n:", n)
	}
	for _, i := range ns {
		fmt.Println("i:", i)
	}

	//loop
	n := 0
	for {
		time.Sleep(1 * time.Second)
		n++
		fmt.Println(n)
	}
}
