package main

import (
	"fmt"
	"os"
)

func main() {
	// f := createFile("defer.txt")
	// writeFile("napis do pliku", f)
	// readFile(f)
	f := openFile("defer.txt")
	readFile(f)
	defer closeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating file", p)
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(data string, f *os.File) {
	fmt.Printf("writing %v to %v\n", data, f.Name())
	fmt.Fprint(f, data)
}

func closeFile(f *os.File) {
	fmt.Printf("closing file %v\n", f)
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func openFile(p string) *os.File {
	f, err := os.Open(p)
	if err != nil {
		panic(err)
	}
	return f
}

func readFile(f *os.File) {
	bytes := make([]byte, 10)
	readBytes, err := f.Read(bytes)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes[:readBytes]))
}
