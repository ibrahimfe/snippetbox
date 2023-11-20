package main

import (
	"log"
	"net/http"
)

// HTTP Handler
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello From SnippetBox"))

}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display Spesific Snippet ...."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new Snippet...."))
}

func main() {
	// Controller for HTTP Request
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting servern on :4000")
	// Start http server on port 4000
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
