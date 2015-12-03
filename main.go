package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseFiles("bin/pages/error.html", "bin/pages/index.html"))

func main() {
	log.Println("Starting server.")
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/script", scriptHandler)
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

/*
rootHandler handles the index page.
*/
func rootHandler(w http.ResponseWriter, r *http.Request) {
	// catch any non-index accessess
	if r.RequestURI != "/" {
		w.WriteHeader(http.StatusNotFound)
		renderError(w, "page doesn't exist", errors.New("no handler defined"))
		return
	}
	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		renderError(w, "server error", err)
		return
	}
}

/*
scriptHandler allows the client to fetch the script for the html files.
*/
func scriptHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("bin/pages/script.js")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		renderError(w, "failed to fetch resources", err)
		return
	}
	_, err = w.Write(data)
	if err != nil {
		log.Println("Failed to write script data:", err)
	}
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
