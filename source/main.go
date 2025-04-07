package main

import (
	"log"
	"os"
	"time"

	"github.com/tucnak/telebot"
	"github.com/gb_tgbot/source/infrastructure"
	"github.com/gb_tgbot/source/repository"
	"github.com/gb_tgbot/source/delivery"
)

var bot *telebot.Bot

func main() {
	// Читаем токен бота из переменной окружения
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN не установлен")
	}

	// Создаем экземпляр бота
	var err error
	bot, err = telebot.NewBot(telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}

    infrastructure.InitDB()

    repo := repository.NewGoodsRepository(infrastructure.DB)

	// Регистрируем обработчики команд
	var handler = delivery.NewBotHandler(bot, repo)
    bot.Handle(telebot.OnText, handler.HandleUpdates)

	// Запускаем бота
	log.Println("Бот запущен!")
	bot.Start()
}
