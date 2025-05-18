package pokedex

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/WuggyX2/boot-dev-pokedex/internal/pokecache"
)

type LocationAreaResult struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func RetrieveLocationItems(areaUrl string, cache *pokecache.Cache) (LocationAreaResult, error) {
	result := LocationAreaResult{}

	var byteValues []byte

	byteValues, exits := cache.Get(areaUrl)

	if !exits {
		res, err := http.Get(areaUrl)
		if err != nil {
			return result, err
		}

		defer res.Body.Close()

		byteValues, err = io.ReadAll(res.Body)

		if err != nil {
			return result, err
		}

		cache.Add(areaUrl, byteValues)
	} 
	if err := json.Unmarshal(byteValues, &result); err != nil {
		return result, err
	}

	return result, nil

}
