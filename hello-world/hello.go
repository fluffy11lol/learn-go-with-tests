package main

import "fmt"

const french = "French"
const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const frenchHelloPrefix = "Bonjour, "
const spanishHelloPrefix = "Hola, "

func Hello(s string, language ...string) string {
	if s == "" {
		s = "World"
	}
	return greetingPrefix(language...) + s + "!"
}
func greetingPrefix(language ...string) string {
	if len(language) == 0 {
		return englishHelloPrefix
	}
	switch language[0] {
	case spanish:
		return spanishHelloPrefix
	case french:
		return frenchHelloPrefix
	default:
		return englishHelloPrefix
	}
}
func main() {
	fmt.Println(Hello("s"))
}
