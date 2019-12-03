package telegram

import (
	"fmt"
	"github.com/finalist736/gokit/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"math/rand"
	"tcs/items"
	"tcs/machanics/miner"
	"tcs/models"
	"tcs/telegram/menu"
	"time"
)

func mineGas(avatar *models.Avatar) {
	if avatar.IsMining {
		msg := tgbotapi.NewMessage(avatar.Player.Chat, cantAction)
		msg.ReplyMarkup = menu.StateOnlyMenu
		Send(avatar.ChatID(), msg)
		return
	}
	if avatar.IsMoving() {
		msg := tgbotapi.NewMessage(avatar.Player.Chat, cantAction)
		msg.ReplyMarkup = menu.MoveMenu
		Send(avatar.ChatID(), msg)
		return
	}
	if avatar.Player.Air == 0 {
		msg := tgbotapi.NewMessage(avatar.Player.Chat, airLessMessage)
		baseLocation(avatar)
		msg.ReplyMarkup = menu.BaseMenu
		Send(avatar.ChatID(), msg)
		return
	}
	if avatar.Player.Location.IsBase() {
		msg := tgbotapi.NewMessage(avatar.ChatID(), "🛠 Нельзя майнить на базе! Дневальный расстроится...")
		msg.ReplyMarkup = menu.BaseMenu
		Send(avatar.ChatID(), msg)
		return
	}
	if !avatar.HasDevices() {
		msg := tgbotapi.NewMessage(avatar.ChatID(), "🛠 Нечем майнить! Купи майнилку...")
		msg.ReplyMarkup = menu.OutSideMenu
		Send(avatar.ChatID(), msg)
		return
	}
	if avatar.UsingDevice == nil {
		msg := tgbotapi.NewMessage(avatar.ChatID(), "🛠 Сначала нужно выбрать устройство майнинга...")
		msg.ReplyMarkup = menu.OutSideMenu
		Send(avatar.ChatID(), msg)
		return
	}
	go mineRoutine(avatar, items.ITEM_GAS)
	msg := tgbotapi.NewMessage(avatar.ChatID(), "🛢 Ждем майн газа")
	msg.ReplyMarkup = menu.StateOnlyMenu
	Send(avatar.ChatID(), msg)
}

func mineResources(avatar *models.Avatar) {
	if avatar.IsMining {
		msg := tgbotapi.NewMessage(avatar.Player.Chat, cantAction)
		msg.ReplyMarkup = menu.StateOnlyMenu
		Send(avatar.ChatID(), msg)
		return
	}
	if avatar.IsMoving() {
		msg := tgbotapi.NewMessage(avatar.Player.Chat, cantAction)
		msg.ReplyMarkup = menu.MoveMenu
		Send(avatar.ChatID(), msg)
		return
	}
	if avatar.Player.Air == 0 {
		msg := tgbotapi.NewMessage(avatar.Player.Chat, airLessMessage)
		baseLocation(avatar)
		msg.ReplyMarkup = menu.BaseMenu
		Send(avatar.ChatID(), msg)
		return
	}
	if avatar.Player.Location.IsBase() {
		msg := tgbotapi.NewMessage(avatar.ChatID(), "🛠 Нельзя майнить на базе! Дневальный расстроится...")
		msg.ReplyMarkup = menu.BaseMenu
		Send(avatar.ChatID(), msg)
		return
	}
	if !avatar.HasDevices() {
		msg := tgbotapi.NewMessage(avatar.ChatID(), "🛠 Нечем майнить! Купи майнилку...")
		msg.ReplyMarkup = menu.OutSideMenu
		Send(avatar.ChatID(), msg)
		return
	}
	if avatar.UsingDevice == nil {
		msg := tgbotapi.NewMessage(avatar.ChatID(), "🛠 Сначала нужно выбрать устройство майнинга...")
		msg.ReplyMarkup = menu.OutSideMenu
		Send(avatar.ChatID(), msg)
		return
	}
	go mineRoutine(avatar, items.ITEM_METAL)
	msg := tgbotapi.NewMessage(avatar.ChatID(), "🧲 Ждем майн ресурсов")
	msg.ReplyMarkup = menu.StateOnlyMenu
	Send(avatar.ChatID(), msg)
	return
}

