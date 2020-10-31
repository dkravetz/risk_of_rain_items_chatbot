package ror2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// GameItem represents the tables from the wiki
type GameItem struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Cooldown     string `json:"cooldown"`
	AcquiredFrom string `json:"acquired_from"`
}

//GameItems is an array of GameItem
type GameItems []GameItem

//NewGameItem is a constructor for GameItem
func NewGameItem(name, description, cooldown, acquiredFrom string) GameItem {
	return GameItem{Name: name, Description: description, Cooldown: cooldown, AcquiredFrom: acquiredFrom}
}

// Stringer method
func (g GameItem) String() string {
	return fmt.Sprintf("%s: %s Scaling/Cooldown: %s", g.Name, g.Description, g.Cooldown)
}

// String (int) is necessary to satisfy the Source type for the fuzzy finder
func (g GameItems) String(i int) string {
	return g[i].Name
}

// Len is necessary to satisfy the Source type for the fuzzy finder
func (g GameItems) Len() int {
	return len(g)
}

// ToJSON stores an array of GameItems into a json file on disk.
// It creates or replaces the specified file.
func (g GameItems) ToJSON(filePath string) error {
	items, err := json.MarshalIndent(g, "", " ")
	if err != nil {
		log.Fatal("Couldn't convert GameItems to json", err)
		return err
	}

	err = ioutil.WriteFile(filePath, items, 0644)
	if err != nil {
		log.Fatal("Couldn't save json file", err)
	}

	return err
}

// FromJSON parses a json file located at filePath
func FromJSON(filePath string) (GameItems, error) {
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("Couldn't read json file", err)
		return nil, err
	}

	var gameItems GameItems
	err = json.Unmarshal(f, &gameItems)
	if err != nil {
		log.Fatal("Couldn't parse json file", err)
		return nil, err
	}

	return gameItems, nil
}
