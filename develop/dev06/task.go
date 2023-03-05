package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IKey interface{}

type Dkey struct {
	deilmiter string
}

func (d *Dkey) Split(s string) []string {
	return strings.Split(s, d.deilmiter)
}

type Skey struct {
	enable bool
}

func (s Skey) Operation(txt []string) bool {
	return len(txt) <= 1
}

type Fkey struct {
	d         Dkey
	s         Skey
	columnNum []int
}

func (f *Fkey) Operation(s string) (string, bool) {
	var resultLine bytes.Buffer
	if f.d.deilmiter == "" {
		resultLine.WriteString(s)
		return resultLine.String(), true
	}
	lineArray := f.d.Split(s)
	if f.s.Operation(lineArray) {
		if f.s.enable {
			return "", false
		}
		resultLine.WriteString(s)
		return resultLine.String(), true
	}

	for i, v := range lineArray {
		for _, idx := range f.columnNum {
			if idx == i {
				resultLine.WriteString(v + f.d.deilmiter)
			}
		}
	}
	if resultLine.Len() == 0 {
		return "", true
	}
	return strings.Trim(resultLine.String(), f.d.deilmiter), true
}
func NewFkey(key string, currKey Fkey) Fkey {
	f := currKey
	intervalCheck := false
	for _, v := range key {
		if num, err := strconv.Atoi(string(v)); err == nil && !intervalCheck {
			f.columnNum = append(f.columnNum, num-1)
			continue
		} else if err == nil && intervalCheck {
			i := f.columnNum[len(f.columnNum)-1] + 1
			for ; i < num; i++ {
				f.columnNum = append(f.columnNum, i)
			}
		}
		if v == '-' {
			intervalCheck = !intervalCheck
			continue
		}
	}
	return f
}

func Cut(commandLine string) (string, error) {
	commands := strings.Split(commandLine, " ")
	if commands[0] != "cut" {
		return "", errors.New("wrong command")
	}
	F := Fkey{}
	for _, v := range commands {
		if string([]rune(v)[:2]) == "-f" {
			F = NewFkey(string([]rune(v)[2:]), F)
			continue
		}
		if string([]rune(v)[:2]) == "-d" {
			F.d.deilmiter = string([]rune(v)[2:])
			continue
		}
		if string([]rune(v)[:2]) == "-s" {
			F.s.enable = true
			continue
		}

	}
	file, err := os.Open(commands[len(commands)-1])
	if err != nil {
		return "", err
	}
	defer file.Close()
	data := bufio.NewScanner(file)

	var curResult bytes.Buffer

	for data.Scan() {
		line := data.Text()
		if s, ok := F.Operation(line); ok {
			curResult.WriteString(s + "\n")
		}
	}
	return strings.Trim(curResult.String(), "\n"), nil
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	result, _ := Cut(input.Text())
	fmt.Println(result)

}
