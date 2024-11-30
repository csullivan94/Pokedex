package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetLocations(url string) (Location, error) {

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("error with get request: ", err)
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
	return locations, nil

}