func mineRandom(avatar *models.Avatar) {
	if avatar.IsMining {
		msg := tgbotapi.NewMessage(avatar.Player.Chat, cantAction)
		msg.ReplyMarkup = menu.StateOnlyMenu
		Send(avatar.ChatID(), msg)
		return
	}
	if avatar.IsMoving() {
		msg := tgbotapi.NewMessage(avatar.Player.Chat, cantAction)
		msg.ReplyMarkup = menu.MoveMenu
		Send(avatar.ChatID(), msg)
		return
	}
	if avatar.Player.Air == 0 {
		msg := tgbotapi.NewMessage(avatar.Player.Chat, airLessMessage)
		baseLocation(avatar)
		msg.ReplyMarkup = menu.BaseMenu
		Send(avatar.ChatID(), msg)
		return
	}
	if avatar.Player.Location.IsBase() {
		msg := tgbotapi.NewMessage(avatar.ChatID(), "🛠 Нельзя майнить на базе! Дневальный расстроится...")
		msg.ReplyMarkup = menu.BaseMenu
		Send(avatar.ChatID(), msg)
		return
	}
	if !avatar.HasDevices() {
		msg := tgbotapi.NewMessage(avatar.ChatID(), "🛠 Нечем майнить! Купи майнилку...")
		msg.ReplyMarkup = menu.OutSideMenu
		Send(avatar.ChatID(), msg)
		return
	}
	if avatar.UsingDevice == nil {
		msg := tgbotapi.NewMessage(avatar.ChatID(), "🛠 Сначала нужно выбрать устройство майнинга...")
		msg.ReplyMarkup = menu.OutSideMenu
		Send(avatar.ChatID(), msg)
		return
	}
	go mineRoutine(avatar, items.ITEM_JUNK)
	msg := tgbotapi.NewMessage(avatar.ChatID(), "🛢🧲 Что-то должно намайнить")
	msg.ReplyMarkup = menu.StateOnlyMenu
	Send(avatar.ChatID(), msg)
	return
}

func mineRoutine(avatar *models.Avatar, mineType items.ItemType) {
	avatar.Lock()
	if avatar.UsingDevice == nil {
		logger.StdOut().Debug("no mine device choiced")
		avatar.Unlock()
		return
	}
	avatar.IsMining = true
	avatar.Player.Air--
	avatar.Player.UpdateAir()
	mineDevice := avatar.UsingDevice
	avatar.Unlock()
	logger.StdOut().Debugf("start mine: %s", mineType)
	<-time.After(time.Second * time.Duration(mineDevice.Time))
	logger.StdOut().Debugf("end mine: %s", mineType)
	avatar.Lock()
	defer avatar.Unlock()

	var (
		item   *items.Item
		itemID int
		//junkRandom int
		mineCount = mineDevice.Mine
	)

	avatar.IsMining = false
	txt := fmt.Sprintf("%s\nРезультаты майна:\n", avatar.LocationMessage())
	//// TODO mine resources by location!!!
	mineCount, itemID = miner.Get(
		avatar.Player.Location,
		mineType,
		mineDevice.Mine,
		avatar.Player.PosX,
		avatar.Player.PosY)
	for i := 0; i < mineCount; i++ {
		item = items.GetByID(itemID)
		txt += item.InventoryItemTitle() + "\n"
		avatar.InventoryAddItem(itemID)
	}

	cnt := rand.Intn(3-1) + 1
	var junks []items.Item
	switch avatar.Player.Location {
	case 3, 4:
		junks = items.GetJunkFor1Location()
	default:
		junks = items.GetByType(items.ITEM_JUNK)
	}
	for i := 0; i < cnt; i++ {
		itemID = junks[rand.Intn(len(junks))].ID
		item = items.GetByID(itemID)
		txt += item.InventoryItemTitle() + "\n"
		avatar.InventoryAddItem(itemID)
	}
	avatar.UsingDevice = nil
	msg := tgbotapi.NewMessage(avatar.ChatID(), txt+"\n")
	msg.ReplyMarkup = menu.OutSideMenu
	Send(avatar.ChatID(), msg)
	mob, exists := models.GetCurrentMob(avatar.Player.Location)
	if exists {
		logger.StdOut().Debugf("mob: %+v", mob)
		mobBattle(avatar, mob)
	}
}
