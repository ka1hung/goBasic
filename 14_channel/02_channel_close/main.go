package main

import (
	"fmt"
	"time"
)

func counter(n int, c chan int) {
	for i := 0; i < n; i++ {
		time.Sleep(time.Second)
		c <- i
	}
	close(c)
}

func main() {
	c := make(chan int)
	go counter(20, c)
	for i := range c {
		fmt.Println(i)
	}
}
