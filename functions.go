package functions

import (
	"log"

	"github.com/vitormoura/libs-lab-gomodules/messages"
)

// GetRandomStringList récupère une liste de strings
func GetRandomStringList() []string {
	const qttMessages = 10

	var phrase messages.Phrase
	var err error

	msgs := make([]string, qttMessages)

	for i := 0; i < qttMessages; i++ {
		if phrase, err = messages.GetRandomPhrase(); err != nil {
			log.Fatal(err)
		}

		msgs = append(msgs, phrase.Text)
	}

	return msgs
}
