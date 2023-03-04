package main

import (
	"testing"
)

func TestGrep(t *testing.T) {
	commandLine := "grep two text.txt"
	currResult := "one two\ntwo three"
	testResult, err := Grep(commandLine)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if testResult != currResult {
		t.Fatalf("ожтдаемый результат: " + currResult + "\nПолученный результат: " + testResult)
		return
	}
	t.Logf("Тест прошел успешно")

}
