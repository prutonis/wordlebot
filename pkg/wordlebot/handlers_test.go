package wordlebot

import (
	"testing"
)

func TestGetCurrentKeyboard(t *testing.T) {
	game := &game{
		word:    "elude",
		gc:      3,
		guesses: [10]string{"aeiuo", "trump", "barak"},
		letters: map[string]int{"e": 2, "l": 1, "u": 1, "d": 1},
	}
	out := `*\\#* _e_ *\\#* _u_ *\\#* \n*\\#* *\\#* *u* *\\#* *\\#* \n*\\#* *\\#* *\\#* *\\#* *\\#* \n`
	gl := ListGuesses(game)
	if gl != out {
		t.Fatal("Something has gone wrong")
	}
}
