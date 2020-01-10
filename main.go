package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	// TODO: formatar o código e permitir carregar dados pelas variáveis de ambiente

	var config Config
	readConfigFile(&config)
	readConfigEnv(&config)
	bot, err := tgbotapi.NewBotAPI(config.Telegram.Botkey)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Autorizado na conta %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			var allowedChatID int64 = config.Telegram.ChatID
			if update.Message.Chat.ID == allowedChatID {
				// NOTA! Não seria melhor organizar cada switch por uma função ou um enum ?
				switch update.Message.Command() {
				case "dump":
					msg.Text = "Um momento, estou tentando realizar o dump."
					bot.Send(msg)
					fileName, err := dumpDatabase(config)
					if err != nil {
						fmt.Println(err)
						bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Ocorreu um erro: "+err.Error()))
						return
					}
					// NOTA! tornar esse path ./dumps/... uma variável para ter apenas um lugar com esse nome.
					// NOTA! organizar esses bot.Send e msg.Text de forma que fique claro o momento de cada um.
					msg := tgbotapi.NewDocumentUpload(update.Message.Chat.ID, "./dumps/"+fileName)
					_, err = bot.Send(msg)
					if err != nil {
						fmt.Println(err)
						bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Ocorreu um erro: "+err.Error()))
						return
					}
					os.Remove("./dumps/" + fileName)
					msgToSend := tgbotapi.NewMessage(update.Message.Chat.ID, "Dump finalizado")
					bot.Send(msgToSend)

				case "chat_id":
					msg.Text = strconv.FormatInt(update.Message.Chat.ID, 10)
					bot.Send(msg)
				default:
					return
				}
			} else {
				msg.Text = "Eu não estou permitido responder você ou seus comandos."
				bot.Send(msg)
			}

		}

	}

}
