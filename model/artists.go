package groupie

import (
	"encoding/json"
	"net/http"
)

type Artist struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func FetchArtists() ([]Artist, error) {
	url := "https://groupietrackers.herokuapp.com/api/artists"

	// Send GET request to the API
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Decode the response body into a slice of Artist structs
	var artists []Artist
	err = json.NewDecoder(response.Body).Decode(&artists)
	if err != nil {
		return nil, err
	}

	return artists, nil
}
