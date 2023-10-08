package main

import (
	"math/rand"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const TOKEN = "6678079673:AAHV-zFEfjF2czip5wpCIF7u2Y9WBLh5uZU"

var bot *tgbotapi.BotAPI
var chatID int64

var TelegramBotNames = [3]string{"стивен", "стив", "кинг"}

var answers = []string{
	"Да",
	"Нет",
	"Иногда самые страшные монстры находятся внутри нас.",
	"Смело иди к своим страхам, и они могут исчезнуть.",
	"В жизни всегда есть место для надежды, даже в самых темных моментах.",
	"Не забывайте, что добро и зло могут скрываться внутри одного и того же человека.",
	"Дружба - это сила, способная преодолеть даже самые страшные испытания.",
	"Не играйте с нечистой силой, она всегда имеет свою цену.",
	"Мы сами создаем многие свои кошмары.",
	"Смех - это лучшее средство от страха.",
	"Постоянное стремление к власти может стать твоим конечным поражением.",
	"Иногда надо пойти на компромисс, чтобы выжить.",
	"Семья - это не всегда те, кто кровно связан с вами.",
	"Правда часто бывает страшнее вымысла.",
	"Зло не всегда является очевидным.",
	"Верьте в свои силы, и они могут помочь вам справиться с любыми трудностями.",
	"Люди способны на ужасающие поступки, но они также могут проявлять героизм и сострадание.",
	"Нельзя бежать от своего прошлого, но можно извлечь уроки из него.",
	"Судьба может быть непредсказуемой, и важно готовиться к любым обстоятельствам.",
	"Счастье часто находится в мелочах.",
	"Зло не победит, если у вас есть верные друзья.",
	"Иногда ты можешь найти спасение только внутри себя.",
	"Страх - это всего лишь обманчивая иллюзия.",
	"Лучший способ понять человека - это посмотреть, как он поступает в экстремальных ситуациях.",
	"Самое важное - сохранить свою человечность, даже в самых ужасных условиях.",
	"Иногда нужно отказаться от прошлого, чтобы двигаться вперед.",
	"Не забывайте, что внутри каждого из нас таится темная сторона.",
	"Судьба не всегда подчиняется нашим планам.",
	"Важно находить радость в моментах спокойствия.",
	"Смерть - это не конец, а новое начало.",
	"Сила дружбы может перемещать горы.",
	"Жизнь - это непредсказуемое приключение, и важно жить ее на полную катушку.",
}

func connectWithTelegram() {
	var err error
	bot, err = tgbotapi.NewBotAPI(TOKEN)
	if err != nil {
		panic("Cannot connect to telegram")
	}
}

func sendMessage(msg string) {
	msgConfig := tgbotapi.NewMessage(chatID, msg)
	_, err := bot.Send(msgConfig)
	if err != nil {
		panic(err)
	}
}

func IsMessageForTelegramBot(update *tgbotapi.Update) bool {
	if update.Message == nil || update.Message.Text == "" {
		return false
	}

	msgInLowerCase := strings.ToLower(update.Message.Text)
	for _, name := range TelegramBotNames {
		if strings.Contains(msgInLowerCase, name) {
			return true
		}
	}

	return false
}

func getBotAnswer() string {
	index := rand.Intn(len(answers))
	return answers[index]
}

func sendAnswer(update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(chatID, getBotAnswer())
	msg.ReplyToMessageID = update.Message.MessageID
	_, err := bot.Send(msg)
	if err != nil {
		panic(err)
	}
}

func main() {
	connectWithTelegram()

	updateConfig := tgbotapi.NewUpdate(0)
	for update := range bot.GetUpdatesChan(updateConfig) {
		if update.Message != nil && update.Message.Text == "/start" {
			chatID = update.Message.Chat.ID
			sendMessage("Задай свой вопрос, назвав меня по имени. Ответом на вопрос должны быть \"да\" либо \"нет\" ")
		}

		if IsMessageForTelegramBot(&update) {
			sendAnswer(&update)
		}
	}
}
