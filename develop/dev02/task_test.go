package main

import (
	"testing"
)

func TestUnpackage(t *testing.T) {
	testString := []string{"a4bc2d5e", "abcd", "45", ""}
	TestResult := make([]string, len(testString))
	CurrResult := []string{"aaaabccddddde", "abcd", "", ""}

	for i, v := range testString {
		TestResult[i] = SimpleUnpackingString(v)
	}
	for i, v := range TestResult {
		if v != CurrResult[i] {

			t.Errorf("Expected: " + CurrResult[i] + "; Get it: " + TestResult[i])
		}
	}
	t.Logf("First test complite")

	testString = []string{`qwe\4\5`, `qwe\45`, `qwe\\5`}
	TestResult = make([]string, len(testString))
	CurrResult = []string{"qwe45", "qwe44444", `qwe\\\\\`}
	for i, v := range testString {
		TestResult[i] = UnpackingBreakString(v)
	}
	for i, v := range TestResult {
		if v != CurrResult[i] {
			t.Errorf("Expected: " + CurrResult[i] + "; Get it: " + TestResult[i])
		}
	}
	t.Logf("Second test complite")

	if t.Failed() {
		t.Logf("Tests Uncorrect")
	} else {
		t.Logf("Tests Correct")
	}
}
