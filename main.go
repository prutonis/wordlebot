package main

import (
	"moro.io/wordlebot/pkg/logger"
	"moro.io/wordlebot/pkg/randword"
)

func main() {
	// x, _ := randword.RandomWord("configs/ro.txt")
	// y, _ := randword.RandomWord("configs/en.txt")
	// fmt.Printf("ro: %s, en: %s\n", x, y)

	logger.Info("Random word: " + randword.RandomWord())
}
