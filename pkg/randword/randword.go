package randword

import (
	"math/rand"

	"moro.io/wordlebot/pkg/utils"
)

func RandomWord() string {
	cfg := utils.GetConfig()
	fileName := utils.WordsFile(cfg.Lang.Lang)
	rwn := rand.Intn(cfg.Lang.WordCount)
	randomWord, _ := utils.GetWord(fileName, rwn)
	return randomWord
}
