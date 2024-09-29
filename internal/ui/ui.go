package ui

import (
	"fmt"
	"strings"

	"github.com/atselvan/ankiconnect"
	"github.com/zzucch/teinen/internal/anki"
	"github.com/zzucch/teinen/internal/parse"
	"github.com/zzucch/teinen/internal/read"
	"github.com/zzucch/teinen/internal/waitlog"
)

func GetModelAndFields(client *ankiconnect.Client) (string, []anki.ModelField) {
	models := anki.GetModels(client)
	model := Choose("model", models)

	modelFields := anki.GetModelFields(client, model)

	const fieldAmount = 3
	chosenFields := make([]anki.ModelField, 0, fieldAmount)

	chosenFields = append(chosenFields, anki.ModelField{
		FieldType: anki.Word,
		Name:      Choose("word field", modelFields),
	})
	chosenFields = append(chosenFields, anki.ModelField{
		FieldType: anki.Meaning,
		Name:      Choose("meaning field", modelFields),
	})
	chosenFields = append(chosenFields, anki.ModelField{
		FieldType: anki.Info,
		Name:      Choose("info field", modelFields),
	})

	return model, chosenFields
}

func GetEntries() []parse.Entry {
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

func CreateAndPopulateDeck(
	client *ankiconnect.Client,
	model string,
	modelFields []anki.ModelField,
	entries []parse.Entry,
) {
	const deckName = "teinen"

	anki.CreateDeck(client, deckName)

	fmt.Println("inserting...")

	for _, entry := range entries {
		fieldData := []anki.FieldData{
			{
				Field: anki.GetField(modelFields, anki.Word).Name,
				Data:  entry.Word,
			},
			{
				Field: anki.GetField(modelFields, anki.Meaning).Name,
				Data:  entry.Meaning,
			},
			{
				Field: anki.GetField(modelFields, anki.Info).Name,
				Data:  entry.Info,
			},
		}

		anki.AddNote(client, deckName, model, fieldData)
	}
}
