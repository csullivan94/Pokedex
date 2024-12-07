package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/csullivan94/pokedex/internal/pokecache"
)

var CacheUsed bool

func GetLocations(url string, cache *pokecache.Cache) (Location, error) {

	value, exists := cache.Get(url)
	if exists {
		CacheUsed = true
		var locations Location
		err := json.Unmarshal(value, &locations)
		if err != nil {
			fmt.Printf("error unmarshaling data: %v ", err)
			return Location{}, err
		}

		return locations, nil
	}

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("error with get request: ", err)
		fmt.Println(url)
		return Location{}, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, err
	}

	defer res.Body.Close()

	var locations Location
	err = json.Unmarshal(data, &locations)
	if err != nil {
		fmt.Printf("error unmarshaling data: %v ", err)
		return Location{}, err
	}

	cache.Add(url, data)

	return locations, nil

}
