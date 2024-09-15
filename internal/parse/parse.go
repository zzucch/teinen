package parse

import (
	"strconv"
	"strings"
)

type Entry struct {
	Word    string
	Meaning string
	Info    string
}

func Parse(lines []string) []Entry {
	entries := make([]Entry, 0)

	var (
		currentEntry     Entry
		parsingInfo      bool
		foundFirstNumber bool
	)

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if !foundFirstNumber {
			if _, err := strconv.Atoi(line); err == nil {
				foundFirstNumber = true
			}

			continue
		}

		if line == "" {
			continue
		}

		if _, err := strconv.Atoi(line); err == nil {
			if parsingInfo && currentEntry.Word != "" {
				entries = append(entries, currentEntry)
				currentEntry = Entry{}
				parsingInfo = false
			}

			continue
		}

		if !parsingInfo {
			if currentEntry.Word == "" {
				currentEntry.Word = line
			} else if currentEntry.Meaning == "" {
				currentEntry.Meaning = line
				parsingInfo = true
			}
		} else {
			if currentEntry.Info == "" {
				currentEntry.Info = line
			} else {
				currentEntry.Info += "\n" + line
			}
		}
	}

	if currentEntry.Word != "" {
		entries = append(entries, currentEntry)
	}

	return entries
}
