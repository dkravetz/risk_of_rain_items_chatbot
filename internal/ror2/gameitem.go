package ror2

// GameItem represents the tables from the wiki
type GameItem struct {
	Name        string
	Description string
	Cooldown    string
}

//GameItems is an array of GameItem
type GameItems []GameItem

//NewGameItem is a constructor for GameItem
func NewGameItem(name, description, cooldown string) *GameItem {
	return &GameItem{Name: name, Description: description, Cooldown: cooldown}
}
