package propertyBasedTests

import "testing"

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Name   string
		Number int
		Want   string
	}{
		{"1 gets converted to I", 1, "I"},
		{"2 gets converted to II", 2, "II"},
		{"3 gets converted to III", 3, "III"},
		{"4 gets converted to IV", 4, "IV"},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			got := ConvertToRoman(test.Number)
			want := test.Want

			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		})
	}
}
