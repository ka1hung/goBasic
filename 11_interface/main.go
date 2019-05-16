//https://yami.io/golang-interface/ good ref
package main

import "fmt"

//interface{} 可以當做任何的變數
func printer1(data interface{}) {
	fmt.Println(data.(string))
}

func printer2(value interface{}) {
	switch v := value.(type) {
	case string:
		fmt.Println("value is string" + v)
	case int:
		fmt.Printf("value is int %d\n", v)
	default:
		fmt.Println("Not support this type")
	}
}

func main() {
	printer1("hello")
	printer2("hello")
	printer2(123)
	printer2(99.99)

	mm := myMath{num1: 99.9, num2: 199}
	fmt.Println(doMath(mm))

	ym := yourMath{num1: 100, num2: 199}
	fmt.Println(doMath(ym))
}

type myMath struct {
	num1 float64
	num2 float64
}

func (mm myMath) add() float64 {
	return mm.num1 + mm.num2
}

type yourMath struct {
	num1 int64
	num2 int64
}

func (ym yourMath) add() float64 {
	return float64(ym.num1 + ym.num2)
}

type math interface {
	add() float64
}

func doMath(m math) float64 {
	return m.add()
}
