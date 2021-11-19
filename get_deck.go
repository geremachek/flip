package main

import (
	"os"
	"io"
	"bufio"
	"strings"
)

const tab = "    "

// Construct a "deck" of files from a list of paths

func getDeck(files []string) ([][]string, error) {
	var deck [][]string

	for _, f := range files {
		reader, err := os.Open(f)

		if err != nil {
			return nil, err
		}

		deck = append(deck, processSlide(reader)...)
	}

	return deck, nil
}

// process a slide (splitting it into multiple slides if needed) for display

func processSlide(rd io.Reader) [][]string {
	var (
		scanner = bufio.NewScanner(rd)
		
		slides = make([][]string, 1)
		i int
	)

	for scanner.Scan() {
		// separate stream into different slides based on the "--SLIDE--" token

		if line := scanner.Text(); line == "--SLIDE--" {
			i++
			slides = append(slides, []string{})
		} else {
			// add the line to the slide + tab display hack

			slides[i] = append(slides[i], strings.ReplaceAll(line, "\t", tab))
		}
	}

	return slides
}
