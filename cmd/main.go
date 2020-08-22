package main

import "fmt"

func getHello() string {
	return "Hello go!"
}

func main() {
	text := getHello()
	fmt.Println(text)
}
