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
		fmt.Fprintln(os.Stderr, errors.New("args is not 3"))
		fmt.Println("example: simpleHTMLhandler <html> <port>")
		os.Exit(1)
	}
	f = os.Args[1]
	port := os.Args[2]

	http.HandleFunc("/", simpleHTMLhandler)
	fmt.Println(fmt.Sprintf("http://localhost:%s", port))
	http.ListenAndServe(":"+port, nil)
}
