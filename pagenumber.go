package main

import (
	"regexp"
	"strconv"

	"github.com/csullivan94/pokedex/internal/pokecache"
)

func getPageNumber(cfg *Config, c *pokecache.Cache) error {
	url := cfg.Current
	re := regexp.MustCompile(`\d+`)
	offsetString := string(re.Find([]byte(url[35:])))
	offset, err := strconv.Atoi(offsetString)
	if err != nil {
		return err
	}
	pageNumber := (offset / 20) + 1

	cfg.PageNum = pageNumber

	return nil
}
