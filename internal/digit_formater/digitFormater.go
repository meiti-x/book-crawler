package digit_formater

import (
	"strings"
)

// ConvertPersianDigitsToEnglish remove all non english digit from a string.
func ConvertPersianDigitsToEnglish(input string) string {
	persianToEnglish := map[rune]rune{
		'۰': '0',
		'۱': '1',
		'۲': '2',
		'۳': '3',
		'۴': '4',
		'۵': '5',
		'۶': '6',
		'۷': '7',
		'۸': '8',
		'۹': '9',
		'.': '.',
	}

	var result strings.Builder
	for _, char := range input {
		if val, ok := persianToEnglish[char]; ok {
			result.WriteRune(val)
		} else if char >= '0' && char <= '9' {
			result.WriteRune(char)
		}
	}

	return result.String()
}
