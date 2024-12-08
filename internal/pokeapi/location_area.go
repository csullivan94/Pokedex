package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/csullivan94/pokedex/internal/pokecache"
)

func GetLocationPokemon(url string, cache *pokecache.Cache) (LocationArea, error) {
	value, exists := cache.Get(url)
	if exists {
		CacheUsed = true
		var locationArea LocationArea
		err := json.Unmarshal(value, &locationArea)
		if err != nil {
			fmt.Printf("error unmarshaling data: %v ", err)
			return LocationArea{}, err
		}

		return locationArea, nil
	}

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("error with get request: ", err)
		fmt.Println(url)
		return LocationArea{}, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	defer res.Body.Close()

	var locationArea LocationArea
	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		fmt.Printf("error unmarshaling data: %v ", err)
		return LocationArea{}, err
	}

	cache.Add(url, data)

	return locationArea, nil

}
