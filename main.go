package main

import (
	"net/http"
)

func main () {
	http.HandleFunc("/", StartPage)
	http.HandleFunc("/user", UserPage)
	http.ListenAndServe(":80", nil)
}