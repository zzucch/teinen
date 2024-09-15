package main

import (
	"github.com/zzucch/teinen/internal/anki"
	"github.com/zzucch/teinen/internal/ui"
)

func main() {
	//	lines, err := read.Read()
	//	if err != nil {
	//		waitlog.Fatal(err)
	//	}
	//
	//	entries := parse.Parse(lines)
	//
	//	println("amount: ", len(entries))

	client := anki.Connect()

	// deckName := "new-deck"

	// anki.CreateDeck(client, deckName)
	models := anki.GetModels(client)
	print(ui.Choose(models))
}
