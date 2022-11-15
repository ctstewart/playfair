package playfair

import (
	"fmt"
	"strings"
)

func convertKeyToStringArray(key string) ([]string, error) {
	key = strings.ToUpper(key)

	// alphabet holds a list of A-Z
	seen_letters := map[string]bool{}

	var strArr []string
	for _, char := range strings.Split(key, "") {
		if char == "J" {
			char = "I"
		}

		if _, was_seen := seen_letters[char]; was_seen {
			return []string{}, fmt.Errorf("duplicate letter in key: %s", char)
		}

		strArr = append(strArr, char)
		seen_letters[char] = true
	}

	for i := int('A'); i <= int('Z'); i++ {
		this_char := string(rune(i))
		if this_char == "J" {
			continue
		}
		if _, was_seen := seen_letters[this_char]; !was_seen {
			strArr = append(strArr, this_char)
		}
	}
	return strArr, nil
}
