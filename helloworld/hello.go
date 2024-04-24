package main

import (
	"fmt"
	"github.com/srikanthankamtw/go-test/integers"
)

const (
	englishHelloPrefix = "Hello"
	spanishHelloPrefix = "Hola"
	frenchHelloPrefix  = "Bonjour"
	spaceDelimiter     = " "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return getLanguage(language) + spaceDelimiter + name
}

func getLanguage(language string) (languageHelloPrefix string) {
	switch language {
	case "French":
		languageHelloPrefix = frenchHelloPrefix
	case "Spanish":
		languageHelloPrefix = spanishHelloPrefix
	default:
		languageHelloPrefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Satya", "Spanish"))
	fmt.Println(integers.Add(1, 2))
}
