package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/dkravetz/risk_of_rain_items_chatbot/internal/ror2"
	"github.com/sahilm/fuzzy"
)

func main() {
	itemPtr := flag.String("item", "", "Name of item to search")
	flag.Parse()

	items, err := ror2.FromJSON("items.json")
	// items, err := ror2.UpdateItemListFromURL("https://riskofrain2.gamepedia.com/Items#Boss")
	if err != nil {
		panic(err)
	}

	// Why can't I just `if (*itemPtr){}` ??
	if *itemPtr == "" {
		interactiveMode(&items)
	} else {
		if results := fuzzy.FindFrom(*itemPtr, items); results != nil {
			fmt.Println(items[results[0].Index])
		} else {
			fmt.Println("Sorry, can't find results")
		}
	}
}

func interactiveMode(items *ror2.GameItems) {
	fmt.Println("Type in the name of the item you're looking for:")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "quit" || scanner.Text() == "exit" {
			os.Exit(0)
		}
		if results := fuzzy.FindFrom(scanner.Text(), items); results != nil {
			fmt.Println((*items)[results[0].Index])
		} else {
			fmt.Println("Sorry, can't find results")
		}
		fmt.Println("Type in the name of the item you're looking for:")
	}
}
