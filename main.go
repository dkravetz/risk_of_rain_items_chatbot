package main

import (
	"flag"
	"fmt"

	"github.com/dkravetz/risk_of_rain_items_chatbot/internal/ror2"
	"github.com/sahilm/fuzzy"
)

func main() {
	wordPtr := flag.String("word", "Tri-Tip Dagger", "Name of item to search")
	flag.Parse()

	items, err := ror2.GetAllItemsFromFile("files/ror2_items.html")
	// items, err := ror2.GetAllItemsFromURL("https://riskofrain2.gamepedia.com/Items#Boss")
	if err != nil {
		panic(err)
	}

	results := fuzzy.FindFrom(*wordPtr, items)

	fmt.Println(items[results[0].Index])

}
