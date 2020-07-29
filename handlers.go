package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"io/ioutil"
)

func StartPage (w http.ResponseWriter, r* http.Request) {
	fmt.Fprintf(w, "Start Page!")
}

func UserPage (w http.ResponseWriter, r* http.Request) {
	fmt.Fprintf(w, "Hey user!")
}

var pathToStatic string = "./templates"

func StaticFilesHandler (path) func (w http.ResponseWriter, r* http.Request) {
	if (path == "/")
		path += "index"
	content, err := ioutil.ReadFile(path.join(pathToStatic, path + ".html"))
	if (err != nil) {
		fmt.Fprintf(w, ioutil.ReadFile(path.join(pathToStatic, path + ".html")));
		return
	}
	fmt.Fprintf(w, content);
}