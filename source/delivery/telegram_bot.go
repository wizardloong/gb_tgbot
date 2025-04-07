package delivery

import (
	"fmt"
	"log"
	"strconv"
	"telegram-bot/internal/domain"
	"telegram-bot/internal/repository"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type BotHandler struct {
	bot  *tgbotapi.BotAPI
	repo *repository.GoodsRepository
}

func NewBotHandler(bot *tgbotapi.BotAPI, repo *repository.GoodsRepository) *BotHandler {
	return &BotHandler{bot: bot, repo: repo}
}

func (h *BotHandler) HandleUpdates(update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	switch update.Message.Command() {
	case "add":
		h.handleAdd(update)
	case "list":
		h.handleList(update)
	case "get":
		h.handleGet(update)
	case "update":
		h.handleUpdate(update)
	case "delete":
		h.handleDelete(update)
	default:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Неизвестная команда.")
		h.bot.Send(msg)
	}
}

// Добавление товара: /add имя_цена_описание
func (h *BotHandler) handleAdd(update tgbotapi.Update) {
	args := update.Message.CommandArguments()
	parts := splitArgs(args, 3)
	if parts == nil {
		h.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Формат: /add имя цена описание"))
		return
	}

	price, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		h.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка: цена должна быть числом"))
		return
	}

	goods := domain.Goods{Name: parts[0], Price: price, Data: parts[2], IsActive: true}
	if err := h.repo.Create(&goods); err != nil {
		h.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка при добавлении товара"))
		return
	}

	h.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Товар успешно добавлен!"))
}

// Получение списка товаров: /list
func (h *BotHandler) handleList(update tgbotapi.Update) {
	goods, err := h.repo.GetAll()
	if err != nil {
		h.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка получения списка товаров"))
		return
	}

	if len(goods) == 0 {
		h.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Список товаров пуст"))
		return
	}

	var response string
	for _, g := range goods {
		response += fmt.Sprintf("ID: %d\nНазвание: %s\nЦена: %.2f\nОписание: %s\n\n",
			g.ID, g.Name, g.Price, g.Data)
	}

	h.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, response))
}

// Получение информации о товаре: /get ID
func (h *BotHandler) handleGet(update tgbotapi.Update) {
	id, err := strconv.Atoi(update.Message.CommandArguments())
	if err != nil {
		h.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка: ID должно быть числом"))
		return
	}

	goods, err := h.repo.GetByID(uint(id))
	if err != nil {
		h.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Товар не найден"))
		return
	}

	h.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID,
		fmt.Sprintf("ID: %d\nНазвание: %s\nЦена: %.2f\nОписание: %s",
			goods.ID, goods.Name, goods.Price, goods.Data)))
}

// Обновление товара: /update ID новое_имя новая_цена новое_описание
func (h *BotHandler) handleUpdate(update tgbotapi.Update) {
	args := update.Message.CommandArguments()
	parts := splitArgs(args, 4)
	if parts == nil {
		h.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Формат: /update ID имя цена описание"))
		return
	}

	id, err := strconv.Atoi(parts[0])
	if err != nil {
		h.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка: ID должно быть числом"))
		return
	}

	price, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		h.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка: цена должна быть числом"))
		return
	}

	goods, err := h.repo.GetByID(uint(id))
	if err != nil {
		h.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Товар не найден"))
		return
	}

	goods.Name = parts[1]
	goods.Price = price
	goods.Data = parts[3]

	if err := h.repo.Update(goods); err != nil {
		h.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка при обновлении товара"))
		return
	}

	h.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Товар успешно обновлён!"))
}

// Удаление товара: /delete ID
func (h *BotHandler) handleDelete(update tgbotapi.Update) {
	id, err := strconv.Atoi(update.Message.CommandArguments())
	if err != nil {
		h.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка: ID должно быть числом"))
		return
	}

	if err := h.repo.Delete(uint(id)); err != nil {
		h.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка при удалении товара"))
		return
	}

	h.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Товар успешно удалён!"))
}

// Вспомогательная функция для разбиения аргументов
func splitArgs(args string, count int) []string {
	parts := make([]string, 0)
	current := ""
	quote := false

	for _, ch := range args {
		if ch == ' ' && !quote {
			parts = append(parts, current)
			current = ""
			continue
		}
		if ch == '"' {
			quote = !quote
			continue
		}
		current += string(ch)
	}

	if len(current) > 0 {
		parts = append(parts, current)
	}

	if len(parts) != count {
		return nil
	}
	return parts
}
