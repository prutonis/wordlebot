package main

import (
	"moro.io/wordlebot/pkg/logger"
	"moro.io/wordlebot/pkg/wordlebot"
)

func main() {
	logger.Info("WordleBot started")
	wordlebot.WordleBot()
}
