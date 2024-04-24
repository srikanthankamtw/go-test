package main

import "fmt"

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

func getLanguage(language string) string {
	languageHelloPrefix := englishHelloPrefix
	switch language {
	case "French":
		languageHelloPrefix = frenchHelloPrefix
	case "Spanish":
		languageHelloPrefix = spanishHelloPrefix
	}
	return languageHelloPrefix
}

func main() {
	fmt.Println(Hello("Satya", "Spanish"))
}
