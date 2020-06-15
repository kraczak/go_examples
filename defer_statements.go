package main

import (
	"fmt"
	"time"
)

func printHelloWorld(){
	defer fmt.Print("World")
	fmt.Println("Hello ")
}


func returnThenSleep(number int) int{
	defer sleepFunc(number)
	return number
}

func sleepFunc(seconds int){
	time.Sleep(5000 * time.Millisecond)
	fmt.Printf("I slept for %v seconds\n", seconds)
}

func main()  {
	printHelloWorld()
	five := returnThenSleep(5)
	fmt.Println(five)
}