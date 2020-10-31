package ror2

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// UpdateItemListFromURL does a network request to the specified URL and parses the results
func UpdateItemListFromURL(url string) (GameItems, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal("Couldn't request URL. ", err)
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("Couldn't parse html. ", err)
		return nil, err
	}
	gameItems := getItemsDocument(doc)
	gameItems.ToJSON("items.json")
	return gameItems, err
}

func getItemsDocument(doc *goquery.Document) GameItems {
	// Get rows into an array. each element represents a GameItem
	var gameItems GameItems
	var columns []string
	var rows [][]string
	doc.Find(".article-table").Each(func(i int, table *goquery.Selection) {
		table.Find("tbody tr").Each(func(j int, tr *goquery.Selection) {
			tr.Find("td").Each(func(index int, cell *goquery.Selection) {
				columns = append(columns, strings.TrimRight(cell.Text(), "\n")) // remove trailing newline
			})
			// TODO: Find out where this nils are coming from
			if columns != nil {
				rows = append(rows, columns)
				columns = nil // Otherwise we append all previous columns to current one
			}
		})
	})

	for i, item := range rows {
		// Ugly workaround for the new column explaining where/how to get boss items
		if len(rows[i]) == 4 {
			gameItems = append(gameItems, NewGameItem(item[0], item[1], item[3], item[2]))
			continue
		}
		gameItems = append(gameItems, NewGameItem(item[0], item[1], item[2], ""))
	}
	return gameItems
}
