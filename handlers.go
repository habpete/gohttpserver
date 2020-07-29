package main

import (
	"fmt"
	"net/http"
)

func StartPage (w http.ResponseWriter, r* http.Request) {
	fmt.Fprintf(w, "Start Page!")
}

func UserPage (w http.ResponseWriter, r* http.Request) {
	fmt.Fprintf(w, "Hey user!")
}