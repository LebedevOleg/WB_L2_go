package main

import "testing"

func TestDFKeys(t *testing.T) {
	inputLine := "cut -d: -f2 text.txt"
	expectedRes := "Company\nCompany1\nCompany2\nCompany3\nCompany4\nCompany5\nCompany6\nitem 7 company 7 322322$ ganre 3 no"
	testResult, err := Cut(inputLine)
	if err != nil {
		t.Fatalf(err.Error())
		return
	}
	if expectedRes != testResult {
		//t.Fatalf("Результат не верен")
		t.Fatalf("expected:\n" + expectedRes + "\n Geted: \n" + testResult)
		return
	}
	t.Logf("Тест прошел успешно")

}

//todo: add more test
