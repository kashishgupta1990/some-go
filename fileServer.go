package main

import (
	"net/http"
	"log"
)

func main() {
	fs := http.FileServer(http.Dir("./"))

	http.Handle("/", fs)

	log.Println("Listing on 5000")
	http.ListenAndServe(":5000", nil)
}
