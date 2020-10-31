package main

import (
	"flag"
	"fmt"

	"github.com/dkravetz/risk_of_rain_items_chatbot/internal/ror2"
	"github.com/sahilm/fuzzy"
)

func main() {
	wordPtr := flag.String("word", "Tri Tip", "Name of item to search")
	flag.Parse()

	items, err := ror2.FromJSON("items.json")
	// items, err := ror2.UpdateItemListFromURL("https://riskofrain2.gamepedia.com/Items#Boss")
	if err != nil {
		panic(err)
	}

	if results := fuzzy.FindFrom(*wordPtr, items); results != nil {
		fmt.Println(items[results[0].Index])
	} else {
		fmt.Println("Sorry, can't find results")
	}
}
