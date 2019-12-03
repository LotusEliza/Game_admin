package telegram

import (
	"github.com/finalist736/gokit/database"
	"github.com/finalist736/gokit/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"tcs/items"
	"tcs/machanics/duty"
	"tcs/models"
	"tcs/telegram/menu"
	"tcs/tools"
	"time"
)

func dutyCommon(avatar *models.Avatar, d int) {
	if avatar.IsDuty() {
		msg := tgbotapi.NewMessage(avatar.Player.Chat, alreadyOnDuty)
		msg.ReplyMarkup = menu.StateOnlyMenu
		Send(avatar.ChatID(), msg)
		return
	}
	avatar.Duty = models.CreateDuty(avatar.Player.Tg, d)
	var (
		txt   string
		after time.Duration
	)

	switch d {
	case duty.DUTY_CANTEEN:
		txt = "Отправляюсь в наряд, вернусь через час"
		after = duty.DUTY_CANTEEN_PERIOD
	case duty.DUTY_GEO:
		txt = "Отправляюсь в наряд, вернусь через 3 часа"
		after = duty.DUTY_GEO_PERIOD
	case duty.DUTY_PATROL:
		txt = "Отправляюсь в наряд, вернусь через 7 часов"
		after = duty.DUTY_PATROL_PERIOD
	}
	msg := tgbotapi.NewMessage(avatar.Player.Chat, txt)
	msg.ReplyMarkup = menu.StateOnlyMenu
	Send(avatar.ChatID(), msg)
	go dutyRoutine(avatar, after, d)
}

func hqDutyCanteen(avatar *models.Avatar) {
	dutyCommon(avatar, 1)
}

func hqDutyGeo(avatar *models.Avatar) {
	dutyCommon(avatar, 2)
}

func hqDutyPatrol(avatar *models.Avatar) {
	dutyCommon(avatar, 3)
}

func hqDuty(avatar *models.Avatar) {
	msg := tgbotapi.NewMessage(avatar.Player.Chat,
		`♟️ Хорошо, что ты пришел. Людей всегда не хватает.
Выбери в какой наряд хочешь заступить:

🍜 Столовая: 1 час; 💰10
🛠 Рейд геологов: 3 часа; 💰20; немного ресурсов
🚶‍♂️ Патруль: 7 часов; 💰100 + все что найдешь на локации
`)
	msg.ReplyMarkup = menu.HeadQuarterDutyMenu
	Send(avatar.ChatID(), msg)
}

func dutyRoutine(avatar *models.Avatar, after time.Duration, d int) {
	logger.StdOut().Debugf("starting duty routine for %d, after %s", avatar.Player.Tg, after.String())
	<-time.After(after)
	logger.StdOut().Debugf("done duty routine for %d, after %s", avatar.Player.Tg, after.String())
	// reward after duty
	var txt string
	avatar.Lock()
	switch d {
	case duty.DUTY_CANTEEN:
		avatar.Wallet.Credits += 10
		txt = "💰10"
	case duty.DUTY_GEO:
		avatar.Wallet.Credits += 20
		for i := 0; i < 20; i++ {
			avatar.InventoryAddItem(8)
			avatar.InventoryAddItem(10)
			avatar.InventoryAddItem(15)
		}
		var item = items.GetByID(8)
		txt = "💰20\n" + item.InventoryItemTitle() + " (20)\n"
		item = items.GetByID(10)
		txt += item.InventoryItemTitle() + " (20)\n"
		item = items.GetByID(15)
		txt += item.InventoryItemTitle() + " (20)\n"
	case duty.DUTY_PATROL:
		avatar.Wallet.Credits += 100
		var id = tools.RandMinMax(1, 6)
		avatar.InventoryAddItem(id)
		var item = items.GetByID(id)
		txt = "💰100\n" + item.InventoryItemTitle() + "\n"
	}
	avatar.Wallet.UpdateMoney()
	avatar.Unlock()
	msg := tgbotapi.NewMessage(avatar.Player.Chat, "Смена в наряде закончилась, твоя награда:\n\n"+txt)
	msg.ReplyMarkup = menu.BaseMenu
	Send(avatar.ChatID(), msg)
}

func RestartDutyRoutines() {
	session := database.GetDefaultSession()
	var (
		tgs    []*models.PlayerDuty
		avatar *models.Avatar
		ok     bool
		after  time.Duration
		now    time.Time
	)
	_, err := session.Select("tg, timeend, duty").From(models.PlayerDutyTableName).Where("timeend>NOW()").Load(&tgs)
	if err != nil {
		logger.StdErr().Errorf("RestartDutyRoutines Select player_duty error: %s", err)
		return
	}
	now = time.Now()
	for _, it := range tgs {
		avatar, ok = models.LoadAvatar(it.Tg)
		if !ok {
			continue
		}
		after = it.TimeEnd.Sub(now)
		cache.Set(avatarCacheName(avatar.Player.Tg), avatar)
		go dutyRoutine(avatar, after, it.Duty)
	}
}
