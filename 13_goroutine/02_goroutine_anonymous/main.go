package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println(time.Now().Format("2006-01-02 15:04:05.999999"))
		}
	}()
	for {
		fmt.Println("main loop")
		time.Sleep(time.Second * 10)
	}
}
