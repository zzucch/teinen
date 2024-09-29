package main

import (
	"github.com/zzucch/teinen/internal/anki"
	"github.com/zzucch/teinen/internal/ui"
	"github.com/zzucch/teinen/internal/waitlog"
)

func main() {
	client := anki.Connect()

	model, modelFields := ui.GetModelAndFields(client)

	entries := ui.GetEntries()

	ui.CreateAndPopulateDeck(client, model, modelFields, entries)

	waitlog.Println("done")
}
