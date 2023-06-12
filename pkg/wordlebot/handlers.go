package wordlebot

import (
	"fmt"
	"strings"

	"github.com/yanzay/tbot/v2"
	"moro.io/wordlebot/pkg/utils"
)

func checkChatCreated(chatId string) *chat {
	Chat, ok := Chats[chatId]
	if !ok {
		Chat = new(chat)
		Chat.chatId = chatId
		Chat.lang = utils.GetConfig().GetDefaultLang()
		Chats[chatId] = Chat
	}
	return Chat
}

func (a *application) helpHandler(m *tbot.Message) {
	a.client.SendMessage(m.Chat.ID, `Wordle Bot menu:
	/help - display help
	/start - start new game
	/giveup - give up and show the number
	/lang - change dictionary language [en, ro]`)
}

func (a *application) startHandler(m *tbot.Message) {
	currentChat := checkChatCreated(m.Chat.ID)
	fmt.Println("starthandler", m.Text)
	Game := currentChat.game
	if Game == nil || Game.isEnded() {
		Game = CreateGame(currentChat.lang)
		fmt.Println("Game created ", Game, m.Chat.ID)
		currentChat.game = Game
		a.client.SendMessage(m.Chat.ID, "A new game started. Try to guess the 5 letter word!")
	} else {
		a.client.SendMessage(m.Chat.ID, "A game is in progress. Continue it or /giveup")
		msg := ListGuesses(Game)
		a.client.SendMessage(m.Chat.ID, msg, tbot.OptParseModeMarkdown)
	}
}

func (a *application) giveUpHandler(m *tbot.Message) {
	Chat := checkChatCreated(m.Chat.ID)
	Game := Chat.game
	fmt.Println("giveup handler")
	if Game != nil {
		if Game.isEnded() {
			a.client.SendMessage(m.Chat.ID, fmt.Sprintf("Game is ended. The word was '%s'. 🚀 Start new one with /start", Game.word))
		} else {
			Game.gc = 10
			a.client.SendMessage(m.Chat.ID, fmt.Sprintf("😎 👎 You are weak. Read more books!\nThe word was '%s'\n%s", Game.word, Chat.lang.GetDictUrl(Game.word)))
		}
	} else {
		a.client.SendMessage(m.Chat.ID, "No game found. 🚀 Start one with /start")
	}
}

func (a *application) messagesHandler(m *tbot.Message) {
	Game := checkChatCreated(m.Chat.ID).game
	uw := m.Text
	fmt.Println("Received: ", m.Text)
	if Game != nil && !Game.isEnded() && len([]rune(uw)) == 5 {
		Game.guesses[Game.gc] = uw
		Game.gc++
		msg := ListGuesses(Game)
		a.client.SendMessage(m.Chat.ID, msg, tbot.OptParseModeMarkdown)
	}
}

func (a *application) languageHandler(m *tbot.Message) {
	cfg := utils.GetConfig()
	lineBtns := make([]tbot.InlineKeyboardButton, len(cfg.Langs))
	i := 0
	for k, _ := range cfg.Langs {
		lineBtns[i] = tbot.InlineKeyboardButton{
			Text:         strings.ToUpper(k),
			CallbackData: k,
		}
		i++
	}
	buttons := tbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]tbot.InlineKeyboardButton{
			lineBtns,
		},
	}
	a.client.SendMessage(m.Chat.ID, "Pick a language", tbot.OptInlineKeyboardMarkup(&buttons))
}

func (a *application) callbackHandler(cq *tbot.CallbackQuery) {
	lang := cq.Data
	currentChat := checkChatCreated(cq.Message.Chat.ID)
	currentChat.lang = utils.GetConfig().Langs[lang]
	a.client.DeleteMessage(cq.Message.Chat.ID, cq.Message.MessageID)
	a.client.SendMessage(cq.Message.Chat.ID, "Selected dictionary language: "+lang)
}

func ListGuesses(Game *game) string {
	vw := strings.Split(Game.word, "")
	sb := strings.Builder{}
	for i := 0; i < Game.gc; i++ {
		gw := strings.Split(Game.guesses[i], "")
		gl := make(map[string]int)
		sb.WriteString("")
		for j, gwl := range gw {
			cl := fmt.Sprintf("*%s* ", gwl)
			if !WordContains(vw, gwl, Game.letters, gl) {
				cl = "*\\#* "
			} else if vw[j] != gw[j] {
				cl = fmt.Sprintf("_%s_ ", gwl)
			}
			sb.WriteString(cl)
		}
		fmt.Println(sb.String())
		sb.WriteString("\n")
	}
	return sb.String()
}

func WordContains(wv []string, letter string, wl map[string]int, gl map[string]int) bool {
	mlc, ok := gl[letter]
	if !ok {
		mlc = 0
	}
	mlc++
	gl[letter] = mlc

	for _, l := range wv {
		if l == letter && mlc <= wl[letter] {
			return true
		}
	}
	return false
}
