package main

import (
	"fmt"
	"sort"
)

type Data struct {
	ID   int
	Name string
}

func main() {
	//sort int
	s1 := []int{8, 2, 6, 3, 1, 4}
	sort.Ints(s1)
	fmt.Println(s1)

	//sot int reverse
	sort.Sort(sort.Reverse(sort.IntSlice(s1)))
	fmt.Println(s1)

	//sort string
	s2 := []string{"aaa", "bbb", "6", "3", "1", "4"}
	sort.Strings(s2)
	fmt.Println(s2)

	//sot string reverse
	sort.Sort(sort.Reverse(sort.StringSlice(s2)))
	fmt.Println(s2)

	//sort float64
	s3 := []float64{1, 1.2, 0, -1.9, -82.333, 99.11}
	sort.Float64s(s3)
	fmt.Println(s3)

	//sot float64 reverse
	sort.Sort(sort.Reverse(sort.Float64Slice(s3)))
	fmt.Println(s3)

	//sort struct
	ds := []Data{}
	ds = append(ds, Data{ID: 49, Name: "kevin"})
	ds = append(ds, Data{ID: 11, Name: "peter"})
	ds = append(ds, Data{ID: 11, Name: "mary"})
	ds = append(ds, Data{ID: 11, Name: "adon"})
	ds = append(ds, Data{ID: 15, Name: "lily"})

	//sort by id
	sort.Slice(ds, func(i, j int) bool {
		return ds[i].ID < ds[j].ID
	})
	fmt.Println(ds)

	//sort by Name
	sort.Slice(ds, func(i, j int) bool {
		return ds[i].Name < ds[j].Name
	})
	fmt.Println(ds)

	//sort by id reverse
	sort.Slice(ds, func(i, j int) bool {
		return ds[i].ID > ds[j].ID
	})
	fmt.Println(ds)
}
