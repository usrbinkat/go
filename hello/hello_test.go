package hello

import "testing"

func TestHello(t *testing.T) {

	t.Run("Say hello in English to default name 'werld'", func(t *testing.T) {
		expect := "Hello, werld"
		result := Hello("english", "")
		checkHelloMsg(t, expect, result)
	})

	t.Run("Say hello in French to Fancy", func(t *testing.T) {
		expect := "Bonjour, Fancy"
		result := Hello("french", "Fancy")
		checkHelloMsg(t, expect, result)
	})

	t.Run("Say hello in Spanish to mundo", func(t *testing.T) {
		expect := "Hola, mundo"
		result := Hello("spanish", "mundo")
		checkHelloMsg(t, expect, result)
	})
}

func checkHelloMsg(t testing.TB, expect, result string) {
	t.Helper()
	if expect != result {
		msg := "\n Expect: %q \n Result: %q"
		t.Errorf(msg, expect, result)
	}
}
