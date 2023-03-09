package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/create_event")
	http.HandleFunc("/update_event")
	http.HandleFunc("/delete_event")
	http.HandleFunc("/event_for_dat")
	http.HandleFunc("/event_for_month")
	http.HandleFunc("/event_for_week")
	err := http.ListenAndServe(":6666", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("server closed")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
