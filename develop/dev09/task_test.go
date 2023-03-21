package main

import "testing"

func TestWGet(t *testing.T) {
	err := Wget("https://habr.com/ru/company/ru_mts/blog/680324/")
	if err != nil {
		t.Errorf("ошибка парсинга страницы ::" + err.Error())
		return
	}
	t.Logf("парсинг прошел успешно")
}
