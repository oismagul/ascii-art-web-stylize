package service

import (
	"errors"
	"strings"
)

func Validation(input string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("empty input")
	}

	input = strings.ReplaceAll(input, `\n`, "\n")

	if err := NotAllowedChar(input); err != nil {
		return "", err
	}

	var result []rune

	if OnlyNewLines(input) {
		for i := 0; i < len(input)-1; i++ {
			if input[i] == '\\' && input[i+1] == 'n' {
				result = append(result, '\n')
				i++
			}
		}
		return string(result), nil
	}

	return input, nil
}

func NotAllowedChar(str string) error {
	for _, char := range str {
		if (char < 32 || char > 126) && char != 10 {
			return errors.New("Invalid input: non-ASCII character '" + string(char) + "'")
		}
	}
	return nil
}

func OnlyNewLines(s string) bool {
	i := 0
	for i < len(s) {
		switch s[i] {
		case '\\':
			if i+1 < len(s) && s[i+1] == 'n' {
				i += 2
			} else {
				return false
			}
		case ' ':
			i++
		default:
			return false
		}
	}
	return true
}
