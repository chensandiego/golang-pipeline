package main

import "fmt"

var version = "dev"

func main() {
	fmt.Println("Version: ", version)
	fmt.Println(hello())
}

func hello() string {
	return "Hello Golang"
}
