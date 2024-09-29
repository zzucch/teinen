package parse

import (
	"errors"
	"unicode"
)

type Entry struct {
	Word    string
	Meaning string
	Info    string
}

func Parse(data string) ([]Entry, error) {
	entries := make([]Entry, 0)
	dataRunes := []rune(data)

	for currentLocation := 0; currentLocation < len(dataRunes); {
		if !unicode.IsDigit(dataRunes[currentLocation]) {
			currentLocation++
			continue
		}

		entry, err := parseEntry(dataRunes, &currentLocation)
		if err != nil {
			return nil, err
		}

		entries = append(entries, entry)
	}

	return entries, nil
}

func parseEntry(data []rune, currentLocation *int) (Entry, error) {
	_, err := parseNumber(data, currentLocation)
	if err != nil {
		return Entry{}, err
	}

	word, err := parseWord(data, currentLocation)
	if err != nil {
		return Entry{}, err
	}

	meaning, err := parseMeaning(data, currentLocation)
	if err != nil {
		return Entry{}, err
	}

	info, err := parseInfo(data, currentLocation)
	if err != nil {
		return Entry{}, err
	}

	return Entry{
		Word:    word,
		Meaning: meaning,
		Info:    info,
	}, nil
}

func parseNumber(data []rune, currentLocation *int) (string, error) {
	start := *currentLocation
	for *currentLocation < len(data) &&
		unicode.IsDigit(data[*currentLocation]) {
		*currentLocation++
	}

	for *currentLocation < len(data) &&
		unicode.IsSpace(data[*currentLocation]) {
		*currentLocation++
	}

	if start == *currentLocation {
		return "", errors.New("expected number")
	}

	return string(data[start:*currentLocation]), nil
}

func parseWord(data []rune, currentLocation *int) (string, error) {
	start := *currentLocation

	insideBrackets := false
	for *currentLocation < len(data) {
		r := data[*currentLocation]

		if r == '（' {
			insideBrackets = true
			*currentLocation++
			continue
		}

		if r == '）' {
			insideBrackets = false
			*currentLocation++
			continue
		}

		if !insideBrackets && (unicode.IsSpace(r)) {
			break
		}

		*currentLocation++
	}

	if start == *currentLocation {
		return "", errors.New("expected word")
	}

	word := bracketsToAnkiFurigana(data[start:*currentLocation])

	for *currentLocation < len(data) &&
		unicode.IsSpace(data[*currentLocation]) {
		*currentLocation++
	}

	return string(word), nil
}

func parseMeaning(data []rune, currentLocation *int) (string, error) {
	start := *currentLocation

	for *currentLocation < len(data) &&
		data[*currentLocation] != '\n' {
		*currentLocation++
	}

	if start == *currentLocation {
		return "", errors.New("expected meaning")
	}

	*currentLocation++

	return string(data[start:*currentLocation]), nil
}

func parseInfo(data []rune, currentLocation *int) (string, error) {
	start := *currentLocation

	for *currentLocation < len(data) {
		r := data[*currentLocation]

		if unicode.IsDigit(r) &&
			(start == 0 || data[*currentLocation-1] == '\n') {
			break
		}

		*currentLocation++
	}

	return string(data[start:*currentLocation]), nil
}

func bracketsToAnkiFurigana(data []rune) []rune {
	for i := range data {
		if data[i] == '（' || data[i] == '(' {
			data[i] = '['
		}

		if data[i] == '）' || data[i] == ')' {
			data[i] = ']'
		}
	}

	return data
}
