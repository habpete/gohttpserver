package main

import (
	"net/http"
)

func main () {
	// http.HandleFunc("/", StartPage)
	// http.HandleFunc("/user", UserPage)
	http.Handle()
	http.ListenAndServe(":80", nil)
}