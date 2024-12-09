package main

import (
	"github.com/csullivan94/pokedex/internal/pokeapi"
	"github.com/csullivan94/pokedex/internal/pokecache"
)

type Config struct {
	CurrentMap     string
	PreviousMap    string
	NextMap        string
	PageNum        int
	CacheUsed      bool
	Cache          *pokecache.Cache
	Argument       string
	CurrentExplore string
	Pokedex        map[string]pokeapi.Pokemon
}
