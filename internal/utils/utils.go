package utils

import "strings"

func ToUpperFirst(text string) string {

	if text == "" {
		return ""
	}

	chars := strings.Split(text, "")
	chars[0] = strings.ToUpper(chars[0])

	return strings.Join(chars, "")
}
