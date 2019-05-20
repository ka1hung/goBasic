package main

import (
	"fmt"
	"log"
	"strconv"
)

//宣告全域變數 如果沒有設值 初始值為0
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

	//變數型態轉換(強轉型)
	var f float64 = 9.99
	fmt.Printf("型態:%T,數值:%v\n", f, f)
	var i int = int(f)
	fmt.Printf("型態:%T,數值:%v\n", i, i)

	// var s = string(f) //float64不能強轉型字串
	var s = string(i)
	fmt.Printf("型態:%T,數值:%v\n", s, s)

	//如果需要轉字串最好使用strconv
	s = strconv.Itoa(i) //int到string
	fmt.Printf("型態:%T,數值:%v\n", s, s)

	//float64轉字串
	s = strconv.FormatFloat(f, 'f', 3, 64)
	fmt.Printf("型態:%T,數值:%v\n", s, s)

	//字串轉float64
	f, err := strconv.ParseFloat(s, 64) //string to float
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("型態:%T,數值:%v\n", f, f)

	//string到int64
	num, err := strconv.ParseInt("1011100111", 2, 64)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("型態:%T,數值:%v\n", num, num)

}
