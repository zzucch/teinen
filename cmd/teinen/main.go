package main

import (
	"fmt"
	"strings"

	"github.com/atselvan/ankiconnect"
	"github.com/zzucch/teinen/internal/anki"
	"github.com/zzucch/teinen/internal/parse"
	"github.com/zzucch/teinen/internal/read"
	"github.com/zzucch/teinen/internal/ui"
	"github.com/zzucch/teinen/internal/waitlog"
)

func main() {
	// client := anki.Connect()

	// model, modelFields := getModelAndFields(client)

	entries := getEntries()
	for _, entry := range entries {
		fmt.Printf(
			"\n\nEntry:\nWord:%s\nMeaning:%s\nInfo:%s",
			entry.Word,
			entry.Meaning,
			entry.Info,
		)
	}

	// createAndPopulateDeck(client, model, modelFields, entries)
}

func getModelAndFields(client *ankiconnect.Client) (string, []anki.ModelField) {
	models := anki.GetModels(client)
	model := ui.Choose("model", models)

	modelFields := anki.GetModelFields(client, model)

	const fieldAmount = 3
	chosenFields := make([]anki.ModelField, 0, fieldAmount)

	chosenFields = append(chosenFields, anki.ModelField{
		FieldType: anki.Word,
		Name:      ui.Choose("word field", modelFields),
	})
	chosenFields = append(chosenFields, anki.ModelField{
		FieldType: anki.Meaning,
		Name:      ui.Choose("meaning field", modelFields),
	})
	chosenFields = append(chosenFields, anki.ModelField{
		FieldType: anki.Info,
		Name:      ui.Choose("info field", modelFields),
	})

	return model, chosenFields
}

func getEntries() []parse.Entry {
	for {
		entries := readAndParse()

		fmt.Printf("found %d entries\n", len(entries))

		if len(entries) == 0 {
			fmt.Println("try once more")
		} else {
			return entries
		}
	}
}

func readAndParse() []parse.Entry {
	lines, err := read.Read()
	if err != nil {
		waitlog.Fatal(err)
	}

	entries, err := parse.Parse(strings.Join(lines, "\n"))
	if err != nil {
		waitlog.Fatal(err)
	}

	return entries
}

func createAndPopulateDeck(
	client *ankiconnect.Client,
	model string,
	modelFields []anki.ModelField,
	entries []parse.Entry,
) {
	const deckName = "teinen"

	anki.CreateDeck(client, deckName)

	for _, entry := range entries {
		fieldData := []anki.FieldData{
			{
				Field: getField(modelFields, anki.Word).Name,
				Data:  entry.Word,
			},
			{
				Field: getField(modelFields, anki.Meaning).Name,
				Data:  entry.Meaning,
			},
			{
				Field: getField(modelFields, anki.Info).Name,
				Data:  entry.Info,
			},
		}
		anki.AddNote(client, deckName, model, fieldData)
	}
}

func getField(
	modelFields []anki.ModelField,
	fieldType anki.FieldType,
) *anki.ModelField {
	for _, field := range modelFields {
		if field.FieldType == fieldType {
			return &field
		}
	}

	waitlog.Fatal("failed to get field")
	return nil // never happends
}
