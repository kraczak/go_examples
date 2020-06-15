package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	var sum int = 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	dict := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range dict {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range dict {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		// c is Unicode code point
		fmt.Println(i, c)
	}
}
