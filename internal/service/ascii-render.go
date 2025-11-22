package service

import (
	"strings"
)

func PrintASCII(input string, banner map[rune][]string) (string, error) {
	result := []rune{}

	input, err := Validation(input)
	if err != nil {
		return "", err
	}

	lines := strings.Split(input, "\n")
	if len(lines) == 2 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	for _, line := range lines {
		if line == "" {
			result = append(result, '\n')
			continue
		}

		for i := 0; i < 8; i++ {
			for _, char := range line {
				if ascii, ok := banner[char]; ok {
					result = append(result, []rune(ascii[i])...)
				} else {
					result = append(result, []rune("        ")...)
				}
			}
			result = append(result, '\n')
		}
	}

	if len(result) > 0 && result[len(result)-1] == '\n' {
		result = result[:len(result)-1]
	}

	return string(result), nil
}
