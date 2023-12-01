package main

import "fmt"

const spanish = "Spanish"
const english = "English"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(name string, language ...string) string {
	lang := english
	if len(language) > 0 {
		lang = language[0]
	}
	if name == "" {
		name = "World"
	}
	if lang == spanish {
		return spanishHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
