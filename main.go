package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Note struct {
	Title         string
	Description   string
	StatusCreated time.Time
}

// store for the Notes collection
var noteStore = make(map[string]Note)

// variable to generate key for the collection
var id int = 0

var templates map[string]*template.Template

// compile view templtes
func init() {
	if template == nil {
		templates = make(map[string]*template.Template)
	}
	template["index"] = template.Must(template.ParseFiles("templates/index.html, templates/base.html"))
	template["add"] = template.Must(template.ParseFiles("templates/add.html, templates/base.html"))
	template["edit"] = template.Must(template.ParseFiles("templates/edit.html, templates/base.html"))
}

//entry point of our program
func main() {
	r := mux.NewRouter().StrictSlash(false)
	fs := http.FileServer(http.Dir("public"))
	r.Handle("/public", fs)

	r.Handle("/", getNotes)
	r.Handle("/notes/add", addNote)
	r.Handle("/notes/save", saveNote)
	r.Handle("/notes/edit/{id}", editNote)
	r.Handle("/notes/update/{id}", updateNote)
	r.Handle("/notes/delete/{id}", deleteNote)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Println("listening on port 8080....")
	server.ListenAndServe()

}
