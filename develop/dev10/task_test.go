package main

import "testing"

func TestTelnetCorrect(t *testing.T) {
	testCommand := "telnet pop.yandex.ru 110"
	err := Telnet(testCommand)
	if err != nil {
		t.Errorf("test Faild: " + err.Error())
	}
	t.Logf("Test Complite")
}
func TestTelnetBad(t *testing.T) {
	testCommand := "telnet --timeout=3s yandex.ru 110"
	err := Telnet(testCommand)
	if err != nil {
		t.Errorf("test Faild: " + err.Error())
	}
	t.Logf("Test Complite")
}
