package romans

import (
	"fmt"
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	testCases := []struct {
		Numeral int
		Roman   string
	}{
		{Numeral: 1, Roman: "I"},
		{Numeral: 2, Roman: "II"},
		{Numeral: 3, Roman: "III"},
		{Numeral: 4, Roman: "IV"},
		{Numeral: 5, Roman: "V"},
		{Numeral: 6, Roman: "VI"},
		{Numeral: 7, Roman: "VII"},
		{Numeral: 8, Roman: "VIII"},
		{Numeral: 9, Roman: "IX"},
		{Numeral: 10, Roman: "X"},
		{Numeral: 14, Roman: "XIV"},
		{Numeral: 18, Roman: "XVIII"},
		{Numeral: 20, Roman: "XX"},
		{Numeral: 39, Roman: "XXXIX"},
		{Numeral: 40, Roman: "XL"},
		{Numeral: 47, Roman: "XLVII"},
		{Numeral: 49, Roman: "XLIX"},
		{Numeral: 50, Roman: "L"},
		{Numeral: 100, Roman: "C"},
		{Numeral: 90, Roman: "XC"},
		{Numeral: 400, Roman: "CD"},
		{Numeral: 500, Roman: "D"},
		{Numeral: 900, Roman: "CM"},
		{Numeral: 1000, Roman: "M"},
		{Numeral: 1984, Roman: "MCMLXXXIV"},
		{Numeral: 3999, Roman: "MMMCMXCIX"},
		{Numeral: 2014, Roman: "MMXIV"},
		{Numeral: 1006, Roman: "MVI"},
		{Numeral: 798, Roman: "DCCXCVIII"},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%d gets converted to %q", testCase.Numeral, testCase.Roman), func(t *testing.T) {
			got := ConvertToRoman(testCase.Numeral)

			if testCase.Roman != got {
				t.Errorf("got %s want %s", got, testCase.Roman)
			}
		})

	}
}
