package main

import (
	"bufio"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func Wget(input string) error {
	site, err := http.Get(input)
	if err != nil {
		return err
	}
	defer site.Body.Close()
	f, createErr := os.Create("test.html")
	if createErr != nil {
		return err
	}
	scanner := bufio.NewScanner(site.Body)
	linkCheck := false
	for scanner.Scan() {
		currLine := scanner.Text()
		if !linkCheck {
			matched, err := regexp.MatchString(`rel="stylesheet"`, currLine)
			if err != nil {
				return err
			}
			if matched {
				lineArr := strings.Split(currLine, " ")
				for i, v := range lineArr {
					href, err := regexp.MatchString(`^href=`, v)
					if err != nil {
						return err
					}
					if href {
						link := strings.Split(v, `"`)
						lineArr[i] = link[0] + `"` + strings.Join(strings.Split(input, "/")[:3], "/") + strings.Join(link[1:], `"`)
						styleSheet := strings.Join(lineArr, " ")
						f.WriteString(styleSheet)
						linkCheck = true
						break
					}
				}
				continue
			}
		}
		f.WriteString(currLine)
	}
	return nil
}

func main() {
	Wget("https://habr.com/ru/company/ru_mts/blog/680324/")
}
