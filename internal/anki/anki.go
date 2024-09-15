package anki

import (
	"strings"

	"github.com/atselvan/ankiconnect"
	"github.com/zzucch/teinen/internal/waitlog"
)

func Run() {
	client := ankiconnect.NewClient()

	restErr := client.Ping()
	if restErr != nil {
		waitlog.Fatal(restErr)
	}

	newDeckName := "new-deck"

	decks, restErr := client.Decks.GetAll()

	if restErr != nil {
		waitlog.Fatal(restErr)
	}

	for _, deck := range *decks {
		if strings.EqualFold(newDeckName, deck) {
			waitlog.Fatal("the deck already exists, delete it")
		}
	}

	if restErr := client.Decks.Create(newDeckName); restErr != nil {
		waitlog.Fatal()
	}

	note := ankiconnect.Note{
		DeckName:  newDeckName,
		ModelName: "Basic",
		Fields: ankiconnect.Fields{
			"Front": "Front data",
			"Back":  "Back data",
		},
	}

	if restErr := client.Notes.Add(note); restErr != nil {
		waitlog.Fatal(restErr.Error)
	}

	waitlog.Println("Done")
}
