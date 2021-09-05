package main

import (
	"strings"
	"io/ioutil"
)

// Construct a "deck" of files from a list of paths

func getDeck(files []string) ([][]string, error) {
	var deck [][]string

	for _, f := range files {
		data, err := ioutil.ReadFile(f)

		if err != nil {
			return nil, err
		}

		deck = append(deck, strings.Split(strings.TrimRight(string(data), "\n"), "\n"))
	}

	return deck, nil
}
