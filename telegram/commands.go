package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
	"strings"
	"tcs/items"
	"tcs/models"
	"tcs/telegram/menu"
)

func commands(avatar *models.Avatar, u *tgbotapi.Update) {
	cmd := u.Message.Command()
	switch {
	case cmd == "start":
		//player := &models.Player{}
	case cmd == "reset":
		avatar.Reset()
		cache.Delete(avatarCacheName(avatar.Player.Tg))
		baseMsg := tgbotapi.NewMessage(avatar.Player.Chat, "Вас сбросили с обрыва! пишите /start для перезапуска!")
		baseMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		Send(avatar.ChatID(), baseMsg)
	case cmd == "home":
		baseLocation(avatar)
		msg := tgbotapi.NewMessage(avatar.ChatID(), avatar.StateMessage())
		msg.ReplyMarkup = menu.BaseMenu
		Send(avatar.ChatID(), msg)
	case cmd == "addhp":
		avatar.Player.HP += 10
		avatar.Player.UpdateHP()
		msg := tgbotapi.NewMessage(avatar.ChatID(), avatar.LocationMessage())
		Send(avatar.ChatID(), msg)
	case cmd == "money":
		avatar.Wallet.Credits += 100
		avatar.Wallet.UpdateMoney()
		msg := tgbotapi.NewMessage(avatar.ChatID(), avatar.MoneyMessage())
		Send(avatar.ChatID(), msg)
	case cmd == "addair":
		avatar.Player.Air += 10
		avatar.Player.UpdateAir()
		msg := tgbotapi.NewMessage(avatar.ChatID(), avatar.LocationMessage())
		Send(avatar.ChatID(), msg)
	case cmd == "menu":
		//avatar.MenuHistory = stack.New()
		msg := tgbotapi.NewMessage(avatar.ChatID(), avatar.StateMessage())
		if avatar.Player.Location.IsBase() {
			msg.ReplyMarkup = menu.BaseMenu
		} else {
			msg.ReplyMarkup = menu.OutSideMenu
		}
		Send(avatar.ChatID(), msg)
	case cmd == "loc":
		locs := u.Message.CommandArguments()
		parts := strings.Split(locs, " ")
		if len(parts) == 2 {
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])
			if x == 0 {
				x = 1
			}
			if y == 0 {
				y = 1
			}
			avatar.Player.PosX = x
			avatar.Player.PosY = y
			avatar.Player.UpdateLocation()
			Send(avatar.ChatID(), tgbotapi.NewMessage(avatar.ChatID(), avatar.LocationMessage()))
		}
	// buy
	case strings.HasPrefix(cmd, "buy_"):
		parts := strings.Split(cmd, "_")
		if len(parts) != 2 {
			Send(avatar.ChatID(), tgbotapi.NewMessage(avatar.ChatID(), "Неверная команда!"))
			return
		}
		itemID, err := strconv.Atoi(parts[1])
		if err != nil {
			Send(avatar.ChatID(), tgbotapi.NewMessage(avatar.ChatID(), "Неверная команда!"))
			return
		}
		if !avatar.CanBuy(itemID) {
			Send(avatar.ChatID(), tgbotapi.NewMessage(avatar.ChatID(), "Недостаточно денег"))
			return
		}
		item, bought := avatar.BuyItem(itemID)
		if !bought {
			Send(avatar.ChatID(), tgbotapi.NewMessage(avatar.ChatID(), "Какой-то глюк при покупке предмета"))
			return
		}
		Send(
			avatar.ChatID(),
			tgbotapi.NewMessage(
				avatar.ChatID(),
				fmt.Sprintf("%s\n\n%s куплен и добавлен в инвентарь",
					avatar.MoneyMessage(), item.InventoryItemTitle()),
			))
		message := ""
		switch item.Type {
		case items.ITEM_ARMOR:
			message = avatar.ShopArmorMessage()
		case items.ITEM_WEAPON:
			message = avatar.ShopWeaponMessage()
		case items.ITEM_METAL:
			message = avatar.ShopResourcesMessage()
		case items.ITEM_GAS:
			message = avatar.ShopGasMessage()
		}
		if message == "" {
			return
		}
		Send(avatar.ChatID(), tgbotapi.NewMessage(avatar.ChatID(), message))
	// equip
	case strings.HasPrefix(cmd, "equip"):
		parts := strings.Split(cmd, "_")
		if len(parts) != 2 {
			Send(avatar.ChatID(), tgbotapi.NewMessage(avatar.ChatID(), "Неверная команда!"))
			return
		}
		itemID, err := strconv.Atoi(parts[1])
		if err != nil {
			Send(avatar.ChatID(), tgbotapi.NewMessage(avatar.ChatID(), "Неверная команда!"))
			return
		}
		item := items.GetByID(itemID)
		if item == nil {
			Send(avatar.ChatID(), tgbotapi.NewMessage(avatar.ChatID(), "Это нельзя носить с собой!"))
			return
		}
		invItem := avatar.InventoryGetItem(item.ID)
		if invItem == nil {
			Send(avatar.ChatID(), tgbotapi.NewMessage(avatar.ChatID(), "У вас нет такого предмета!"))
			return
		}
		if avatar.InventoryEquipItem(invItem) {
			if item.Type == items.ITEM_BALLOON {
				go airRoutine(avatar)
			}
			Send(avatar.ChatID(), tgbotapi.NewMessage(avatar.ChatID(), avatar.StateMessage()))
			Send(avatar.ChatID(), tgbotapi.NewMessage(avatar.ChatID(), avatar.InventoryMessage()))
			return
		}

	case strings.HasPrefix(cmd, "disass"):
		parts := strings.Split(cmd, "_")
		if len(parts) != 2 {
			Send(avatar.ChatID(), tgbotapi.NewMessage(avatar.ChatID(), "Неверная команда!"))
			return
		}
		itemID, err := strconv.Atoi(parts[1])
		if err != nil {
			Send(avatar.ChatID(), tgbotapi.NewMessage(avatar.ChatID(), "Неверная команда!"))
			return
		}
		item := items.GetByID(itemID)
		if item == nil {
			Send(avatar.ChatID(), tgbotapi.NewMessage(avatar.ChatID(), "Это нельзя разбирать!"))
			return
		}
		if avatar.InventoryGetItem(item.ID) == nil {
			Send(avatar.ChatID(), tgbotapi.NewMessage(avatar.ChatID(), "Нет такого предмета в инвентаре!"))
			return
		}
		factoryDisassembly(avatar, itemID)
	case strings.HasPrefix(cmd, "craft"):
		parts := strings.Split(cmd, "_")
		if len(parts) != 2 {
			Send(avatar.ChatID(), tgbotapi.NewMessage(avatar.ChatID(), "Неверная команда!"))
			return
		}

		craftID, err := strconv.Atoi(parts[1])
		if err != nil {
			Send(avatar.ChatID(), tgbotapi.NewMessage(avatar.ChatID(), "Неверная команда!"))
			return
		}

		factoryCraft(avatar, craftID)

	default:
		Send(avatar.ChatID(), tgbotapi.NewMessage(avatar.ChatID(), "I don't understand you!"))
	}
}
