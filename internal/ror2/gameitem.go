package ror2

import "fmt"

// GameItem represents the tables from the wiki
type GameItem struct {
	Name         string
	Description  string
	Cooldown     string
	AcquiredFrom string
}

//GameItems is an array of GameItem
type GameItems []GameItem

//NewGameItem is a constructor for GameItem
func NewGameItem(name, description, cooldown, acquiredFrom string) GameItem {
	return GameItem{Name: name, Description: description, Cooldown: cooldown, AcquiredFrom: acquiredFrom}
}

func (g GameItem) String() string {
	return fmt.Sprintf("%s: %s Scaling/Cooldown: %s", g.Name, g.Description, g.Cooldown)
}
