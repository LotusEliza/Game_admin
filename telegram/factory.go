package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"math/rand"
	"tcs/items"
	"tcs/machanics/craft"
	"tcs/models"
	"tcs/telegram/menu"
	"time"
)

func factory(avatar *models.Avatar) {
	msg := tgbotapi.NewMessage(avatar.Player.Chat, "Мастерская!\n")
	msg.ReplyMarkup = menu.FactoryMenu
	Send(avatar.ChatID(), msg)
}

func factoryCraftList(avatar *models.Avatar) {
	if !avatar.Player.Location.IsBase() {
		msg := tgbotapi.NewMessage(avatar.Player.Chat, "Печатать устройста можно только на базе!")
		msg.ReplyMarkup = menu.FactoryMenu
		Send(avatar.ChatID(), msg)
		return
	}

	list, ok := craft.GetLocationCraftList(avatar.Player.Location)
	if !ok {
		msg := tgbotapi.NewMessage(avatar.Player.Chat, "Здесь нет 3D принтера!")
		msg.ReplyMarkup = menu.FactoryMenu
		Send(avatar.ChatID(), msg)
		return
	}
	var (
		recTxt string
	)
	recTxt = "🖨 Список рецептов:\n\n"
	for _, it := range list {
		recTxt += it.Title() + "\n\n"
	}
	msg := tgbotapi.NewMessage(avatar.Player.Chat, recTxt)
	msg.ReplyMarkup = menu.SimpleBackMenu
	Send(avatar.ChatID(), msg)
}

func factoryDisassemblyList(avatar *models.Avatar) {
	if !avatar.InventoryHasDevices() {
		msg := tgbotapi.NewMessage(avatar.Player.Chat, "В инвентаре нет устройств, чтоб разобрать")
		msg.ReplyMarkup = menu.FactoryMenu
		Send(avatar.ChatID(), msg)
		return
	}

	var showItemsType = []items.ItemType{items.ITEM_WEAPON, items.ITEM_ARMOR, items.ITEM_DEVICE, items.ITEM_BALLOON}
	typesText := ""
	for _, typ := range showItemsType {
		its := avatar.GetInventoryItemsByType(typ)
		if len(its) == 0 {
			continue
		}
		typesText += typ.Name() + ":\n"
		for _, it := range its {
			typesText += it.DisassemblyItemTitle() + "\n"
		}
		typesText += "\n"
	}
	txt := fmt.Sprintf(
		"%s\n\nЧто можно разобрать:\n%s",
		avatar.MoneyMessage(), typesText)
	msg := tgbotapi.NewMessage(avatar.Player.Chat, txt)
	msg.ReplyMarkup = menu.SimpleBackMenu
	Send(avatar.ChatID(), msg)
}

func factoryDisassembly(avatar *models.Avatar, itemID int) {
	var (
		i          int
		rnd        int
		metal, gas int
		item       *items.Item
		itemMetal  *items.Item
		itemGas    *items.Item
	)
	item = items.GetByID(itemID)
	for i = 0; i < item.Sell; i++ {
		rnd = rand.Intn(2)
		if rnd == 0 {
			// metal
			avatar.InventoryAddItem(8)
			metal++
		} else {
			// gas
			avatar.InventoryAddItem(10)
			gas++
		}
	}

	avatar.InventoryRemoveItem(itemID)

	itemMetal = items.GetByID(8)
	itemGas = items.GetByID(10)

	txt := fmt.Sprintf("Разобрав %s\n получили:\n%s (%d)\n%s (%d)\n",
		item.InventoryItemTitle(),
		itemMetal.InventoryItemTitle(), metal,
		itemGas.InventoryItemTitle(), gas)
	msg := tgbotapi.NewMessage(avatar.Player.Chat, txt)
	msg.ReplyMarkup = menu.FactoryMenu
	Send(avatar.ChatID(), msg)
}

