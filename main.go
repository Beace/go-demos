package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/foo", fooHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func fooHandler (w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1 style='text-align:center'>hello go</h1>"))
}