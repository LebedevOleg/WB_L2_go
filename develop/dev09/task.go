package main

import (
	"io"
	"net/http"
	"os"
)

func Wget(input string) {
	site, err := http.Get("https://zetcode.com/golang/net-html/")
	if err != nil {
		return
	}
	defer site.Body.Close()
	f, createErr := os.Create("test.html")
	if createErr != nil {
		return
	}

	io.Copy(f, site.Body)

}

func main() {
	/* 	input := bufio.NewScanner(os.Stdin)
	   	input.Scan() */
	Wget("")
}
