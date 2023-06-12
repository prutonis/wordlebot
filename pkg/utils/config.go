package utils

import (
	"sync"

	"github.com/spf13/viper"
	"moro.io/wordlebot/pkg/logger"
)

type WbLang struct {
	Lang      string
	WordCount int
}

type WbConfig struct {
	Langs []WbLang
	Lang  *WbLang
}

var instance *WbConfig
var once sync.Once

func LoadConfig() {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("wordlebot") // Register config file name (no extension)
	viper.SetConfigType("toml")      // Look for specific type
	viper.ReadInConfig()
	logger.Info("Config loaded.")
}

func GetConfig() *WbConfig {
	once.Do(func() {
		LoadConfig()
		instance = &WbConfig{}
		langs := viper.GetStringSlice("words.langs")
		for _, l := range langs {
			wc := LineCounter(WordsFile(l))
			lang := WbLang{Lang: l, WordCount: wc}
			instance.Langs = append(instance.Langs, lang)
		}
		if len(instance.Langs) == 0 {
			panic("No word files found!")
		}
		instance.Lang = &instance.Langs[0]
	})
	return instance
}
