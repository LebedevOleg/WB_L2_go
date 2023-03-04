package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"
)

func Sort(commandLine string) (string, error) {
	commands := strings.Split(commandLine, " ")
	if commands[0] != "sort" {
		return "", errors.New("Введена не верная команда")
	}

	file, err := os.Open(commands[len(commands)-1])
	if err != nil {
		return "", err
	}
	defer file.Close()

	text := bufio.NewScanner(file)
	var sortedtext bytes.Buffer
	var previousWord string
	for text.Scan() {

	}
	return sortedtext.String(), nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	commandLine := scanner.Text()
	res, _ := Sort(commandLine)
	fmt.Println(res)
}
