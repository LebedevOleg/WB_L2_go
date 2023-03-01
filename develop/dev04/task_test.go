package main

import "testing"

func TestAnagramms(t *testing.T) {
	testWords := []string{"пятак", "сЛиток", "тяпкА", "пятак", "листок", "пятка", "столик", "ключ"}
	testResult := SearchAnagramms(testWords)
	t.Log(testResult)
}
