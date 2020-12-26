package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	polishHelloPrefix  = "Witaj, "
)

func main() {
	fmt.Println(Hello("world", ""))
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "French":
		prefix = frenchHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	case "Polish":
		prefix = polishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
