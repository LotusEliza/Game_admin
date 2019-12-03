package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"math/rand"
	"tcs/machanics/miner"
	"tcs/models"
	"time"
)

var canteenIdleMessages = []string{
	//"- Слышал, Льва Абалкина убили.\n- А кто это?\n- Ну ты че? Это специалист по голованам!\n- Да точно, вспомнил!!",
	"- Погода отличная на планете! Хоть на море едь...",
	"- Тахорги - это здоровенные чудовища с пандоры! Говорят его видели на этой планете...\n- Он сюда сам пешком пришел?",
	"- Я лично не видел, но знаю парня, который точно видел катализатор, который из воды выделяет кислород и на ходу можно баллоны пополнить!\n- А он рассказывал как этот катализатор выглядит?",
	"- Тагоряне отозвали дипмиссию с Земли 9 лет назад!\n- Быстро время летит, как вчера было...",
	"- Ночью лучше не ходи в пустыню, там живность очень злая!\n- Я уже ходил! Встретил аномалию, оказалась сильнее чем днем, но добыча вдохновляет...",
	"- Стояли звери около двери, в них стреляли, они умирали...",
	"- В столовой хорошо! Поел - ❤️ ХП восстановил!",
	//"- На Саракше сейчас !",
}

func canteen(avatar *models.Avatar) {
	avatar.Player.HP = 10
	avatar.Player.UpdateHP()
	text := avatar.HealthMessage()
	text += "\n🍜 Столовая\nВо время обеда ты подслушал разговор за соседним столом:\n"
	if time.Now().Sub(avatar.CanteenMessage) < time.Minute {
		text += canteenIdleMessages[rand.Intn(len(canteenIdleMessages))]
	} else {
		text += miner.GetAllSpots(avatar.Player.Location)
		avatar.CanteenMessage = time.Now()
	}
	msg := tgbotapi.NewMessage(avatar.ChatID(), text)
	//msg.ReplyMarkup = menu.CanteenMenu
	//avatar.MenuHistory.Push(&menu.BaseMenu)
	Send(avatar.ChatID(), msg)
}
