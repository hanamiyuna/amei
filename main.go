package main

import (
	"io/ioutil"
	"log"
	"path/filepath"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gopkg.in/yaml.v2"
)

// Configuration Interface
type config struct {
	Token string `yaml:"token"`
}

// Read in conf
func (confFile *config) getConf() *config {
	fileName, _ := filepath.Abs("./conf.yaml")
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panic(err)
	}

	err = yaml.Unmarshal(yamlFile, confFile)
	return confFile
}

func main() {

	var c config
	conf := c.getConf()

	bot, err := tgbotapi.NewBotAPI(conf.Token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Connected @%s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
