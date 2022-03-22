package main

import (
	"fmt"
	"sort"

	"github.com/facette/natsort"
)

type Data struct {
	ID   string
	Name string
}

func main() {
	//sort string
	s2 := []string{"Device2", "Device11", "Device1", "Device22", "Device13", "Device3"}
	sort.Strings(s2)
	fmt.Println(s2)

	//natsort string
	s2 = []string{"Device2", "Device11", "Device1", "Device22", "Device13", "Device3"}
	natsort.Sort(s2)
	fmt.Println(s2)

	//natsort struct slice
	ds := []Data{}
	ds = append(ds, Data{ID: "2", Name: "kevin"})
	ds = append(ds, Data{ID: "11", Name: "peter"})
	ds = append(ds, Data{ID: "1", Name: "mary"})
	ds = append(ds, Data{ID: "3", Name: "adon"})
	ds = append(ds, Data{ID: "12", Name: "lily"})
	//sort by id
	sort.Slice(ds, func(i, j int) bool {
		return natsort.Compare(ds[i].ID, ds[j].ID)
	})

	fmt.Println(ds)
}
