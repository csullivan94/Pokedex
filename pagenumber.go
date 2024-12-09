package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func GetPageNumber(cfg *Config) error {
	url := cfg.CurrentMap
	re := regexp.MustCompile(`\d+`)
	offsetString := string(re.Find([]byte(url[40:])))
	offset, err := strconv.Atoi(offsetString)
	if err != nil {
		return err
	}
	pageNumber := (offset / 20) + 1

	cfg.PageNum = pageNumber

	return nil
}

func GivePageNumber(cfg *Config, page int) error {
	if page > 53 {
		fmt.Println("Outside of page range, last page 52")
		return fmt.Errorf("outside of page range")
	}

	url := cfg.CurrentMap
	newOffsetInt := (page - 1) * 20
	newOffset := strconv.Itoa(newOffsetInt) + "&"

	re := regexp.MustCompile(`\d+&`)
	offsetString := string(re.ReplaceAll([]byte(url), []byte(newOffset)))
	cfg.CurrentMap = offsetString
	return nil
}
