//https://yami.io/golang-interface/ good ref
package main

import "fmt"

//interface{} 可以當做任何的變數
func printer1(data interface{}) {
	fmt.Println(data.(string))
}

func printer2(value interface{}) {
	// 透過型態斷言揭露 interface{} 真正的型態。
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
}
