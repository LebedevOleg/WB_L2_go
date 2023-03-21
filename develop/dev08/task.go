package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func Selector(command string) (IOperation, error) {
	command = strings.Trim(command, " ")
	commandArr := strings.Split(command, " ")
	switch commandArr[0] {
	case "echo":
		return &Echo{text: strings.Join(commandArr[1:], " ")}, nil
	case "pwd":
		return &PWD{}, nil
	case "cd":
		return &CD{path: strings.Join(commandArr[1:], " ")}, nil
	case "kill":
		return &Kill{processName: commandArr[1]}, nil
	case "ps":
		return &PS{}, nil
	case "\\quit":
		return nil, errors.New("exit")
	default:
		return nil, errors.New("Введенной команды не существует")
	}
}

func Command(comandLine string) {
	comandLine = strings.Trim(comandLine, " ")
	commandPipes := strings.Split(comandLine, "|")
	for i := 0; i < len(commandPipes); i++ {
		operation, err := Selector(commandPipes[i])
		if err != nil {
			fmt.Println(err)
			return
		}
		operation.Operation()
	}
}

func main() {
	for {
		scan := bufio.NewScanner(os.Stdin)
		scan.Scan()
		Command(scan.Text())
	}
}
