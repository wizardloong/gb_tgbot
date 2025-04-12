package main

import (
	"log"
	"os"

	"github.com/tucnak/telebot"
	"gb_tgbot/infrastructure"
	"gb_tgbot/repository"
	"gb_tgbot/delivery"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

var bot *telebot.Bot

func main() {
	// Загрузить .env
	err := godotenv.Load() // Ищет .env в текущей директории
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Читаем токен бота из переменной окружения
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN не установлен")
	}

	// Создаем экземпляр бота
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

    infrastructure.InitDB()

    repo := repository.NewGoodsRepository(infrastructure.DB)

	// Регистрируем обработчики команд
	var handler = delivery.NewBotHandler(bot, repo)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		handler.HandleUpdates(update)
	}
}
