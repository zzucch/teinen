package anki

import (
	"strings"

	"github.com/atselvan/ankiconnect"
	"github.com/zzucch/teinen/internal/waitlog"
)

func Connect() *ankiconnect.Client {
	return ankiconnect.NewClient()
}

func CreateDeck(client *ankiconnect.Client, deckName string) {
	decks, restErr := client.Decks.GetAll()
	if restErr != nil {
		waitlog.Fatal(restErr)
	}

	for _, deck := range *decks {
		if strings.EqualFold(deckName, deck) {
			waitlog.Fatal("the deck already exists, delete it")
		}
	}

	if restErr := client.Decks.Create(deckName); restErr != nil {
		waitlog.Fatal(restErr)
	}
}

func GetModels(client *ankiconnect.Client) *[]string {
	models, restErr := client.Models.GetAll()
	if restErr != nil {
		waitlog.Fatal(restErr)
	}

	return models
}

func GetModelFields(client *ankiconnect.Client, modelName string) *[]string {
	fields, restErr := client.Models.GetFields(modelName)
	if restErr != nil {
		waitlog.Fatal(restErr)
	}

	return fields
}

func AddNote(
	client *ankiconnect.Client,
	deckName, modelName string,
) {
	note := ankiconnect.Note{
		DeckName:  deckName,
		ModelName: modelName,
		Fields: ankiconnect.Fields{
			"Front": "Front data",
			"Back":  "Back data",
		},
	}

	if restErr := client.Notes.Add(note); restErr != nil {
		waitlog.Fatal(restErr.Error)
	}
}
