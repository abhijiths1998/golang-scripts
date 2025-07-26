package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

// Note represents a single note
type Note struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

// store holds notes in memory
var (
	notes  []Note
	mu     sync.Mutex
	nextID int
)

func main() {
	http.HandleFunc("/notes", notesHandler)
	http.HandleFunc("/", serveIndex)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// serveIndex serves the static HTML page
func serveIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, "static/index.html")
}

// notesHandler handles GET and POST for notes
func notesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		mu.Lock()
		defer mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(notes)
	case http.MethodPost:
		var n Note
		if err := json.NewDecoder(r.Body).Decode(&n); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		mu.Lock()
		n.ID = nextID
		nextID++
		notes = append(notes, n)
		mu.Unlock()
		w.WriteHeader(http.StatusCreated)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
