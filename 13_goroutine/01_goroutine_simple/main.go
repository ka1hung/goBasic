package main

import (
	"fmt"
	"time"
)

func doSomething(name string) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Printf("%v count %v\n", name, i)
	}
}
func main() {
	go doSomething("hi")
	go doSomething("hello")
	time.Sleep(1 * time.Second) //wait to see result
}
