package wordlebot

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/yanzay/tbot/v2"
	"moro.io/wordlebot/pkg/randword"
	"moro.io/wordlebot/pkg/utils"
)

type game struct {
	word    string
	letters map[string]int
	guesses [6]string
	gc      int
}

type chat struct {
	chatId string
	game   *game
	lang   *utils.WbLang
}

type chats map[string]*chat

type application struct {
	client *tbot.Client
}

var (
	App   application
	Bot   *tbot.Server
	Token string
	Chats chats
)

func init() {
	e := godotenv.Load()
	if e != nil {
		panic("Cannot load .env file")
	}
	Token = os.Getenv("TBOT_TOKEN")
}

func WordleBot() {
	Chats = make(chats)
	Bot = tbot.New(Token)
	App.client = Bot.Client()
	Bot.HandleMessage("/start", App.startHandler)
	Bot.HandleMessage("/giveup", App.giveUpHandler)
	Bot.HandleMessage(".{5}$", App.messagesHandler)
	//Bot.HandleCallback(App.callbackHandler)
	log.Fatal(Bot.Start())
}

func (g game) isGuessed() bool {
	for _, w := range g.guesses {
		if w == g.word {
			return true
		}
	}
	return false
}

func CreateGame(lang *utils.WbLang) *game {
	word := randword.RandomWord(lang)
	g := &game{word: word, letters: LetterCountMap(word), gc: 0}
	return g
}

func (g game) isEnded() bool {
	return g.gc > 5 || g.isGuessed()
}

// func (g game) isGivenUp() bool {
// 	return g.gc >= 10
// }

func LetterCountMap(word string) map[string]int {
	wordLetters := make(map[string]int)
	for _, letter := range strings.Split(word, "") {
		count, ok := wordLetters[letter]
		if !ok {
			count = 0
		}
		count++
		wordLetters[letter] = count
	}
	return wordLetters
}
