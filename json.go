package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	boolB, _ := json.Marshal(true)
	fmt.Println(string(boolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	floatB, _ := json.Marshal(2.34)
	fmt.Println(string(floatB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	sliceD := []string{"apple", "banana", "kiwi"}
	sliceB, _ := json.Marshal(sliceD)
	fmt.Println(string(sliceB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &response1{Page: 1, Fruits: []string{"apple", "kiwi", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &response2{Page: 1, Fruits: []string{"apple", "kiwi", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	str := `{"num": 6.13, "strs": ["a","b"]}`
	byt := []byte(str)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1, strs)

	str = `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)

	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
