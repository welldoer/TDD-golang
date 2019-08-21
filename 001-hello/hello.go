package main

import "fmt"

func Hello() (message string) {
	return "Hello, World"
}

func main() {
	fmt.Println(Hello())
}
