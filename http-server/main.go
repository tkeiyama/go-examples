package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello world\n")
}

func main() {
	http.HandleFunc("/", hello)

	err := http.ListenAndServe(":8080", nil)
	if err == nil {
		log.Fatal("Couldn't start the server")
	}
}
