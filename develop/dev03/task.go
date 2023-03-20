package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

type manSort struct {
	fileName    string
	sortedText  []string
	sortColumn  int
	numSort     bool
	reverseSort bool
	skipRepeat  bool
}

func (ms *manSort) GetText() error {
	file, err := os.Open(ms.fileName)

	if err != nil {
		return err
	}
	defer file.Close()
	text := bufio.NewScanner(file)
	for text.Scan() {
		ms.sortedText = append(ms.sortedText, text.Text())
	}
	return nil
}
func (ms *manSort) Sort() error {
	if ms.skipRepeat {
		for i, v := range ms.sortedText {
			if i+1 < len(ms.sortedText) && v == ms.sortedText[i+1] {
				ms.sortedText = append(ms.sortedText[:i], ms.sortedText[i+1:]...)
			}
		}
	}
	if ms.reverseSort {
		sort.Slice(ms.sortedText, func(i, j int) bool {
			if ms.numSort {
				num1, err1 := strconv.Atoi(strings.Split(ms.sortedText[i], " ")[ms.sortColumn-1])
				if err1 != nil {
					return strings.Split(ms.sortedText[i], " ")[ms.sortColumn-1] > strings.Split(ms.sortedText[j], " ")[ms.sortColumn-1]
				}
				num2, err2 := strconv.Atoi(strings.Split(ms.sortedText[j], " ")[ms.sortColumn-1])
				if err2 != nil {
					return strings.Split(ms.sortedText[i], " ")[ms.sortColumn-1] > strings.Split(ms.sortedText[j], " ")[ms.sortColumn-1]
				}
				return num1 > num2
			}
			return strings.Split(ms.sortedText[i], " ")[ms.sortColumn-1] > strings.Split(ms.sortedText[j], " ")[ms.sortColumn-1]
		})
	} else {
		sort.Slice(ms.sortedText, func(i, j int) bool {
			if ms.numSort {
				num1, err1 := strconv.Atoi(strings.Split(ms.sortedText[i], " ")[ms.sortColumn-1])
				if err1 != nil {
					return strings.Split(ms.sortedText[i], " ")[ms.sortColumn-1] < strings.Split(ms.sortedText[j], " ")[ms.sortColumn-1]
				}
				num2, err2 := strconv.Atoi(strings.Split(ms.sortedText[j], " ")[ms.sortColumn-1])
				if err2 != nil {
					return strings.Split(ms.sortedText[i], " ")[ms.sortColumn-1] < strings.Split(ms.sortedText[j], " ")[ms.sortColumn-1]
				}
				return num1 < num2
			}
			return strings.Split(ms.sortedText[i], " ")[ms.sortColumn-1] < strings.Split(ms.sortedText[j], " ")[ms.sortColumn-1]
		})
	}
	return nil
}
func (ms *manSort) WriteToFile() error {
	var sortedtext bytes.Buffer
	for _, v := range ms.sortedText {
		sortedtext.WriteString(v + "\n")
	}
	err := ioutil.WriteFile("(sorted)"+ms.fileName, sortedtext.Bytes(), 0777)
	if err != nil {
		return err
	}
	return nil
}

func Execute(commandLine string) (string, error) {
	commands := strings.Split(commandLine, " ")
	if commands[0] != "sort" {
		return "", errors.New("введена не верная команда")
	}
	mSort := manSort{fileName: commands[len(commands)-1], sortColumn: 1, numSort: false, reverseSort: false, skipRepeat: false}
	for i := 0; i < len(commands); i++ {
		switch commands[i] {
		case "-k":
			mSort.sortColumn, _ = strconv.Atoi(commands[i+1])
		case "-n":
			mSort.numSort = true
		case "-r":
			mSort.reverseSort = true
		case "-u":
			mSort.skipRepeat = true
		}
	}
	err := mSort.GetText()
	if err != nil {
		return "Ошибка при получении текста", err
	}
	err = mSort.Sort()
	if err != nil {
		return "Ошибка при сортировке", err
	}
	err = mSort.WriteToFile()
	if err != nil {
		return "Ошибка при записи", err
	}
	return "", nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	commandLine := scanner.Text()
	res, _ := Execute(commandLine)
	fmt.Println(res)
}
