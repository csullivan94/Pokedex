package main

import (
	"fmt"
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

func givePageNumber(cfg *Config, c *pokecache.Cache, page int) error {
	if page > 52 {
		fmt.Println("Outside of page range, last page 51")
		return fmt.Errorf("outside of page range")
	}

	url := cfg.Current
	newOffsetInt := (page - 1) * 20
	newOffset := strconv.Itoa(newOffsetInt) + "&"

	re := regexp.MustCompile(`\d+&`)
	offsetString := string(re.ReplaceAll([]byte(url), []byte(newOffset)))
	cfg.Current = offsetString
	return nil
}
