package main

import (
	"bytes"
	"strconv"
)

func SimpleUnpackingString(text string) string {
	runeText := []rune(text)
	var resultText bytes.Buffer
	for i := 0; i < len(runeText); i++ {
		if number, err := strconv.Atoi(string(runeText[i])); err == nil {
			if i == 0 {
				break
			}
			for j := 1; j < number; j++ {
				resultText.WriteRune(runeText[i-1])
			}
			continue
		}
		resultText.WriteRune(runeText[i])

	}

	return resultText.String()
}

func UnpackingBreakString(text string) string {
	runeText := []rune(text)
	var resultText bytes.Buffer
	checkBreak := false
	for i := 0; i < len(runeText); i++ {
		if number, err := strconv.Atoi(string(runeText[i])); err == nil && !checkBreak {
			if i == 0 {
				break
			}
			for j := 1; j < number; j++ {
				resultText.WriteRune(runeText[i-1])
			}
			continue
		}
		if runeText[i] == '\\' {
			if checkBreak {
				resultText.WriteRune(runeText[i])
				checkBreak = !checkBreak
				continue
			}
			checkBreak = !checkBreak
			continue
		}
		resultText.WriteRune(runeText[i])
		if checkBreak {
			checkBreak = !checkBreak

		}

	}
	return resultText.String()
}
