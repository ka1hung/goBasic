package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go func() {
		for {
			time.Sleep(time.Second)
			c <- time.Now().Format("2006-01-02 15:04:05.999999")
		}
	}()
	for {
		fmt.Println(<-c)
	}
}
