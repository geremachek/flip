package main

import (
	"strings"
	"io/ioutil"
)

const tab = "    "

// Construct a "deck" of files from a list of paths

func getDeck(files []string) ([][]string, error) {
	var (
		deck [][]string
		trimmed string
	)

	for _, f := range files {
		data, err := ioutil.ReadFile(f)

		if err != nil {
			return nil, err
		}

		// remove trailing newlines (these will take longer to render anyway)

		trimmed = strings.TrimRight(string(data), "\n")

		// make sure tabs will render properly :)

		trimmed = strings.ReplaceAll(trimmed, "\t", tab)

		// split up individual files into "slides" with the "--SLIDE--" token

		for _, slide := range strings.Split(trimmed, "\n--SLIDE--\n") {
			deck = append(deck, strings.Split(slide, "\n"))
		}
	}

	return deck, nil
}
