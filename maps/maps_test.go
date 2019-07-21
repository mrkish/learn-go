package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("gooble")

		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "test"
		want := "this is just a test"

		err := dictionary.Add(key, want)

		assertError(t, err, nil)
		assertValue(t, dictionary, key, want)
	})

	t.Run("existing word", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		dictionary := Dictionary{key: value}

		err := dictionary.Add(key, value)

		assertError(t, err, ErrWordExists)
		assertValue(t, dictionary, key, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		newValue := "this is another test"
		dict := Dictionary{key: value}

		err := dict.Update(key, newValue)

		assertError(t, err, nil)
		assertValue(t, dict, key, newValue)
	})

	t.Run("new word", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		dict := Dictionary{}

		err := dict.Update(key, value)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	key := "test"
	value := "this is just a test"
	dict := Dictionary{key: value}

	dict.Delete(key)
	_, err := dict.Search(key)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", key)
	}
}

func assertValue(t *testing.T, d Dictionary, key, value string) {
	t.Helper()
	got, err := d.Search(key)

	if err != nil {
		t.Fatal("didn't find expected value:", err)
	}

	if value != got {
		t.Errorf("got %q want %q", got, value)
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {

		t.Errorf("got %q want %q, given %q", got, want, "test")
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
