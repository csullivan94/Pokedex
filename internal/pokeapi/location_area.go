package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/csullivan94/pokedex/internal/pokecache"
)

func GetLocationPokemon(url string, cache *pokecache.Cache) (Location_area, error) {
	value, exists := cache.Get(url)
	if exists {
		CacheUsed = true
		var locationArea Location_area
		err := json.Unmarshal(value, &locationArea)
		if err != nil {
			fmt.Printf("error unmarshaling data: %v ", err)
			return Location_area{}, err
		}

		return locationArea, nil
	}

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("error with get request: ", err)
		fmt.Println(url)
		return Location_area{}, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Location_area{}, err
	}

	defer res.Body.Close()

	var locationArea Location_area
	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		fmt.Printf("error unmarshaling data: %v ", err)
		return Location_area{}, err
	}

	cache.Add(url, data)

	return locationArea, nil

}
