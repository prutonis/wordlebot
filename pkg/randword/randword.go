package randword

import (
	"math/rand"
	"time"

	"moro.io/wordlebot/pkg/utils"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().Unix()))
}

func RandomWord(lang *utils.WbLang) string {
	fileName := utils.WordsFile(lang.Lang)
	rwn := r.Intn(lang.WordCount)
	randomWord, _ := utils.GetWord(fileName, rwn)
	return randomWord
}
