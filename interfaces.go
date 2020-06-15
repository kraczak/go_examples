package main

import (
	"fmt"
	"math"
	"reflect"
)

func getType(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	}
	return t.Name()
}

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2 * (r.height + r.width)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(getType(g), g)
	fmt.Println("area:", g.area())
	fmt.Println("perim:", g.perim())
}

func main() {
	var side float64 = 4
	r := rect{width: side, height: side}
	c := circle{radius: side / 2}
	measure(r)
	measure(c)
	fmt.Println(r.area()/c.area() == 4/math.Pi)
}
