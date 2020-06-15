package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f := openFile("defer.txt")
	readFile(3, f)
	defer closeFile(f)
	writeFile(f)
}

func openFile(p string) *os.File {
	fmt.Println("Opening file", p)
	f, err := os.Open(p)
	check(err)
	return f
}

func readFile(byteNo int, f *os.File) {
	bytes := make([]byte, byteNo)
	realBytesNo, err := f.Read(bytes)
	check(err)
	fmt.Println("File contains:", string(bytes[:realBytesNo]))
}

func createFile(path string) *os.File {
	fmt.Println("Creating new file")
	f, err := os.Create(path)
	check(err)
	return f
}

func closeFile(f *os.File) {
	fmt.Println("Closing file")
	err := f.Close()
	check(err)
}

func writeFile(f *os.File) {
	fmt.Println("Writing to file")
	fmt.Fprintln(f, "Data")
}
