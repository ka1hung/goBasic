//https://golang.org/ref/spec#Struct_types
package main

import "fmt"

type person struct {
	frist string
	last  string
	fav   []string
}
type employee struct {
	person
	department string
}

func main() {
	//
	//use struct
	//
	p1 := person{
		frist: "kevin",
		last:  "hsiao",
		fav:   []string{"apple", "banana"},
	}
	fmt.Println(p1)
	fmt.Println(p1.frist, p1.last, p1.fav)

	p2 := person{}
	p2.frist = "中山"
	p2.last = "孫"
	p2.fav = []string{"蘋果", "香蕉"}
	fmt.Println(p2)

	fmt.Println("")
	fmt.Println("")
	//
	//use struct array
	//
	ps := []person{}
	ps = append(ps, p1)
	ps = append(ps, p2)
	fmt.Println(ps)

	for _, p := range ps {
		love := ""
		for _, f := range p.fav {
			love += f + " "
		}

		fmt.Printf("%s,%s 喜歡吃 %s\n", p.frist, p.last, love)
	}

	fmt.Println("")
	fmt.Println("")
	//
	//embed struct
	//
	emp := employee{
		person:     p1,
		department: "工程部",
	}
	fmt.Println(emp)
	fmt.Println(emp.frist)
	fmt.Println(emp.person.frist)

	fmt.Println("")
	fmt.Println("")
	//
	//anonymous struct
	//
	s := struct {
		phone   string
		address string
	}{
		phone:   "1234567890",
		address: "台灣",
	}
	fmt.Printf("%v\n", s)
}
