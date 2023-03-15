package main

import (
	"fmt"
	"os"
)

type IOperation interface {
	Operation()
}

type Echo struct {
	text string
}

func (e *Echo) Operation() {
	fmt.Println(e.text)
}

type PWD struct{}

func (pwd *PWD) Operation() {
	path, err := os.Getwd()
	if err != nil {
		return
	}
	fmt.Println(path)
}

type CD struct {
	path string
}

func (cd *CD) Operation() {
	os.Chdir(cd.path)
}
