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
		var err error

		byteValues, err = queryAndCache(areaUrl, cache)

		if err != nil {
			return result, err
		}
	}

	if err := json.Unmarshal(byteValues, &result); err != nil {
		return result, err
	}

	return result, nil

}

type PokemonsInAreaResult struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func GetPokemonsInArea(areaUrl string, cache *pokecache.Cache) (PokemonsInAreaResult, error) {
	result := PokemonsInAreaResult{}

	var byteValues []byte
	byteValues, exists := cache.Get(areaUrl)

	if !exists {
		var err error

		byteValues, err = queryAndCache(areaUrl, cache)

		if err != nil {
			return result, err
		}
	}

	if err := json.Unmarshal(byteValues, &result); err != nil {
		return result, err
	}

	return result, nil
}

func queryAndCache(url string, cache *pokecache.Cache) ([]byte, error) {

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	byteValues, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	cache.Add(url, byteValues)

	return byteValues, nil
}
