package main

import (
	"fmt"

	"github.com/dkravetz/risk_of_rain_items_chatbot/internal/ror2"
)

func main() {
	items, err := ror2.GetAllItemsFromFile("files/ror2_items.html")
	if err != nil {
		panic(err)
	}
	fmt.Println(items[83])

}
