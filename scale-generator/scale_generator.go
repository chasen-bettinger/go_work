package scale

import (
	"fmt"
	"strings"
)

var flats = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
var sharps = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}

// parse a sequence of intervals to absolute semitones above the tonic
func getNotes(interval string) []int {
	if interval == "" {
		return []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	}
	notes := []int{}
	currentNote := 0
	intervalScale := map[rune]int{'m': 1, 'M': 2, 'A': 3}
	for _, currentInterval := range interval {
		notes = append(notes, currentNote)
		currentNote += intervalScale[currentInterval]
	}
	return notes
}

//find the index of a tonic (as major or minor) in a list
func tonicIndex(tonic string, chromatic []string) int {
	for i := range chromatic {
		if strings.ToUpper(chromatic[i]) == strings.ToUpper(tonic) {
			return i
		}
	}
	panic(fmt.Sprintf("tonic %v not found", tonic))
}

//Scale returns the list of notes in scale given by a tonic and a set of intervals
func Scale(tonic string, interval string) []string {
	notes := getNotes(interval)

	var chromatic []string
	switch tonic {
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		chromatic = flats
	default:
		chromatic = sharps
	}

	scale := []string{}
	tonicI := tonicIndex(tonic, chromatic)
	for i := range notes {
		currentNote := (tonicI + notes[i]) % 12
		currentScale := chromatic[currentNote]
		scale = append(scale, currentScale)
	}
	return scale
}
