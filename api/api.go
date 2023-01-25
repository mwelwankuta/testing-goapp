package api

import (
	"encoding/json"
	"io"
	"net/http"
)

type Pokimon struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Pokimen = []Pokimon

type Response struct {
	Results []Pokimon `json:"results"`
}

func GetPokimon() (Pokimen, string) {
	responseData := Response{}

	errorMsg := ""

	response, err := http.Get("https://pokeapi.co/api/v2/pokemon")
	if err != nil {
		errorMsg = err.Error()
	}

	body, err := io.ReadAll(response.Body)
	json.Unmarshal([]byte(body), &responseData)

	if err != nil {
		errorMsg = err.Error()
	}

	return responseData.Results, errorMsg
}
