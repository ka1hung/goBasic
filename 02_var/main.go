package main

import "fmt"

//宣告全域變數 預設如果沒有設值初始值為0
var (
	i int
	j float64
)

func main() {
	fmt.Printf("i的形態為:%T,數值:%v\n", i, i)
	fmt.Printf("j的形態為:%T,數值:%v\n", j, j)

	//三種宣告方法都可行 但是你會愛上:=
	var a1 int = 100
	var a2 = 100
	a3 := 100 //編譯器自行判斷型別
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)

	//同時宣告兩個變數並設值
	var b, c int = 99, 11
	fmt.Println(b, c)

	//go支援兩變數值直接互換
	b, c = c, b
	fmt.Println(b, c)

	//最後宣告了不用go會生氣(土撥鼠不喜歡浪費)
	var waste int
	fmt.Println(waste) //試著把這行刪掉試試 //waste declared and not used

}
