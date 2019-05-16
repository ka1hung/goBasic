package main

import "fmt"

type myMath struct {
	num1 float64
	num2 float64
}

func NewmyMath(n1, n2 float64) *myMath {
	return &myMath{
		num1: n1,
		num2: n2,
	}
}

func (mm *myMath) add() float64 {
	return mm.num1 + mm.num2
}

//使用指標 隨你揉捏
func (mm *myMath) set(n1, n2 float64) {
	mm.num1 = n1
	mm.num2 = n2
}

//不用指標的話 只能read only
func (mm myMath) setFake(n1, n2 float64) {
	mm.num1 = n1
	mm.num2 = n2
}

func (mm myMath) toString() string {
	s := fmt.Sprintf("num1:%v, num2:%v", mm.num1, mm.num2)
	return s
}

func main() {

	m := myMathNew(0.99, 1.99)
	fmt.Println(m.toString())
	fmt.Printf("sum:%v\n", m.add())

	m.set(100, 55)
	fmt.Println(m.toString())
	m.setFake(1010, 5511)
	fmt.Println(m.toString())
}
