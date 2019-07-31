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
		{"5 gets converted to V", 5, "V"},
		{"6 gets converted to VI", 6, "VI"},
		{"7 gets converted to VII", 7, "VII"},
		{"8 gets converted to VIII", 8, "VIII"},
		{"9 gets converted to IX", 9, "IX"},
		{"10 gets converted to X", 10, "X"},
		{"20 gets converted to XX", 20, "XX"},
		{"39 gets converted to XXXIX", 39, "XXXIX"},
		{"40 gets converted to XL", 40, "XL"},
		{"47 gets converted to XLVII", 47, "XLVII"},
		{"49 gets converted to XLIX", 49, "XLIX"},
		{"50 gets converted to L", 50, "L"},
		{"100 gets converted to C", 100, "C"},
		{"400 gets converted to CD", 400, "CD"},
		{"500 gets converted to D", 500, "D"},
		{"900 gets converted to CM", 900, "CM"},
		{"1000 gets converted to M", 1000, "M"},
		{"2019 gets converted to MMXIX", 2019, "MMXIX"},
		{"4000 gets converted to MMMM", 4000, "MMMM"},
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
