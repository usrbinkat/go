package hello

import "fmt"

var name = ""
var greet = "Hello"

func Hello(greet string, name string) string {

	greeting := greet + ", "

	if greet == "" {
		greet = "Hello"
	}

	if name == "" {
		name = "werld"
	}

	return greeting + name
}

func main() {
	fmt.Println(Hello(greet, name))
}
