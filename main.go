package main

import (
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	// Render the home HTML page from static folder
	http.ServeFile(w, r, "static/public/index.html")
}

func main() {
	// Serve static files from the "static/public" directory
	fs := http.FileServer(http.Dir("static/public"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Serve the home page
	http.HandleFunc("/", homePage)

	// Start the server on port 8080
	log.Println("Listening on :8080...")
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
