package main

import "fmt"

func main() {
	//show address of a
	a := 100
	fmt.Printf("&a:%v\n", &a)

	//var pointer p to a
	p := &a // var p *int = &a
	fmt.Printf("p: type:%T *p:%v p:%v\n", p, *p, p)

	//use pointer for slice
	ns := []int{100, 200, 300}
	ps := &ns
	fmt.Printf("ps: type:%T *ps:%v ps:%v\n", ps, *ps, ps)
	pps := *ps
	for i, v := range pps {
		fmt.Println(i, v)
	}

	//pass pointer to func
	a, b := 1, 2
	changeVal(a, b)
	fmt.Println("after changeVal:", a, b)
	changeRef(&a, &b)
	fmt.Println("after changeRef:", a, b)
}
func changeVal(a, b int) {
	a, b = b, a
	fmt.Println("in changeVal:", a, b)
}
func changeRef(a, b *int) {
	*a, *b = *b, *a
	fmt.Println("in changeVal:", *a, *b)
}
