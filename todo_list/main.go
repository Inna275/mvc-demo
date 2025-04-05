package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HandleIndex)
	http.HandleFunc("/add", HandleAddTask)
	http.HandleFunc("/remove/", HandleRemoveTask)
	http.HandleFunc("/complete/", HandleCompleteTask)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
