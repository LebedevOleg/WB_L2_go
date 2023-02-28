package main

import (
	"os"
	"testing"
)

// todo: Дописать тесты
func TestCurrTime(t *testing.T) {
	currTime, err := GetCurrTime()
	if err != nil {
		t.Errorf("Curr time: " + err.Error())
		os.Exit(126)
	}
	t.Logf("Curr time: " + currTime)
	exactTime, err := GetExactTime()
	if err != nil {
		t.Errorf("Exact time: " + err.Error())
		os.Exit(126)
	}
	t.Logf("Exact time: " + exactTime)
}
