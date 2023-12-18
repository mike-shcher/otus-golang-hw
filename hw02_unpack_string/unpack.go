package hw02unpackstring

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(testString string) (string, error) {
	re := regexp.MustCompile(`[^a-z0-9]+`)
	if len(re.FindStringIndex(testString)) > 0 {
		return "", ErrInvalidString
	}
	var resultString strings.Builder
	bufferedChar := "initial value"

	for i, char := range testString {
		isDigit := false

		numberOfRepetitions, err := strconv.Atoi(string(char))
		if err != nil {
			if bufferedChar == "initial value" ||
				bufferedChar == "" &&
					i != len(testString)-1 {
				bufferedChar = string(char)

				continue
			}
		} else {
			isDigit = true
		}

		if i == 0 && isDigit {
			return "", ErrInvalidString
		}

		if bufferedChar != "" && isDigit {
			resultString.WriteString(strings.Repeat(bufferedChar, numberOfRepetitions))
			bufferedChar = ""
		} else if bufferedChar == "" && isDigit {
			return "", ErrInvalidString
		}

		if bufferedChar != "" && !isDigit {
			resultString.WriteString(bufferedChar)
			bufferedChar = string(char)
		}

		if i == len(testString)-1 && !isDigit {
			resultString.WriteString(string(char))
		}
	}

	return resultString.String(), nil
}
