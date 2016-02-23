package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//notes object definition
type Note struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedOn   time.Time `json:"createdon"`
}

//store for the notes collection
var noteStore = make(map[string]Note)

//Variable to generate key for the store collection
var id int = 0

//HTTP POST /api/notes
func PostNoteHandler(w http.ResponseWriter, r *http.Request) {
	var note Note
	//Decode the incoming Note json
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		panic(err)
	}
	note.CreatedOn = time.Now()
	id++
	k := strconv.Itoa(id)
	noteStore[k] = note

	j, err := json.Marshal(note)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

//HTTP GET /api/notes
func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	var notes []Note
	for _, v := range noteStore {
		notes = append(notes, v)
	}
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(notes)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//HTTP PUT /api/notes/{id}
func PutNoteHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	k := vars["id"]
	var noteToUpd Note

	// Decode the incoming Note json
	err = json.NewDecoder(r.Body).Decode(&noteToUpd)
	if err != nil {
		panic(err)
	}
	if note, ok := noteStore[k]; ok {
		noteToUpd.CreatedOn = note.CreatedOn
		// delete the existing item and add the updated item
		delete(noteStore, k)
		noteStore[k] = noteToUpd
	} else {
		log.Printf("Could not find key of note %s to delete", k)
	}
	w.WriteHeader(http.StatusNoContent)
}

// HTTP DELETE api/note/{id}
func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	k := vars["id"]

	//Remove from store
	if _, ok := noteStore[k]; ok {
		//delete existing item
		delete(noteStore, k)
	} else {
		log.Printf("Could not find key of Note %s to delete", k)
	}
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("api/notes", GetNoteHandler).Methods("GET")
	r.HandleFunc("api/notes", PostNoteHandler).Methods("POST")
	r.HandleFunc("api/notes/{id}", PutNoteHandler).Methods("PUT")
	r.HandleFunc("api/notes/{id}", DeleteNoteHandler).Methods("DELETE")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Println("Listening on port 8080....")
	server.ListenAndServe()
}
