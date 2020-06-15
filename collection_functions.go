package main

import (
	"fmt"
	"strings"
)

func Index(vs []string, t string) int {
	for i, val := range vs {
		if val == t {
			return i
		}
	}
	return -1
}

func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

func Any(vs []string, f func(string) bool) bool {
	for _, val := range vs {
		if f(val) {
			return true
		}
	}
	return false
}

func All(vs []string, f func(string) bool) bool {
	for _, val := range vs {
		if !f(val) {
			return false
		}
	}
	return true
}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, val := range vs {
		vsm[i] = f(val)
	}
	return vsm
}

func main() {
	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "pear"))

	fmt.Println(Include(strs, "grape"))

	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(All(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	fmt.Println(Map(strs, strings.ToUpper))
}
