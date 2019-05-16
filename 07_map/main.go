package main

import "fmt"

func main() {
	//宣告
	var m = make(map[int]string)

	//add
	m[1] = "Monday"
	m[2] = "Tuesday"
	m[3] = "Wednesday"
	m[4] = "Thursday"
	m[5] = "Friday"
	m[6] = "Saturday"
	m[7] = "Sunday"

	//read
	fmt.Println(m[1])

	//check
	day, ok := m[1]
	fmt.Println(day, ok)

	day, ok = m[9]
	fmt.Println(day, ok)

	//range *注意key不會排序*
	for key := range m {
		fmt.Print(m[key], " ")
	}

	//delete
	delete(m, 1)
	fmt.Println(m[1])

	//clear all
	fmt.Println(m)
	m = make(map[int]string)
	fmt.Println(m)
}