func factorySellRes(avatar *models.Avatar) {
	var metal = avatar.GetInventoryItemsByType(items.ITEM_METAL)
	var gas = avatar.GetInventoryItemsByType(items.ITEM_GAS)
	if len(metal) == 0 && len(gas) == 0 {
		msg := tgbotapi.NewMessage(avatar.Player.Chat, "В инвентаре нет ресурсов")
		msg.ReplyMarkup = menu.FactoryMenu
		Send(avatar.ChatID(), msg)
		return
	}

	var total int
	for _, it := range metal {
		total++
		avatar.InventoryRemoveItem(it.ID)
	}
	for _, it := range gas {
		total++
		avatar.InventoryRemoveItem(it.ID)
	}
	avatar.Wallet.Credits += total
	avatar.Wallet.UpdateMoney()
	txt := fmt.Sprintf("Ресурсы проданы на сумму 💰 %d\n", total)
	msg := tgbotapi.NewMessage(avatar.Player.Chat, txt)
	msg.ReplyMarkup = menu.BaseMenu
	Send(avatar.ChatID(), msg)
}

func factorySellJunk(avatar *models.Avatar) {

	if !avatar.InventoryHasJunk() {
		msg := tgbotapi.NewMessage(avatar.Player.Chat, "Хлама нет в инвентаре")
		msg.ReplyMarkup = menu.FactoryMenu
		Send(avatar.ChatID(), msg)
		return
	}
	junks := avatar.GetInventoryItemsByType(items.ITEM_JUNK)
	total := 0
	for _, it := range junks {
		total++
		avatar.InventoryRemoveItem(it.ID)
	}
	total /= 3
	avatar.Wallet.Credits += total
	avatar.Wallet.UpdateMoney()
	txt := fmt.Sprintf("Хлам продан на сумму 💰 %d\n", total)
	msg := tgbotapi.NewMessage(avatar.Player.Chat, txt)
	msg.ReplyMarkup = menu.BaseMenu
	Send(avatar.ChatID(), msg)
}

func factoryCraft(avatar *models.Avatar, craftID int) {
	if avatar.IsPrinting {
		Send(avatar.ChatID(), tgbotapi.NewMessage(avatar.ChatID(), "Печать уже идет!"))
		return
	}
	if !craft.IsReceiptInLocation(craftID, avatar.Player.Location) {
		Send(avatar.ChatID(), tgbotapi.NewMessage(avatar.ChatID(), "Здесь нет такого рецепта!"))
		return
	}
	// TODO add sql table
	receipt := craft.Get(craftID)
	if receipt == nil {
		Send(avatar.ChatID(), tgbotapi.NewMessage(avatar.ChatID(), "Здесь нет такого рецепта!"))
		return
	}

	if avatar.Wallet.Credits < receipt.Credits {
		Send(avatar.ChatID(), tgbotapi.NewMessage(avatar.ChatID(), "Недостаточно денег!"))
		return
	}
	for _, res := range receipt.Resources {
		if avatar.InventoryGetResourcesCount(res.ID) < res.Count {
			Send(avatar.ChatID(), tgbotapi.NewMessage(avatar.ChatID(), "Недостаточно ресурсов!"))
			return
		}
	}
	avatar.Wallet.Credits -= receipt.Credits
	avatar.Wallet.UpdateMoney()
	for _, res := range receipt.Resources {
		avatar.InventoryRemoveResourceCount(res.ID, res.Count)
	}
	item := items.GetByID(receipt.Item)
	txt := fmt.Sprintf("🖨 Печать %s начата, закончится через ⏱ %d м.\n", item.SimpleItemTitle(), int(receipt.Time.Minutes()))
	Send(avatar.ChatID(), tgbotapi.NewMessage(avatar.ChatID(), txt))
	avatar.IsPrinting = true
	go craftRoutine(avatar, receipt)
}

func craftRoutine(avatar *models.Avatar, receipt *craft.Craft) {
	<-time.After(receipt.Time)
	avatar.Lock()
	item := items.GetByID(receipt.Item)
	avatar.InventoryAddItem(item.ID)
	avatar.IsPrinting = false
	avatar.Unlock()
	txt := fmt.Sprintf("🖨 Печать %s выполнена, проверьте свой инвентарь\n", item.SimpleItemTitle())
	Send(avatar.ChatID(), tgbotapi.NewMessage(avatar.ChatID(), txt))
}
