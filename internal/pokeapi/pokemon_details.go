package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/csullivan94/pokedex/internal/pokecache"
)

func GetPokemonDetails(url string, cache *pokecache.Cache) (Pokemon, error) {
	value, exists := cache.Get(url)
	if exists {
		CacheUsed = true
		var pokemonDetails Pokemon
		err := json.Unmarshal(value, &pokemonDetails)
		if err != nil {
			fmt.Printf("No pokemon of that name here!")
			return Pokemon{}, err
		}

		return pokemonDetails, nil
	}

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("error with get request: ", err)
		fmt.Println(url)
		return Pokemon{}, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	defer res.Body.Close()

	var pokemonDetails Pokemon
	err = json.Unmarshal(data, &pokemonDetails)
	if err != nil {
		fmt.Printf("No pokemon of that name here!")
		return Pokemon{}, err
	}

	cache.Add(url, data)

	return pokemonDetails, nil
}
