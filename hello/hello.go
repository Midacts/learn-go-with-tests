package hello

const french = "French"
const frenchHelloPrefix = "Bonjour, "
const german = "German"
const germanHelloPrefix = "Guten tag, "
const hindi = "Hindi"
const hindiHelloPrefix = "Namste, "
const italian = "Italian"
const italianHelloPrefix = "Ciao, "
const japanese = "Japanese"
const japaneseHelloPrefix = "Konnichiwa, "
const persian = "Persian"
const persianHelloPrefix = "Salaam, "
const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "
const englishHelloPrefix = "Hello, "

// Hello says hello to the requestor in the specified language
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// greetingPrefix stores chooses the hello in the specified language
func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case german:
		prefix = germanHelloPrefix
	case hindi:
		prefix = hindiHelloPrefix
	case italian:
		prefix = italianHelloPrefix
	case japanese:
		prefix = japaneseHelloPrefix
	case persian:
		prefix = persianHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

/*
func main() {
	fmt.Println(Hello("world", ""))
}
*/
