package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseFiles("bin/pages/error.html"))

func main() {
	log.Println("Starting server.")
	http.HandleFunc("/", handler)
	http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		renderError(w, "test error", errors.New("test error, please ignore"))
	})
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Println("Failed to start server:", err)
		return
	}
	log.Println("Stopping server.")
}

func handler(w http.ResponseWriter, r *http.Request) {
	// catch any non-index accessess
	if r.RequestURI != "/" {
		renderError(w, "page doesn't exist", errors.New("no handler defined"))
		return
	}
	fmt.Fprintf(w, "<html><head><title>Test</title></head><body>Hello World!</br>User Agent: %s</br>Referer: %s</br><a href=\"http://localhost:8000/error\">Click here to see an error.</a></body></html>", r.UserAgent(), r.Referer())
}

/*
renderError renders the user render page and logs the error for the server.
*/
func renderError(w http.ResponseWriter, reason string, err error) {
	log.Printf("User error: <%s> Explanation given: <%s>.", err, reason)
	err = templates.ExecuteTemplate(w, "error.html", reason)
	if err != nil {
		log.Println("renderError failed on template execute:", err)
	}
}
