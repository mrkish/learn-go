package main

func main() {

	newDoc = Doctor{
		"Bob",
		"STL",
		"Sue",
	}

}

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

type Doctor struct {
	name     string
	location string
	company  string
	patients string
}

type company interface {
	getCompanyAddress(companyName string) string
}

func (c *company) getCompanyAddress(companyName string) string {
	return companyName
}

func Hello(name, language string) string {
	if len(name) == 0 {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
