package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestSortWithoutKeys(t *testing.T) {
	testCommand := "sort notSortedText.txt"
	expectResult := "1 one two\n2 three four\n3 five six\n4 nine ten\n5 nine ten\n5 nine ten"
	text, err := Execute(testCommand)
	if err != nil {
		t.Errorf(text)
		return
	}
	file, fErr := os.Open(text)
	if fErr != nil {
		t.Errorf(fErr.Error())
	}
	testResult, rErr := ioutil.ReadAll(file)
	if rErr != nil {
		t.Errorf(rErr.Error())
	}
	if expectResult != string(testResult) {
		t.Errorf("Данные не совпадают")
		return
	}
	t.Logf("Тест без Ключей прошел успешно")
}
func TestSortR(t *testing.T) {
	testCommand := "sort -r notSortedText.txt"
	expectResult := `5 nine ten
5 nine ten
4 nine ten
3 five six
2 three four
1 one two`
	text, err := Execute(testCommand)
	if err != nil {
		t.Errorf(text)
		return
	}
	file, fErr := os.Open(text)
	if fErr != nil {
		t.Errorf(fErr.Error())
	}
	testResult, rErr := ioutil.ReadAll(file)
	if rErr != nil {
		t.Errorf(rErr.Error())
	}
	if expectResult != string(testResult) {
		t.Errorf("Данные не совпадают")
		return
	}
	t.Logf("Тест с -r прошел успешно")
}

func TestSortN(t *testing.T) {
	testCommand := "sort -n notSortedText.txt"
	expectResult := "1 one two\n2 three four\n3 five six\n4 nine ten\n5 nine ten\n5 nine ten"
	text, err := Execute(testCommand)
	if err != nil {
		t.Errorf(text)
		return
	}
	file, fErr := os.Open(text)
	if fErr != nil {
		t.Errorf(fErr.Error())
	}
	testResult, rErr := ioutil.ReadAll(file)
	if rErr != nil {
		t.Errorf(rErr.Error())
	}
	if expectResult != string(testResult) {
		t.Errorf("Данные не совпадают")
		return
	}
	t.Logf("Тест с -n прошел успешно")
}

func TestSortU(t *testing.T) {
	testCommand := "sort -u notSortedText.txt"
	expectResult := "1 one two\n2 three four\n3 five six\n4 nine ten\n5 nine ten"
	text, err := Execute(testCommand)
	if err != nil {
		t.Errorf(text)
		return
	}
	file, fErr := os.Open(text)
	if fErr != nil {
		t.Errorf(fErr.Error())
	}
	testResult, rErr := ioutil.ReadAll(file)
	if rErr != nil {
		t.Errorf(rErr.Error())
	}
	if expectResult != string(testResult) {
		t.Errorf("Данные не совпадают")
		return
	}
	t.Logf("Тест с -u прошел успешно")
}
func TestSortComboUR(t *testing.T) {
	testCommand := "sort -u -r notSortedText.txt"
	expectResult := `5 nine ten
4 nine ten
3 five six
2 three four
1 one two`
	text, err := Execute(testCommand)
	if err != nil {
		t.Errorf(text)
		return
	}
	file, fErr := os.Open(text)
	if fErr != nil {
		t.Errorf(fErr.Error())
	}
	testResult, rErr := ioutil.ReadAll(file)
	if rErr != nil {
		t.Errorf(rErr.Error())
	}
	if expectResult != string(testResult) {
		t.Errorf("Данные не совпадают")
		return
	}
	t.Logf("Тест с -u, -r прошел успешно")
}
