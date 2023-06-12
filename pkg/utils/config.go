package utils

import (
	"fmt"
	"net/url"
	"sync"

	"github.com/spf13/viper"
	"moro.io/wordlebot/pkg/logger"
)

type WbLang struct {
	Lang      string
	WordCount int
	DictUrl   string
}

type WbConfig struct {
	Langs       map[string]*WbLang
	strictWords bool
}

var instance *WbConfig
var once sync.Once

func init() {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("wordlebot") // Register config file name (no extension)
	viper.SetConfigType("toml")      // Look for specific type
	viper.ReadInConfig()
	logger.Info("Config loaded.")
}

func GetConfig() *WbConfig {
	once.Do(func() {
		langsMap := viper.GetStringMapString("langs")
		instance = &WbConfig{Langs: make(map[string]*WbLang, len(langsMap)), strictWords: false}
		for k, v := range langsMap {
			wc := LineCounter(WordsFile(k))
			lang := WbLang{Lang: k, WordCount: wc, DictUrl: v}
			//instance.Langs = append(instance.Langs, lang)
			instance.Langs[lang.Lang] = &lang
		}
		if len(instance.Langs) == 0 {
			panic("No word files found!")
		}
	})
	return instance
}

func (wc WbConfig) GetDefaultLang() *WbLang {
	return wc.Langs["ro"]
}

func (wl WbLang) GetDictUrl(word string) string {
	return fmt.Sprintf("%s%s", wl.DictUrl, url.QueryEscape(word))
}
