package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("know word", func(t *testing.T) {

		got, err := dictionary.Search("test")
		want := "this is just a test"
		if err != nil {
			t.Fatal(err)
		}
		assertString(t, got, want)
	})
	t.Run("unknow word", func(t *testing.T) {
		_, got := dictionary.Search("unknow")

		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {

		dictionary := Dictionary{}

		word := "test"
		definition := "this is just a test"
		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})
	t.Run("existing word", func(t *testing.T) {

		word := "test"
		definition := "this is just a test"

		dictrionary := Dictionary{word: definition}

		err := dictrionary.Add(word, definition)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictrionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("exiting word", func(t *testing.T) {

		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definiton"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})
	t.Run("new word", func(t *testing.T) {

		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}
		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {

	word := "test"
	definition := "this is just a test"
	dictionary := Dictionary{word: definition}
	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrNotFound {

		t.Errorf("Expected %q to be deleted", word)
	}

}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)

	}

	assertString(t, got, definition)

}

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q got wanted %q", got, want)
	}

}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q got wanted %q", got, want)
	}

}
