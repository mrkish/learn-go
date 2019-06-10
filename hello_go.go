package main

func main () {
}

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

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
