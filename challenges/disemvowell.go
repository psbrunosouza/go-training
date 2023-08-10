package challenges

import "fmt"

func Disemvowel(comment string) (result string) {
	vowels := map[string]string{
		"a": "a",
		"e": "e",
		"i": "i",
		"o": "o",
		"u": "u",
		"A": "A",
		"E": "E",
		"I": "I",
		"O": "O",
		"U": "U",
	} 

	for _, v := range comment {
		if value := vowels[fmt.Sprintf("%c", v)]; !(value == fmt.Sprintf("%c", v)) {
			result += fmt.Sprintf("%c", v)
		}
	}
	return
}

