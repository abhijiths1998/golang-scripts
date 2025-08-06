package main

import (
	"html/template"
	"log"
	"net/http"
	"sync"
)

type Post struct {
	Title   string
	Content string
}

var (
	templates = template.Must(template.ParseGlob("templates/*.html"))
	posts     []Post
	mu        sync.Mutex
)

func main() {
	http.HandleFunc("/", listPosts)
	http.HandleFunc("/new", newPost)
	http.HandleFunc("/create", createPost)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func listPosts(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	if err := templates.ExecuteTemplate(w, "index.html", posts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func newPost(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "new.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func createPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}
	post := Post{
		Title:   r.FormValue("title"),
		Content: r.FormValue("content"),
	}
	mu.Lock()
	posts = append(posts, post)
	mu.Unlock()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
