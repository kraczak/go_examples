package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println(strs)

	ints := []int{9, 3, 1, 2, 5, 4, 6}
	sort.Ints(ints)
	fmt.Println(ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println(s)
}
