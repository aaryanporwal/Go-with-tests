package main

import "fmt"

const englishHelloPrefix = "Hello, "
const suffix = "!"

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name + suffix
}

func main() {
	fmt.Println(Hello("Aaryan"))
}
