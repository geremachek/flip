package main

import (
	"strings"
	"io/ioutil"
)

const tab = "    "

// Construct a "deck" of files from a list of paths

func getDeck(files []string) ([][]string, error) {
	var deck [][]string

	for _, f := range files {
		data, err := ioutil.ReadFile(f)

		if err != nil {
			return nil, err
		}

		deck = append(deck, processSlide(string(data))...)
	}

	return deck, nil
}

// process a slide (splitting it into multiple slides if needed) for display

func processSlide(slide string) (slides [][]string) {
	// remove trailing newlines (these will take longer to render anyway)

	trimmed := strings.TrimRight(slide, "\n")

	// make sure tabs will render properly :)

	trimmed = strings.ReplaceAll(trimmed, "\t", tab)

	// split up individual files into "slides" with the "--SLIDE--" token

	for _, slide := range strings.Split(trimmed, "\n--SLIDE--\n") {
		slides = append(slides, strings.Split(slide, "\n"))
	}

	return
}
