package anki

import (
	"log"
	"strings"

	"github.com/atselvan/ankiconnect"
	"github.com/zzucch/teinen/internal/waitlog"
)

type FieldType int

const (
	Word FieldType = iota
	Meaning
	Info
)

type FieldData struct {
	Field string
	Data  string
}

type ModelField struct {
	FieldType
	Name string
}

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
			waitlog.Println("the deck already exists, delete it")
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

func GetModelFields(
	client *ankiconnect.Client,
	modelName string,
) *[]string {
	fields, restErr := client.Models.GetFields(modelName)
	if restErr != nil {
		waitlog.Fatal(restErr)
	}

	return fields
}

func AddNote(
	client *ankiconnect.Client,
	deckName, modelName string,
	fieldData []FieldData,
) {
	if len(fieldData) != 3 {
		waitlog.Fatal("invalid fields amount")
	}

	note := ankiconnect.Note{
		DeckName:  deckName,
		ModelName: modelName,
		Fields: ankiconnect.Fields{
			fieldData[0].Field: fieldData[0].Data,
			fieldData[1].Field: fieldData[1].Data,
			fieldData[2].Field: fieldData[2].Data,
		},
	}

	log.Println(note)

	if restErr := client.Notes.Add(note); restErr != nil {
		waitlog.Fatal(restErr.Error)
	}
}
