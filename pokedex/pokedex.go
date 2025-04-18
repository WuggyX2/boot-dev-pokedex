package pokedex

import (
	"encoding/json"
	"net/http"
)

type LocationAreaResult struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func RetrieveLocationItems(areaUrl string) (LocationAreaResult, error) {
	result := LocationAreaResult{}
	res, err := http.Get(areaUrl)

	if err != nil {
		return result, err
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	if err := decoder.Decode(&result); err != nil {
		return result, err
	}

	return result, nil

}
