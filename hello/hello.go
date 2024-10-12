package main

import "fmt"

const (
	russian = "Russian"
	french  = "French"

	englishHelloPrefix = "Hello, "
	russianHelloPrefix = "Привет, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case russian:
		prefix = russianHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
