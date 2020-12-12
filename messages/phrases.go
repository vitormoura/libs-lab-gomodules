package messages

import (
	"bufio"
	"math/rand"
	"strings"
)

// Phrase represents une phrase avec sa signification
type Phrase struct {
	Text    string
	Meaning string
}

// GetRandomPhrase récupère une phrase au hasard
func GetRandomPhrase() (Phrase, error) {
	var phrases []Phrase
	var err error

	if phrases, err = getListOfPhrases(); err != nil {
		return Phrase{}, err
	}

	pos := rand.Intn(len(phrases))

	return phrases[pos], nil
}

func getListOfPhrases() ([]Phrase, error) {
	data, _ := Asset("data/phrases.txt")

	file := strings.NewReader(string(data))

	scanner := bufio.NewScanner(file)
	text := ""
	phrases := make([]Phrase, 50)
	phrase := Phrase{}

	for scanner.Scan() {
		text = scanner.Text()

		if len(phrase.Text) == 0 {
			phrase.Text = text
		} else {
			phrase.Meaning = text
			phrases = append(phrases, phrase)

			phrase = Phrase{}
		}
	}

	return phrases, nil
}
