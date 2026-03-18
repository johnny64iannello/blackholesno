package main

import (
	"log"
	"net/http"
)

func main() {
	// Serve i file statici sotto /static/
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Serve index.html sulla root
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

	log.Println("Server su http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
