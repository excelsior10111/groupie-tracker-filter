package main

import (
	"flag"
	app "groupie/internal/services"
	"log"
	"net/http"
	"time"
)

func main() {
	h := new(app.Handler)
	mux := http.NewServeMux()
	mux.HandleFunc("/", h.Home)
	mux.HandleFunc("/artist", h.ArtistInfo)
	port := flag.String("listenAddr", ":5000", "The listen adress of the API server")
	flag.Parse()
	fileserver := http.FileServer(http.Dir("./ui/static/"))
	srv := &http.Server{
		Handler:      mux,
		Addr:         *port,
		WriteTimeout: time.Second * 9,
		ReadTimeout:  time.Second * 7,
	}
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))
	err := srv.ListenAndServe()
	log.Fatal(err)
}
