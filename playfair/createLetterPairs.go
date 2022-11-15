package playfair

import (
	"regexp"
	"strings"
)

func createLetterPairs(s string) [][2]string {
	// Sanitize the string by removing non uppecase letters
	s = strings.ToUpper(s)
	invaldCharRegex := regexp.MustCompile("[^A-Z]")
	s = string(invaldCharRegex.ReplaceAll([]byte(s), []byte{}))

	strArr := strings.Split(s, "")

	var lp [][2]string
	for {
		// If there are no letters, then we're done
		if len(strArr) == 0 {
			break
		}

		var innerLp [2]string

		// If there is only one letter left
		if len(strArr) == 1 {
			innerLp[0] = strArr[0]
			innerLp[1] = "Z"
			lp = append(lp, innerLp)
			break
		}

		// Add the first letter to the list
		innerLp[0] = strArr[0]

		// Add an x if the next letter is the same, and loop with [1:]
		if strArr[0] == strArr[1] {
			innerLp[1] = "X"
			lp = append(lp, innerLp)
			strArr = strArr[1:]
		} else {
			innerLp[1] = strArr[1]
			lp = append(lp, innerLp)
			strArr = strArr[2:]
		}
	}
	return lp
}
