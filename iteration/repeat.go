package iteration

import "strings"

func Repeat(character string, repeatCount int) string {
	var repeatedString string
	for i := 0; i < repeatCount; i++ {
		repeatedString += character
	}
	return repeatedString
}

func RepeatBuiltIn(character string, repeatCount int) string {
	return strings.Repeat(character, repeatCount)
}
