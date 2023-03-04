package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type grepKey struct {
	ignore      bool
	fixed       bool
	invert      bool
	lineNum     bool
	after       bool
	afterCount  int
	before      bool
	beforeCount int
	context     bool
	count       bool
}

func Grep(commandLine string) (string, error) {
	command := strings.Split(strings.TrimSpace(commandLine), " ")
	if command[0] != "grep" {
		return "", errors.New("не найдена команда " + command[0])
	}
	var grepOperations grepKey
	for i, v := range command {
		switch v {
		case "-v":
			grepOperations.invert = true
		case "-A":
			grepOperations.after = true
			num, err := strconv.Atoi(command[i+1])
			if err != nil {
				return "", err
			}
			grepOperations.afterCount = num
		case "-B":
			grepOperations.before = true
			num, err := strconv.Atoi(command[i+1])
			if err != nil {
				return "", err
			}
			grepOperations.beforeCount = num
		case "-C":
			grepOperations.context = true
			num, err := strconv.Atoi(command[i+1])
			if err != nil {
				return "", err
			}
			grepOperations.beforeCount = num
			grepOperations.afterCount = num
		case "-c":
			grepOperations.count = true
		case "-F":
			grepOperations.fixed = true
		case "-i":
			grepOperations.ignore = true
		case "-n":
			grepOperations.lineNum = true
		}
	}

	file, err := os.Open(command[len(command)-1])
	if err != nil {
		return "", errors.New("файла с таким именем не найдено")
	}
	defer file.Close()

	textLine := bufio.NewScanner(file)
	var resultText bytes.Buffer
	lineCount := 0
	resultCount := 0
	nextLine := 0
	previousLines := make([]string, 0, grepOperations.beforeCount)

	for textLine.Scan() {
		lineCount++
		txt := textLine.Text()
		matchRes := false
		if grepOperations.ignore {
			txt = strings.ToLower(txt)
		}
		if grepOperations.fixed {
			if strings.Contains(txt, command[len(command)-2]) {
				matchRes = true
			}
		} else {
			matchRes, _ = regexp.MatchString(command[len(command)-2], txt)
		}
		if grepOperations.invert {
			matchRes = !matchRes
		}
		if !matchRes {
			if nextLine != 0 {
				resultText.WriteString(txt + "\n")
				nextLine--
			}
			if grepOperations.before {
				if len(previousLines) < cap(previousLines) {
					if grepOperations.lineNum {
						previousLines = append(previousLines,
							strconv.FormatInt(int64(lineCount), 10)+"-"+txt)
					} else {
						previousLines = append(previousLines, txt)
					}
				} else {
					previousLines = previousLines[1:]
					if grepOperations.lineNum {
						previousLines = append(previousLines,
							strconv.FormatInt(int64(lineCount), 10)+"-"+txt)
					} else {
						previousLines = append(previousLines, txt)
					}
				}
			}
			continue
		}
		resultCount++
		if grepOperations.count {
			continue
		}
		if grepOperations.before {
			for _, v := range previousLines {
				resultText.WriteString(v + "\n")
			}
			previousLines = make([]string, 0, grepOperations.beforeCount)
		}
		if grepOperations.lineNum {
			resultText.WriteString(strconv.FormatInt(int64(lineCount), 10) + ":")
		}

		if grepOperations.after {
			nextLine = grepOperations.afterCount
		}
		if grepOperations.context {
			for _, v := range previousLines {
				resultText.WriteString(v + "\n")
			}
			previousLines = previousLines[len(previousLines):]
			nextLine = grepOperations.afterCount
		}
		resultText.WriteString(txt + "\n")
	}
	if grepOperations.count {
		resultText.WriteString(strconv.FormatInt(int64(resultCount), 10))
	}
	return strings.TrimSuffix(resultText.String(), "\n"), nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	commandLine := scanner.Text()
	res, _ := Grep(commandLine)
	fmt.Println(res)
}
