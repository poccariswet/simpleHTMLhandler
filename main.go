package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

var f string

func simpleHTMLhandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	t.Execute(w, nil)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, errors.New("args is not 2"))
		os.Exit(1)
	}
	f = os.Args[1]

	http.HandleFunc("/", simpleHTMLhandler)
	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
