package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Dmitrii", "")
		want := "Hello, Dmitrii"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Russian", func(t *testing.T) {
		got := Hello("Дмитрий", "Russian")
		want := "Привет, Дмитрий"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Macron", "French")
		want := "Bonjour, Macron"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// Отмечаем что это функция-хелпер, и если тест упадёт, он не будем показывать строку кода направляющую в эту функцию,
	// А будет направлять на строчку функции которая вызвала эту функцию
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
