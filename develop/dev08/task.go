package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func Selector(command []string) (IOperation, error) {
	switch command[0] {
	case "echo":
		return &Echo{strings.Join(command[1:], " ")}, nil
	case "pwd":
		return &PWD{}, nil
	case "cd":
		return &CD{strings.Join(command[1:], " ")}, nil
	case "kill":
		return &Kill{processName: command[1]}, nil
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
	comandLineArr := strings.Split(comandLine, " ")
	operation, err := Selector(comandLineArr)
	if err != nil {
		if err.Error() == "exit" {
			os.Exit(0)
		}
		fmt.Println(err)
		return
	}
	operation.Operation()
}

func main() {
	for {
		scan := bufio.NewScanner(os.Stdin)
		scan.Scan()
		Command(scan.Text())
	}
}
