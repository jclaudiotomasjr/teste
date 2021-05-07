package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"d", "b", "a", "c"}
	sort.Strings(strs)
	fmt.Println("Strings: ", strs)

	ints := []int{14, 8, 9, 3}
	sort.Ints(ints)
	fmt.Println("ints: ", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Ordenado: ", s)
}
