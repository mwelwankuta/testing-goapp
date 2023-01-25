package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Pokimon struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Pokimen = []Pokimon

type PokimonResponse struct {
	Results []Pokimon `json:"results"`
}

func GetPokimon() (Pokimen, string) {
	responseData := PokimonResponse{}

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

type Sprites struct {
	FrontDefault string `json:"front_default"`
	BackDefault  string `json:"back_default"`
}

type CurrentPokimon struct {
	Name    string  `json:"name"`
	Sprites Sprites `json:"sprites"`
	Url     string  `json:"url"`
}

func GetCurrentPokimon(url string) (CurrentPokimon, error) {
	responseData := CurrentPokimon{}

	errorMsg := errors.New("")

	response, err := http.Get(url)
	if err != nil {
		errorMsg = err
	}

	body, err := io.ReadAll(response.Body)
	fmt.Println(url, "url")
	json.Unmarshal([]byte(body), &responseData)

	if err != nil {
		errorMsg = err
	}

	return CurrentPokimon{
		Name:    responseData.Name,
		Url:     responseData.Url,
		Sprites: responseData.Sprites,
	}, errorMsg
}
