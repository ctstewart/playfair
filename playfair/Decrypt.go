package playfair

import (
	"strings"

	"github.com/pkg/errors"
)

func Decrypt(message string, key string) (string, error) {
	var ans string
	keyTable, err := convertKeyToStringArray(key)
	if err != nil {
		return "", errors.Wrap(err, "creating string array from key")
	}
	letterPairs := createLetterPairs(message)
	for _, pair := range letterPairs {
		c1, c2 := getCoordFromLetter(pair[0], keyTable), getCoordFromLetter(pair[1], keyTable)
		if c1[1] == c2[1] {
			// Same row
			c1[0] = c1[0] - 1
			if c1[0] == -1 {
				c1[0] = 4
			}

			c2[0] = c2[0] - 1
			if c2[0] == -1 {
				c2[0] = 4
			}
		} else if c1[0] == c2[0] {
			// Same column
			c1[1] = c1[1] - 1
			if c1[1] == -1 {
				c1[1] = 4
			}

			c2[1] = c2[1] - 1
			if c2[1] == -1 {
				c2[1] = 4
			}
		} else {
			// Get the other corners of the square
			c1[0], c2[0] = c2[0], c1[0]
		}
		ans += getLetterFromCoord(c1[0], c1[1], keyTable)
		ans += getLetterFromCoord(c2[0], c2[1], keyTable)
	}
	return strings.ToLower(ans), nil
}
