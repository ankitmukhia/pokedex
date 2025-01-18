package internal

import (
	"net/http"
	"encoding/json"
	"log"
	"io"
)

const (
	baseUrl = "https://pokeapi.co/api/v2"
)

func locationLists() (RespLocations, error) {
	res, err := http.Get(baseUrl + "/location-area")
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)

	cfg := RespLocations{}

	err = json.Unmarshal(body, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	return cfg, nil
}
