package romans

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

var romanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(numeral int) string {
	var result strings.Builder
	for _, romanNumeral := range romanNumerals {
		for numeral >= romanNumeral.Value {
			result.WriteString(romanNumeral.Symbol)
			numeral -= romanNumeral.Value
		}
	}
	return result.String()
}
