package telegram

import (
	"fmt"
	"github.com/finalist736/gokit/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"tcs/items"
	"tcs/machanics/miner"
	"tcs/models"
	"tcs/telegram/menu"
	"time"
)

func outsideDevices(avatar *models.Avatar) {
	if avatar.Player.Location.IsBase() {
		return
	}
	if avatar.HasDevices() {
		msg := tgbotapi.NewMessage(avatar.Player.Chat, avatar.DevicesMessage())
		prevMenu := avatar.OutsideDevicesMenu()
		msg.ReplyMarkup = prevMenu
		Send(avatar.ChatID(), msg)
	} else {
		msg := tgbotapi.NewMessage(avatar.Player.Chat, "У вас нет устройств!")
		msg.ReplyMarkup = menu.OutSideMenu
		Send(avatar.ChatID(), msg)
	}
}

func tryOutsideDevices(avatar *models.Avatar, txt string) (tgbotapi.Chattable, bool) {
	if avatar.Player.Location.IsBase() {
		return nil, false
	}
	if !avatar.HasDevices() {
		return nil, false
	}
	devs := avatar.GetInventoryItemsByType(items.ITEM_DEVICE)
	for _, dev := range devs {
		if dev.InventoryItemTitle() != txt {
			continue
		}

		switch dev.SubType {
		case items.ITEM_SUB_BOER:
			avatar.UsingDevice = dev
			msg := tgbotapi.NewMessage(avatar.ChatID(), "🛢🧲 Что майнить будем?")
			msg.ReplyMarkup = menu.ChoiceResourceMenu
			return msg, true
		case items.ITEM_SUB_NULL_T:
			if time.Now().Sub(avatar.LastNullT) < time.Minute*time.Duration(dev.Time) {
				msg := tgbotapi.NewMessage(
					avatar.ChatID(),
					fmt.Sprintf("🌀 Еще нельзя использовать %s!\nЖди еще %d минут",
						dev.Title,
						int((time.Minute*time.Duration(dev.Time)-time.Now().Sub(avatar.LastNullT)).Minutes())))
				msg.ReplyMarkup = menu.OutSideMenu
				return msg, true
			}
			msg := tgbotapi.NewMessage(
				avatar.ChatID(),
				fmt.Sprintf("🌀 Произвести NULL-Транспортировку на базу?"))
			msg.ReplyMarkup = menu.NULLtMenu
			return msg, true
		case items.ITEM_SUB_SCANNER:
			if avatar.IsMining {
				msg := tgbotapi.NewMessage(avatar.Player.Chat, cantAction)
				msg.ReplyMarkup = menu.StateOnlyMenu
				return msg, true
			}
			if avatar.IsScanning {
				msg := tgbotapi.NewMessage(avatar.Player.Chat, cantAction)
				msg.ReplyMarkup = menu.StateOnlyMenu
				return msg, true
			}
			if avatar.IsMoving() {
				msg := tgbotapi.NewMessage(avatar.Player.Chat, cantAction)
				msg.ReplyMarkup = menu.MoveMenu
				return msg, true
			}
			if avatar.Player.Air == 0 {
				msg := tgbotapi.NewMessage(avatar.Player.Chat, airLessMessage)
				baseLocation(avatar)
				msg.ReplyMarkup = menu.BaseMenu
				return msg, true
			}
			if avatar.Player.Location.IsBase() {
				msg := tgbotapi.NewMessage(avatar.ChatID(), "🛠 Нельзя сканировать на базе! Дневальный расстроится...")
				msg.ReplyMarkup = menu.BaseMenu
				return msg, true
			}
			if !avatar.HasDevices() {
				msg := tgbotapi.NewMessage(avatar.ChatID(), "🛠 Нечем сканировать! Сначала купи коробочку...")
				msg.ReplyMarkup = menu.OutSideMenu
				return msg, true
			}
			msg := tgbotapi.NewMessage(
				avatar.ChatID(),
				fmt.Sprintf("📡 Поиск ресурсов..."))
			msg.ReplyMarkup = menu.StateOnlyMenu
			avatar.UsingDevice = dev
			go scanningRoutine(avatar)
			return msg, true
		}

		//Send(avatar.ChatID(), msg)
		//go mineRoutine(avatar, dev)
	}
	return nil, false
}

func nullTransporting(avatar *models.Avatar) {
	if avatar.Player.Location.IsBase() {
		return
	}
	if !avatar.HasDevices() {
		return
	}
	msg := tgbotapi.NewMessage(
		avatar.ChatID(),
		fmt.Sprintf("🌀 Начинаю NULL-Транспортировку на базу"))
	msg.ReplyMarkup = menu.StateOnlyMenu
	Send(avatar.ChatID(), msg)
	switch avatar.Player.Faction {
	case 1:
		avatar.Player.Location = models.CorporationBase
		avatar.Player.PosY = 1
		avatar.Player.PosX = 1
	case 2:
		avatar.Player.Location = models.PiratesBase
		avatar.Player.PosY = 25
		avatar.Player.PosX = 25
	}
	avatar.LastNullT = time.Now()
	avatar.Player.UpdateLocation()
	avatar.SetMoving()
	go movingRoutine(avatar)
}

func scanningRoutine(avatar *models.Avatar) {

	avatar.Lock()
	if avatar.UsingDevice == nil {
		logger.StdOut().Debug("no mine device choiced")
		avatar.Unlock()
		return
	}
	avatar.IsScanning = true
	avatar.Player.Air--
	avatar.Player.UpdateAir()
	avatar.Unlock()
	logger.StdOut().Debugf("start scanning")
	<-time.After(time.Second * time.Duration(avatar.UsingDevice.Time))
	logger.StdOut().Debugf("end scanning")
	avatar.Lock()
	defer avatar.Unlock()

	ok, nearestSpot := miner.GetNearestSpot(avatar.Player.Location, avatar.Player.PosX, avatar.Player.PosY)
	if !ok {
		msg := tgbotapi.NewMessage(
			avatar.ChatID(),
			fmt.Sprintf("Ничего не найдено"))
		msg.ReplyMarkup = menu.OutSideMenu
		Send(avatar.ChatID(), msg)
	} else {
		txt := fmt.Sprintf(
			"Найдено месторождение %sа в координатах [%d:%d]",
			nearestSpot.Type.Name(),
			nearestSpot.Coordinates[0],
			nearestSpot.Coordinates[1])
		msg := tgbotapi.NewMessage(
			avatar.ChatID(),
			fmt.Sprintf(txt))
		msg.ReplyMarkup = menu.OutSideMenu
		Send(avatar.ChatID(), msg)
	}
	avatar.IsScanning = false
	avatar.UsingDevice = nil
}
