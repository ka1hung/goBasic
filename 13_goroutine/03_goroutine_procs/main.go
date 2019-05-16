package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("number of cpu in this pc:", runtime.NumCPU())
	runtime.GOMAXPROCS(1) //limit only 1 cpu in use; default is max cpu number(8)
	go task()
	go task()
	go task()
	go task()
	go task()
	go task()
	go task()
	go task()
	go task()
	fmt.Println("number of gorountine are runing:", runtime.NumGoroutine())

	for {
		time.Sleep(time.Second * 10)
	}

}

func task() {

	for {
		time.Sleep(time.Second)
	}

}
