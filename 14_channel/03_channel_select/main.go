package main

import (
	"fmt"
	"time"
)

func counter(n int, c chan int, quit chan bool) {
	for i := 0; i < n; i++ {
		time.Sleep(time.Second)
		c <- i
	}
	quit <- true
}

func main() {
	c := make(chan int)
	quit := make(chan bool)
	go counter(5, c, quit)

	num := 999
	for {
		select {
		case num = <-c:
			fmt.Println(num)
		case <-quit:
			fmt.Println("quit")
			return
		}
	}

}
