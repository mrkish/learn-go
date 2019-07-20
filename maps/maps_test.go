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
	dictionary := Dictionary{}
	key := "test"
	want := "this is just a test"

	dictionary.Add(key, want)

	assertValue(t, dictionary, key, want)
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
