package main

import (

	groupie_tracker "groupie_tracker/funcs"


	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	staticDir := http.Dir("css")
	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(staticDir)))

	mux.HandleFunc("/", groupie_tracker.MainPage)
	// Artist details route
	mux.HandleFunc("/artist", groupie_tracker.ArtistPage)

	log.Println("Server is running at http://localhost:2000")
	log.Fatal(http.ListenAndServe(":2000", mux))

}
