package main

import (
    "fmt"
    "sort"
)

type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) More(i, j int) bool {
	return !s.Less(i, j)
}

func (s byLength) Less(i, j int) bool{
	return len(s[i]) < len(s[j])
}



func main(){
	fruits := []string{"kiwin", "kiwinolada", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}