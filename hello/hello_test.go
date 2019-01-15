package hello

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Peoria", "French")
		want := "Bonjour, Peoria"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in German", func(t *testing.T) {
		got := Hello("Galileo", "German")
		want := "Guten tag, Galileo"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Hindi", func(t *testing.T) {
		got := Hello("Devioush", "Hindi")
		want := "Namste, Devioush"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Italian", func(t *testing.T) {
		got := Hello("Gapeto", "Italian")
		want := "Ciao, Gapeto"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Japanese", func(t *testing.T) {
		got := Hello("Yusake", "Japanese")
		want := "Konnichiwa, Yusake"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Persian", func(t *testing.T) {
		got := Hello("Halo", "Persian")
		want := "Salaam, Halo"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
}

// https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go
func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hello("test", "English")
	}
}

// https://blog.golang.org/examples
func ExampleHello() {
	fmt.Println(Hello("test", "English"))
	// Output: Hello, test
}
