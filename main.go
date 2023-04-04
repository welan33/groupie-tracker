package main

import (
	"fmt"
	g "groupie/fonction"
	"net/http"
)

func main() {
	g.APIRequest()
	g.SelectCountry()
	g.RepareTab()
	http.HandleFunc("/", g.MainHandler)
	http.HandleFunc("/artist", g.ArtistHandler)
	http.HandleFunc("/filter", g.NbMembres)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Printf("Serveur lancé\n")
	err := http.ListenAndServe("localhost:8888", nil)

	if err != nil {
		fmt.Println(err)
	}
}
