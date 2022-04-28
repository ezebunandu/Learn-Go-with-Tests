package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}
	prefix := englishHelloPrefix
	if language == "Spanish" {
		prefix = spanishHelloPrefix
	}
	return prefix + name
}
func main() {
	fmt.Println(Hello("", "world"))
}
