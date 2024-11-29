package pokeapi

func getLocations() {
	if cfg.Next == "" {
		cfg.Next = "https://pokeapi.co/api/v2/location"
	}

	cfg.PageNum += 1

	res, err := http.Get(cfg.Next)
	if err != nil {
		fmt.Println("error with get request: ", err)
		return err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	var locations Location
	err = json.Unmarshal(data, &locations)
	if err != nil {
		fmt.Printf("error unmarshaling data: %v ", err)
		return err
	}

}
