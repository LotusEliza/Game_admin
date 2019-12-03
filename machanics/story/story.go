package story

import (
	"github.com/finalist736/gokit/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"tcs/models"
)

type (
	Faction int
	Answer  struct {
		Answer       string
		NextQuestion int
	}
	Questions struct {
		ID       int
		Question string
		Answers  []Answer
		Faction  Faction
		Location int
	}
)

const (
	None        = 0
	Corporation = 1
	Pirates     = 2
)

var (
	storyQuestions = []Questions{
		{
			ID:       1,
			Question: `🚀 Вы летите на корабле корпорации на планету Проксима b ☄️. Вы, геолог по образованию, устроились работать на тамошний горно-обогатительный комбинат. Вот корабль вышел из гипер-скорости, как тут бац💥 его захватили пираты!🏴‍☠️`,
			Answers: []Answer{
				{
					NextQuestion: 2,
					Answer:       "Хорошая история 👍",
				},
				{
					NextQuestion: 2,
					Answer:       "Дальше ⏭️",
				},
			},
			Faction:  None,
			Location: 0,
		},
		{
			ID:       2,
			Question: "Один пассажир предложил бежать к спасательным капсулам 🚎, но пираты по громкой связи 📢 передали информацию: Уважаемые пассажиры 👨‍👩‍👧‍👦, корпорация не доплачивает здесь зарплату 💵. А у нас есть печеньки 🍔, переходите на нашу сторону! Перед вами стоит выбор: 🤷‍♂️ свалить на спасательной капсуле и работать на заводе 🏭, или перейти на темную сторону 🧟‍♂️?",
			Answers: []Answer{
				{
					Answer:       "Скорее в капсулу! 🚣‍♂️",
					NextQuestion: 3,
				},
				{
					Answer:       "Буду бандитом! 🏴‍☠️",
					NextQuestion: 4,
				},
			},
			Faction:  None,
			Location: 0,
		},
		{
			ID:       3,
			Question: "Вам удалось отстоять очередь 🧘‍♂️ на спасательную капсулу, и теперь вы доблестный сотрудник галактической корпорации, поздравляем!",
			Answers:  nil,
			Faction:  Corporation,
			Location: 1,
		},
		{
			ID:       4,
			Question: "По длинному коридору Вы пошли к капитанскому мостику 👨‍🚀. Вас очень радушно встретили пираты и угостили чаем ☕ с овсяными пряниками! Теперь Вы на темной стороне 🦹‍♂️, поздравляем!",
			Answers:  nil,
			Faction:  Pirates,
			Location: 2,
		},
	}
)

func GetStoryByID(id int) *Questions {
	for _, it := range storyQuestions {
		if it.ID == id {
			return &it
		}
	}
	return nil
}

func GetButtonsFromAnswers(answers []Answer) tgbotapi.ReplyKeyboardMarkup {
	var keyboard [][]tgbotapi.KeyboardButton
	var keys []tgbotapi.KeyboardButton
	for i, it := range answers {
		if i+1%2 == 0 {
			keyboard = append(keyboard, keys)
			keys = make([]tgbotapi.KeyboardButton, 0)
		}
		keys = append(keys, tgbotapi.NewKeyboardButton(it.Answer))
	}
	if len(keys) > 0 {
		keyboard = append(keyboard, keys)
	}
	//logger.StdOut().Debugf("keyboard: %+v", keyboard)
	return tgbotapi.NewReplyKeyboard(keyboard...)
}

func GetCurrentStory(player *models.Player, chatID int64) tgbotapi.Chattable {
	currentStory := GetStoryByID(player.Story)
	if currentStory == nil {
		logger.StdErr().Errorf("[GetCurrentStory] no such story question: %d", player.Story)
		return nil
	}

	msg := tgbotapi.NewMessage(chatID, currentStory.Question)
	if len(currentStory.Answers) == 0 {
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	} else {
		msg.ReplyMarkup = GetButtonsFromAnswers(currentStory.Answers)
	}

	return msg
}

func CheckCorrectAnswer(msg string, player *models.Player) int {
	currentStory := GetStoryByID(player.Story)
	if currentStory == nil {
		logger.StdErr().Errorf("[CheckCorrectAnswer] no such story question: %d", player.Story)
		return -1
	}
	for _, answer := range currentStory.Answers {
		if answer.Answer == msg {
			return answer.NextQuestion
		}
	}
	return -1
}

func CheckEndOfStory(player *models.Player) bool {
	currentStory := GetStoryByID(player.Story)
	if currentStory == nil {
		logger.StdErr().Errorf("[CheckEndOfStory] no such story question: %d", player.Story)
		return false
	}
	if currentStory.Faction == None {
		return false
	}
	player.Faction = int(currentStory.Faction)
	player.UpdateFaction()
	player.Location = models.Location(currentStory.Location)
	player.PosX = models.BasesExits[player.Location].PosX
	player.PosY = models.BasesExits[player.Location].PosY
	player.UpdateLocation()
	return true
}
