package main

import "fmt"

func main() {
	//slice
	slice := []int{1, 2, 3, 4, 5, 6, 7}   //宣告並設值
	fmt.Println("slice[2:5]", slice[2:5]) //指定slice的資料slice[n:m]= slice[n]~slice[m-1]
	fmt.Println("slice[:3]", slice[:3])   // =slice[0:3]
	fmt.Println("slice[3:]", slice[3:])   // =slice[3:len(slice)-1]
	fmt.Println()
	fmt.Println()

	//
	//append
	//
	s1 := []int{}
	//cap:可容許的資料長度 len:資料長度 append:多加一筆資料在最後面
	fmt.Printf("len: %v, cap: %v, data: %v\n", len(s1), cap(s1), s1)

	for i := 0; i < 10; i++ {
		//使用append 如過超過cap值 cap會自動擴展
		s1 = append(s1, i)
		fmt.Printf("len: %v, cap: %v, data: %v\n", len(s1), cap(s1), s1)
	}

	s2 := []int{9999, 11199, 33}
	s1 = append(s1, s2...)
	fmt.Printf("len: %v, cap: %v, data: %v\n", len(s1), cap(s1), s1)
	fmt.Println()
	fmt.Println()
	//
	//copy
	//
	copy(s2, s1)
	fmt.Println(s1)
	copy(s1, []int{0, 1, 2, 3, 4})
	fmt.Println(s1)

	fmt.Println("")
	fmt.Println("")
	//remove the 5st data
	s1 = append(s1[:5], s1[6:]...)
	fmt.Println(s1)

	fmt.Println("")
	fmt.Println("")
}
