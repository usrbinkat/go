package hello

import "fmt"

var name = ""
var language = ""
var greet = ""

func main() {
	fmt.Println(Hello(language, name))
}

func Hello(language string, name string) string {

	if name == "" {
		name = "werld"
	}
	
	greeting := greetPrefixSelect(language) + ", " + name
	return greeting
}

func greetPrefixSelect(language string) (prefix string) {

	switch language {
	case "english":
		greet = "Hello"
	case "spanish":
		greet = "Hola"
	case "french":
		greet = "Bonjour"
	}
    return greet
}