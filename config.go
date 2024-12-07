package main

import "github.com/csullivan94/pokedex/internal/pokecache"

type Config struct {
	Current   string
	Previous  string
	Next      string
	PageNum   int
	CacheUsed bool
	Cache     *pokecache.Cache
	Argument  string
}
