package main

import (
	"strings"
	"io/ioutil"
)

// GetDeck
// Construct a "deck" of files from a list of paths

func GetDeck(files []string) ([][]string, error) {
	var deck [][]string

	for _, f := range files {
		data, err := ioutil.ReadFile(f)

		if err != nil {
			return [][]string{}, err
		}

		deck = append(deck, strings.Split(strings.TrimRight(string(data), "\n"), "\n"))
	}

	return deck, nil
}
