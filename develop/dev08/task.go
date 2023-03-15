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
	default:
		fmt.Println("Введенной команды не существует")
		return nil, errors.New("bad command")
	}
}

func Command(comandLine string) {
	comandLine = strings.Trim(comandLine, " ")
	comandLineArr := strings.Split(comandLine, " ")
	operation, err := Selector(comandLineArr)
	if err != nil {
		fmt.Println(err)
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
