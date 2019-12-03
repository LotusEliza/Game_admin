package telegram

import (
	"fmt"
	"github.com/finalist736/gokit/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"tcs/items"
	"tcs/machanics/norms"
	"tcs/models"
	"tcs/telegram/menu"
	"tcs/tools"
	"time"
)

func hq(avatar *models.Avatar) {
	msg := tgbotapi.NewMessage(avatar.Player.Chat, avatar.HqMessage())
	msg.ReplyMarkup = menu.HeadQuarterMenu
	Send(avatar.ChatID(), msg)
}

func hqPutNorm(avatar *models.Avatar) {
	norm := models.PlayerNormLoadCurrent(avatar.Player.Tg)
	if norm != nil {
		msg := tgbotapi.NewMessage(avatar.Player.Chat, "🏆 Ты уже сдал текущие нормы!")
		msg.ReplyMarkup = menu.BaseMenu
		Send(avatar.ChatID(), msg)
		return
	}
	needNorm := norms.GetNormByID(tools.WeekDayNorm())
	var resourcesExists = true
	var cnt = 0
	for _, it := range needNorm.Norms {
		cnt = avatar.InventoryGetResourcesCount(it.Resource)
		if cnt < it.Count {
			resourcesExists = false
			break
		}
	}
	if !resourcesExists {
		msg := tgbotapi.NewMessage(avatar.Player.Chat, "Недостаточно ресурсов!")
		msg.ReplyMarkup = menu.BaseMenu
		Send(avatar.ChatID(), msg)
		return
	}
	curDate := time.Now().Format("2006-01-02")

	for _, it := range needNorm.Norms {
		pn := models.PlayerNorms{
			Tg:       avatar.Player.Tg,
			Date:     curDate,
			Resource: it.Resource,
			Amount:   it.Count,
		}
		err := models.PlayerNormCreate(&pn)
		if err != nil {
			logger.StdErr().Error("PlayerNormCreate error: %s", err)
		}
		avatar.InventoryRemoveResourceCount(it.Resource, it.Count)
	}
	txt := fmt.Sprintf("🌠 Отлично! \nТекущий норматив сдан, твоя награда:\n💰: %d\n🏆: %d",
		needNorm.RewardCredits,
		needNorm.RewardReputation)
	avatar.Wallet.Credits += needNorm.RewardCredits
	avatar.Wallet.Reputation += needNorm.RewardReputation
	avatar.Wallet.UpdateMoney()
	msg := tgbotapi.NewMessage(avatar.Player.Chat, txt)
	msg.ReplyMarkup = menu.BaseMenu
	Send(avatar.ChatID(), msg)
}

func hqNorm(avatar *models.Avatar) {
	norm := norms.GetNormByID(tools.WeekDayNorm())
	txt := avatar.HqNormMessage()
	txt += "\nТекущий норматив:\n"
	for _, it := range norm.Norms {
		res := items.GetByID(it.Resource)
		txt += res.InventoryItemTitle() + fmt.Sprintf(" = %d\n", it.Count)
	}
	txt += "\n"
	msg := tgbotapi.NewMessage(avatar.Player.Chat, txt)
	msg.ReplyMarkup = menu.HeadQuarterNormMenu
	Send(avatar.ChatID(), msg)
}
