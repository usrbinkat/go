package hello

import "testing"

func TestHello(t *testing.T) {

	t.Run("Say hello to Fancy", func(t *testing.T) {
		expect := "Hello, Fancy"
		result := Hello("Hello", "Fancy")
		checkHelloMsg(t, expect, result)
	})

	t.Run("Say hello to Werld", func(t *testing.T) {
		expect := "Hola, mundo"
		result := Hello("Hola", "mundo")
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
