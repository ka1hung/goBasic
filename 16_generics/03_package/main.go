package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

type Data struct {
	ID   int
	Name string
}

// constraints.Ordered is useful
func sumVal[T constraints.Ordered](vals []T) T {
	var sum T
	for i := range vals {
		sum += vals[i]
	}
	return sum
}

func main() {
	//
	//constraints package
	//
	fmt.Println(sumVal([]int{1, 2, 3}))           //6
	fmt.Println(sumVal([]float64{1.1, 2.1, 3.1})) //6.3

	//
	//slices package
	//
	//index find the index of the same value in slice
	//return -1 if not exist
	v := []int{36, 2, 1}
	fmt.Println(slices.Index(v, 2)) //1
	fmt.Println(slices.Index(v, 4)) //-1

	//clone
	v = []int{1, 2, 3}
	v1 := v //do assign
	v1[0] = 99
	fmt.Println(v[0], v1[0]) //99 99

	v = []int{1, 2, 3}
	vClone := slices.Clone(v) //do clone
	vClone[0] = 99
	fmt.Println(v[0], vClone[0]) //1 99

	ds := []Data{}
	ds = append(ds, Data{ID: 1, Name: "aaa"})
	ds = append(ds, Data{ID: 2, Name: "bbb"})
	ds = append(ds, Data{ID: 3, Name: "ccc"})
	dsCopy := slices.Clone(ds)
	fmt.Println(dsCopy)

	//insert
	v = []int{1, 2, 3}
	index := 0
	fmt.Println(slices.Insert(v, index, -1), v)                     //[-1 1 2 3] [1 2 3]  *v=slices.Insert(v, 0, -1)*
	fmt.Println(slices.Insert(ds, 0, Data{ID: 0, Name: "xxx"}), ds) //[{0 xxx} {1 aaa} {2 bbb} {3 ccc}] [{0 xxx} {1 aaa} {2 bbb}]

	//delete
	v = []int{1, 2, 3}
	fmt.Println(slices.Delete(v, 0, 1), v) //[2 3] [2 3 3]
	ds = []Data{}
	ds = append(ds, Data{ID: 1, Name: "aaa"})
	ds = append(ds, Data{ID: 2, Name: "bbb"})
	ds = append(ds, Data{ID: 3, Name: "ccc"})
	fmt.Println(slices.Delete(ds, 1, 2), ds) //[{1 aaa} {3 ccc}] [{1 aaa} {3 ccc} {3 ccc}]

	//equal
	fmt.Println(slices.Equal([]int{1, 2, 3}, []int{1, 2, 3}))    //true
	fmt.Println(slices.Equal([]int{1, 4, 3}, []int{1, 2, 3}))    //false
	fmt.Println(slices.Equal([]int{1, 2, 3}, []int{1, 2, 3, 4})) //false

	//contains
	fmt.Println(slices.Contains([]int{1, 2, 3}, 1)) //true
	fmt.Println(slices.Contains([]int{1, 2, 3}, 9)) //false

	ds = []Data{{ID: 1, Name: "aaa"}, {ID: 2, Name: "bbb"}}
	fmt.Println(slices.Contains(ds, Data{ID: 1, Name: "aaa"})) //true
	fmt.Println(slices.Contains(ds, Data{ID: 1}))              //false

	//
	//maps package
	//
	m := map[int]string{
		123: "123",
		456: "456",
		112: "112",
		111: "111",
		110: "110",
	}
	m1 := m //do assign
	m1[123] = "999"
	fmt.Println(m[123], m1[123]) //999 999

	fmt.Println(maps.Keys(m))   // random key
	fmt.Println(maps.Values(m)) // random value

	m = map[int]string{
		123: "123",
		456: "456",
	}
	m2 := maps.Clone(m) //do clone
	m2[123] = "999"
	fmt.Println(m2[123], m[123]) //999 123

	m2 = maps.Clone(m)                    //do clone
	fmt.Println(maps.Equal(m2, m), m2, m) // true
	m3 := map[int]string{
		789: "baz",
		123: "aaa",
	}

	maps.Copy(m3, m)                      //more like add, over write value if key exist
	fmt.Println(maps.Equal(m3, m), m3, m) // false
}
