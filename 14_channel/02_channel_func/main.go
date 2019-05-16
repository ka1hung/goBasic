package main

import (
	"fmt"
	"time"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		time.Sleep(time.Second)
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int)
	go fibonacci(20, c)
	for i := range c {
		fmt.Println(i)
	}
}
