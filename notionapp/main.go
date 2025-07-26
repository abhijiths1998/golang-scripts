package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

// Todo represents a todo item
type Todo struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

// Page represents a page with markdown content and todos
type Page struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	ParentID *int   `json:"parent_id,omitempty"`
	Todos    []Todo `json:"todos"`
	Children []int  `json:"children"`
}

var (
	pages   = make(map[int]*Page)
	pagesMu sync.Mutex
	nextID  int
	todoID  int
)

func main() {
	http.HandleFunc("/api/pages", pagesHandler)
	http.HandleFunc("/api/pages/", pageHandler)
	http.HandleFunc("/", serveIndex)

	log.Println("Starting server on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, "frontend/index.html")
}

// pagesHandler handles listing and creating top-level pages
func pagesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		pagesMu.Lock()
		defer pagesMu.Unlock()
		list := []Page{}
		for _, p := range pages {
			if p.ParentID == nil {
				list = append(list, *p)
			}
		}
		json.NewEncoder(w).Encode(list)
	case http.MethodPost:
		var p Page
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		pagesMu.Lock()
		p.ID = nextID
		nextID++
		pages[p.ID] = &p
		if p.ParentID != nil {
			if parent, ok := pages[*p.ParentID]; ok {
				parent.Children = append(parent.Children, p.ID)
			}
		}
		pagesMu.Unlock()
		json.NewEncoder(w).Encode(p)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

// pageHandler handles fetching a page and manipulating todos
func pageHandler(w http.ResponseWriter, r *http.Request) {
	rest := r.URL.Path[len("/api/pages/"):]
	if rest == "" {
		http.NotFound(w, r)
		return
	}
	parts := strings.Split(rest, "/")
	idStr := parts[0]
	pagesMu.Lock()
	p, ok := pages[toInt(idStr)]
	pagesMu.Unlock()
	if !ok {
		http.NotFound(w, r)
		return
	}

	// handle todo update
	if len(parts) == 3 && parts[1] == "todos" {
		todoIDStr := parts[2]
		if r.Method != http.MethodPut {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		pagesMu.Lock()
		for i := range p.Todos {
			if p.Todos[i].ID == toInt(todoIDStr) {
				var body struct {
					Done bool `json:"done"`
				}
				if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
					pagesMu.Unlock()
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				p.Todos[i].Done = body.Done
				json.NewEncoder(w).Encode(p.Todos[i])
				pagesMu.Unlock()
				return
			}
		}
		pagesMu.Unlock()
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(p)
	case http.MethodPut:
		var updated Page
		if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		pagesMu.Lock()
		p.Title = updated.Title
		p.Content = updated.Content
		pagesMu.Unlock()
		json.NewEncoder(w).Encode(p)
	case http.MethodPost:
		// create todo
		var t Todo
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		pagesMu.Lock()
		t.ID = todoID
		todoID++
		p.Todos = append(p.Todos, t)
		pagesMu.Unlock()
		json.NewEncoder(w).Encode(t)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
