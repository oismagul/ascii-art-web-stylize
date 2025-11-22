package service

import (
	"os"
	"strings"
)

// LoadBanner loads the ASCII art characters from a banner file
func LoadBanner(bannerName string) (map[rune][]string, error) {
	data, err := os.ReadFile("././banners/" + strings.ToLower(bannerName) + ".txt")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	for len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	banner := make(map[rune][]string)

	index := 1
	currentRune := ' '

	if strings.Contains(bannerName, "thinkertoy") {
		for index+8 <= len(lines) {
			charLines := lines[index : index+8]
			for i := range charLines {
				charLines[i] = strings.ReplaceAll(charLines[i], "\r", "")
			}
			banner[currentRune] = charLines
			currentRune++
			index += 9
		}
	} else {
		for index+8 <= len(lines) {
			charLines := lines[index : index+8]
			banner[currentRune] = charLines
			currentRune++
			index += 9
		}
	}

	return banner, nil
}
