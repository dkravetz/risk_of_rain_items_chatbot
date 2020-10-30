package ror2

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// GetAllItemsFromFile parses an html file instead of doing a network request
func GetAllItemsFromFile(filePath string) (GameItems, error) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Couldn't open file. ", err)
	}
	defer f.Close()

	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal("Couldn't parse html. ", err)
	}

	return getItemsDocument(doc), err
}

// GetAllItemsFromURL does a network request to the specified URL and parses the results
func GetAllItemsFromURL(url string) (GameItems, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal("Couldn't request URL. ", err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("Couldn't parse html. ", err)
	}
	return getItemsDocument(doc), err
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
