package main

import (
	"fmt"
	groupie "groupie-tracker/model"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	PORT = ":8080"
	path = "view/"
)

func main() {

	http.HandleFunc("/", WelcomePage)
	http.HandleFunc("/artists", Artists)
	http.HandleFunc("/locations", Locations)

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	fmt.Println("(http://localhost:8080) Server started on port:", PORT)
	http.ListenAndServe(PORT, nil)
}

func WelcomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		fmt.Fprintf(w, "Page error")
		return
	}

	renderTemplate(w, "index", nil)
}

func Artists(w http.ResponseWriter, r *http.Request) {
	artists, err := groupie.FetchArtists()
	if err != nil {
		log.Fatal(err)
	}

	// Print the ID and name of each artist
	for _, artist := range artists {
		fmt.Printf("ID: %d, Name: %s\n", artist.Id, artist.Name)
	}
}

func Locations(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(w, string(responseData))
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles(path + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
