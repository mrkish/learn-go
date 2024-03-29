package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"}},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		// struct literals (previous to declaring Person and Profile structs in test file)
		// {
		// 	"Struct with nested structs",
		// 	struct {
		// 		Name    string
		// 		Profile struct {
		// 			Age  int
		// 			City string
		// 		}
		// 	}{"Chris", struct {
		// 		Age  int
		// 		City string
		// 	}{33, "London"}},
		// 	[]string{"Chris", "London"},
		// },
		{
			"Nested fields",
			Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Pointers to things",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Slices",
			[]Profile{
				{33, "London"},
				{34, "Reykjavik"},
			},
			[]string{"London", "Reykjavik"},
		},
		{
			"Maps",
			map[string]string{
				"Foo": "Bar",
			[]string{"Bar", "Boz"},
		},
	}

	for _, test := range cases {

		t.Run(test.Name, func(t *testing.T) {

			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
