package main
import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"io/ioutil"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))

	f, err := os.Open("/tmp/dat")
	check(err)

	var sName, number string
	_,err:=f.Read(sName)
	_,err:=f.Read(number)


	var mapmap = map[string]string{}

	// подключаемся к боту с помощью токена
	bot, err := tgbotapi.NewBotAPI("442632858:AAGT6aDU-axkUJIyQ1M6dmAwNustMGfcPEA")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// инициализируем канал, куда будут прилетать обновления от API
	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	updates, _ := bot.GetUpdatesChan(ucfg)
	// читаем обновления из канала
	for {
		for update := range updates{
			// Пользователь, который написал боту
			if update.Message == nil {
				continue
			}
			UserName := update.Message.From.UserName

			// ID чата/диалога.
			// Может быть идентификатором как чата с пользователем
			// (тогда он равен UserID) так и публичного чата/канала
			ChatID := update.Message.Chat.ID

			// Текст сообщения
			Text := update.Message.Text

			log.Printf("[%s] %d %s", UserName, ChatID, Text)

			switch update.Message.Command() {
			case "get":
				reply := "Привет. Я телеграм-бот"
				msg:=tgbotapi.NewMessage(ChatID, reply)
				bot.Send(msg)
			}

			// Ответим пользователю его же сообщением
			//reply := Text
			// Созадаем сообщение
			//msg := tgbotapi.NewMessage(ChatID, reply)
			// и отправляем его

		}

	}
}