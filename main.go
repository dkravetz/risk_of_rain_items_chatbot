package main

import (
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
)

type GameItem struct {
	Name        string
	Description string
	Cooldown    string
}

func New(name, description, cooldown string) *GameItem {
	return &GameItem{Name: name, Description: description, Cooldown: cooldown}
}

type GameItems []GameItem

func main() {
	f, err := os.Open("files/ror2_items.html")
	if err != nil {
		log.Fatal("Couldn't open file. ", err)
		panic("Cannot continue.")
	}
	defer f.Close()

	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal("Couldn't parse html. ", err)
		panic("Cannot continue.")
	}

	// Get rows into an array. each element represents a GameItem
	var gameItems GameItems
	var columns []string
	var rows [][]string
	doc.Find(".article-table").Each(func(i int, table *goquery.Selection) {
		table.Find("tbody tr").Each(func(j int, tr *goquery.Selection) {
			tr.Find("td").Each(func(index int, cell *goquery.Selection) {
				columns = append(columns, cell.Text())
			})
			rows = append(rows, columns)
			columns = nil // Otherwise we append all previous columns to current one
		})
	})

	for _, item := range rows {
		gameItems = append(gameItems, GameItem{item[0], item[1], item[2]})
	}

}
