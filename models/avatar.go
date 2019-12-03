package models

import (
	"fmt"
	"github.com/finalist736/gokit/database"
	"github.com/finalist736/gokit/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
	"sync"
	"tcs/items"
	"tcs/machanics/daytime"
	"time"
)

type Avatar struct {
	*sync.Mutex
	Player    *Player
	Wallet    *PlayerWallet
	Inventory []*PlayerInventory

	isMoving   bool
	IsFilling  bool
	IsMining   bool
	IsScanning bool
	IsPrinting bool

	// timers
	LastNullT      time.Time
	CanteenMessage time.Time

	// coordinates
	WaitCoordinates bool
	Destination     [2]int

	UsingDevice *items.Item

	// duty
	Duty *PlayerDuty

	//MenuHistory *stack.Stack
}

type inventoryItemsCount struct {
	item  *items.Item
	count int
}

func LoadAvatar(tg int) (*Avatar, bool) {
	var ava = new(Avatar)
	ava.Mutex = new(sync.Mutex)
	var exists bool
	ava.Player, exists = PlayerLoad(tg)
	if !exists {
		return nil, false
	}
	ava.Wallet = WalletLoadByPlayerID(tg)
	ava.Inventory = InventoryLoadByPlayerID(tg)

	ava.CanteenMessage = time.Now().Add(-2 * time.Minute)
	ava.LastNullT = time.Now().Add(-100 * time.Minute)

	ava.Duty, _ = DutyLoadLast(tg)

	return ava, true
}

func (s *Avatar) IsDuty() bool {
	if s.Duty == nil {
		return false
	}
	if s.Duty.TimeEnd.Unix() > time.Now().Unix() {
		return true
	}
	return false
}

func (s *Avatar) SetMoving() {
	s.isMoving = true
}

func (s *Avatar) SetUnMoving() {
	s.isMoving = false
}

func (s *Avatar) IsMoving() bool {
	return s.isMoving
}

func (s *Avatar) Reset() {
	session := database.GetDefaultSession()
	_, err := session.DeleteFrom("players").Where("tg=?", s.Player.Tg).Exec()
	if err != nil {
		logger.StdErr().Errorf("delete from players error: %s", err)
		return
	}
	_, err = session.DeleteFrom("player_wallet").Where("tg=?", s.Wallet.Tg).Exec()
	if err != nil {
		logger.StdErr().Errorf("delete from player_wallet error: %s", err)
		return
	}
	_, err = session.DeleteFrom("player_inventory").Where("tg=?", s.Player.Tg).Exec()
	if err != nil {
		logger.StdErr().Errorf("delete from player_wallet error: %s", err)
		return
	}
	_, err = session.DeleteFrom("player_norms").Where("tg=?", s.Player.Tg).Exec()
	if err != nil {
		logger.StdErr().Errorf("delete from player_wallet error: %s", err)
		return
	}
	//s.MenuHistory = stack.New()
}

func (s *Avatar) LocationMiniMessage() string {
	return fmt.Sprintf(
		"Место: %s\nВремя суток: %s\nКоординаты: [%d:%d]",
		s.Player.Location.Name(),
		daytime.Get(),
		s.Player.PosX,
		s.Player.PosY)
}

func (s *Avatar) HealthMessage() string {

	return fmt.Sprintf("❤️ %d/%d 🌬️ %d",
		s.Player.HP,
		10,
		s.Player.Air)
}

func (s *Avatar) MoneyMessage() string {
	return fmt.Sprintf("💰 %d 💎 %d 🏆 %d", s.Wallet.Credits, s.Wallet.Gold, s.Wallet.Reputation)
}

func (s *Avatar) HqMessage() string {
	return "🏤 Штаб!"
}

func (s *Avatar) HqNormMessage() string {
	return "🏆 Чтоб получить репутацию нужно каждый день сдавать норму ресурсов\n"
}

func (s *Avatar) DevicesMessage() string {
	return "Список устройств для использования на поверхности планеты"
}

//func GetButtonsFromAnswers(answers []Answer) tgbotapi.ReplyKeyboardMarkup {
//	var keyboard [][]tgbotapi.KeyboardButton
//	var keys []tgbotapi.KeyboardButton
//	for i, it := range answers {
//		if i+1%2 == 0 {
//			keyboard = append(keyboard, keys)
//			keys = make([]tgbotapi.KeyboardButton, 0)
//		}
//		keys = append(keys, tgbotapi.NewKeyboardButton(it.Answer))
//	}
//	if len(keys) > 0 {
//		keyboard = append(keyboard, keys)
//	}
//	//logger.StdOut().Debugf("keyboard: %+v", keyboard)
//	return tgbotapi.NewReplyKeyboard(keyboard...)
//}

func (s *Avatar) HasDevices() bool {
	return len(s.GetInventoryItemsByType(items.ITEM_DEVICE)) > 0
}

func (s *Avatar) OutsideDevicesMenu() tgbotapi.ReplyKeyboardMarkup {
	var keyboard [][]tgbotapi.KeyboardButton
	var keys []tgbotapi.KeyboardButton

	devices := s.GetInventoryItemsByType(items.ITEM_DEVICE)
	for i, dev := range devices {
		if i+1%2 == 0 {
			keyboard = append(keyboard, keys)
			keys = make([]tgbotapi.KeyboardButton, 0)
		}
		keys = append(keys, tgbotapi.NewKeyboardButton(dev.InventoryItemTitle()))
	}
	if len(keys) > 0 {
		keyboard = append(keyboard, keys)
	}
	//keys = make([]tgbotapi.KeyboardButton, 0)
	//keys = append(keys, tgbotapi.NewKeyboardButton(menu.MENU_BACK))
	//keyboard = append(keyboard, keys)
	return tgbotapi.NewReplyKeyboard(keyboard...)
}

func (s *Avatar) StateMessage() string {
	eqW := items.GetByID(s.Player.EquipW)
	eqA := items.GetByID(s.Player.EquipA)
	eqB := items.GetByID(s.Player.EquipB)
	eqC := items.GetByID(s.Player.EquipC)

	eqWStr := "-"
	eqAStr := "-"
	eqBStr := "-"
	eqCStr := "-"

	if eqW != nil {
		eqWStr = eqW.SimpleItemTitle()
	}

	if eqA != nil {
		eqAStr = eqA.SimpleItemTitle()
	}
	if eqB != nil {
		eqBStr = eqB.SimpleItemTitle()
	}
	if eqC != nil {
		eqCStr = eqC.SimpleItemTitle()
	}

	var dutyText = ""
	if s.IsDuty() {
		dutyText = "\nНаряд: " + s.Duty.Message()
	}

	return fmt.Sprintf(`Ваше состояние: 
id: %d
Время суток: %s
💰 %d 💎 %d 🏆 %d

❤️ %d/%d 🌬️ %d

Экипировка:
Оружие: %s
Броня: %s
Баллоны: %s
Консоль: %s

Фракция: %s
Место: %s
Координаты: [%d:%d]%s`,
		s.Player.Tg,
		daytime.Get(),
		s.Wallet.Credits,
		s.Wallet.Gold,
		s.Wallet.Reputation,
		s.Player.HP,
		10,
		s.Player.Air,
		eqWStr,
		eqAStr,
		eqBStr,
		eqCStr,
		s.Player.FactionName(),
		s.Player.Location.Name(),
		s.Player.PosX,
		s.Player.PosY,
		dutyText,
	)
}

func (s *Avatar) ChatID() int64 {
	return s.Player.Chat
}

func (s *Avatar) CanBuy(id int) bool {
	item := items.GetByID(id)
	if item == nil {
		return false
	}
	if s.Wallet.Credits < item.Price {
		return false
	}
	return true
}

func (s *Avatar) BuyItem(itemID int) (*items.Item, bool) {
	item := items.GetByID(itemID)
	if item == nil {
		return nil, false
	}
	plinv := NewPlayerInventory()
	plinv.Tg = s.Player.Tg
	plinv.ItemType = item.Type.String()
	plinv.ItemValue = item.ID
	if plinv.CreateItem() {
		s.Wallet.Credits -= item.Price
		s.Wallet.UpdateMoney()
		s.Inventory = append(s.Inventory, plinv)
		return item, true
	} else {
		return nil, false
	}
}

func (s *Avatar) ShopMessage() string {
	return fmt.Sprintf(`Магазин!
Здесь можно приобрести по сносной цене: оружие, скафандры, разное оборудование и может даже динамит!
`)
}

func (s *Avatar) ShopArmorMessage() string {
	its := items.GetByType(items.ITEM_ARMOR)
	itemsText := ""
	for _, it := range its {
		itemsText += fmt.Sprintf("%s /buy_%d\n", it.ShopItemTitle(), it.ID)
	}
	return fmt.Sprintf(`%s

Броня!

%s

`, s.MoneyMessage(), itemsText)
}

func (s *Avatar) ShopWeaponMessage() string {
	its := items.GetByType(items.ITEM_WEAPON)
	itemsText := ""
	for _, it := range its {
		itemsText += fmt.Sprintf("%s /buy_%d\n", it.ShopItemTitle(), it.ID)
	}
	return fmt.Sprintf(`%s

Оружие!

%s

`, s.MoneyMessage(), itemsText)
}

func (s *Avatar) ShopBalloonsMessage() string {
	its := items.GetCorpBaseBalloons()
	itemsText := ""
	for _, it := range its {
		itemsText += fmt.Sprintf("%s /buy_%d\n", it.ShopItemTitle(), it.ID)
	}
	return fmt.Sprintf(`%s

Устройства!

%s

`, s.MoneyMessage(), itemsText)
}

func (s *Avatar) ShopDevicesMessage() string {
	its := items.GetCorpBaseDevices()
	itemsText := ""
	for _, it := range its {
		itemsText += fmt.Sprintf("%s /buy_%d\n", it.ShopItemTitle(), it.ID)
	}
	return fmt.Sprintf(`%s

Устройства!

%s

`, s.MoneyMessage(), itemsText)
}

func (s *Avatar) ShopResourcesMessage() string {
	its := items.GetByType(items.ITEM_METAL)
	itemsText := ""
	for _, it := range its {
		itemsText += fmt.Sprintf("%s /buy_%d\n", it.ShopItemTitle(), it.ID)
	}
	return fmt.Sprintf(`%s

Ресурсы!

%s

`, s.MoneyMessage(), itemsText)
}

func (s *Avatar) ShopGasMessage() string {
	its := items.GetByType(items.ITEM_GAS)
	itemsText := ""
	for _, it := range its {
		itemsText += fmt.Sprintf("%s /buy_%d\n", it.ShopItemTitle(), it.ID)
	}
	return fmt.Sprintf(`%s

Газ!

%s

`, s.MoneyMessage(), itemsText)
}

func (s *Avatar) InventoryMessage() string {
	if len(s.Inventory) == 0 {
		return fmt.Sprintf(
			"%s\n\nИнвентарь:\nЗаглянув сюда ты видишь множество пустоты в своем чемодане!",
			s.MoneyMessage())
	}

	var showItemsType = []items.ItemType{items.ITEM_WEAPON, items.ITEM_ARMOR, items.ITEM_DEVICE, items.ITEM_BALLOON, items.ITEM_METAL, items.ITEM_GAS, items.ITEM_JUNK}
	typesText := ""
	for _, typ := range showItemsType {
		its := s.GetInventoryItemsByType(typ)
		if len(its) == 0 {
			continue
		}
		typesText += typ.Name() + ":\n"

		var cnt string
		switch typ {
		case items.ITEM_JUNK, items.ITEM_METAL, items.ITEM_GAS:
			var itmcnt []*inventoryItemsCount
			for _, it := range its {
				if !isItemInInventory(itmcnt, it) {
					itmcnt = append(itmcnt, &inventoryItemsCount{item: it, count: 1})
				}
			}
			for _, it := range itmcnt {
				cnt = strconv.Itoa(it.count)
				typesText += it.item.InventoryItemTitle() + "(" + cnt + "); "
			}
		default:
			for _, it := range its {
				typesText += it.InventoryItemTitle() + "\n"
			}
		}
		typesText += "\n"
	}
	return fmt.Sprintf(
		"%s\n\nИнвентарь:\n%s",
		s.MoneyMessage(), typesText)
}

func (s *Avatar) InventoryGetResourcesCount(id int) (cnt int) {
	for _, it := range s.Inventory {
		if it.ItemValue != id {
			continue
		}
		cnt++
	}
	return
}

func (s *Avatar) GetInventoryItemsByType(t items.ItemType) []*items.Item {
	var its []*items.Item
	for _, it := range s.Inventory {
		if it.ItemType != t.String() {
			continue
		}
		item := items.GetByID(it.ItemValue)
		its = append(its, item)
	}
	return its
}

func (s *Avatar) InventoryHasJunk() bool {
	for _, it := range s.Inventory {
		if it.ItemType == items.ITEM_JUNK.String() {
			return true
		}
	}
	return false
}

func (s *Avatar) InventoryHasDevices() bool {
	for _, it := range s.Inventory {
		if it.ItemType == items.ITEM_DEVICE.String() ||
			it.ItemType == items.ITEM_WEAPON.String() ||
			it.ItemType == items.ITEM_ARMOR.String() ||
			it.ItemType == items.ITEM_BALLOON.String() {
			return true
		}
	}
	return false
}

func (s *Avatar) InventoryGetItem(id int) *PlayerInventory {
	for _, it := range s.Inventory {
		if it.ItemValue == id {
			return it
		}
	}
	return nil
}

func (s *Avatar) InventoryAddItem(id int) {
	item := items.GetByID(id)
	plinv := NewPlayerInventory()
	plinv.Tg = s.Player.Tg
	plinv.ItemValue = item.ID
	plinv.ItemType = item.Type.String()
	plinv.CreateItem()
	s.Inventory = append(s.Inventory, plinv)
}

func (s *Avatar) InventoryRemoveResourceCount(id, count int) {
	for count > 0 {
		for i, it := range s.Inventory {
			if it.ItemValue == id {
				count--
				if it.RemoveItem() {
					s.Inventory = append(s.Inventory[:i], s.Inventory[i+1:]...)
				}
				break
			}
		}
	}
}

func (s *Avatar) InventoryRemoveItem(id int) {
	item := items.GetByID(id)
	if item == nil {
		return
	}
	for index, it := range s.Inventory {
		if it.ItemValue != id {
			continue
		}
		if it.RemoveItem() {
			s.Inventory = append(s.Inventory[:index], s.Inventory[index+1:]...)
		}
		break
	}
}

func (s *Avatar) LocationMessage() string {
	if s.Player.Location.IsBase() {
		var dutyText = ""
		if s.IsDuty() {
			dutyText = "\nНаряд: " + s.Duty.Message()
		}
		return fmt.Sprintf(
			"%s\nМесто: %s\nВремя суток: %s%s",
			s.HealthMessage(),
			s.Player.Location.Name(), daytime.Get(), dutyText)
	} else {
		doingMessage := ""
		if s.isMoving {
			doingMessage += "\n⏳🏃 Ты куда-то идешь..."
		} else if s.IsMining {
			doingMessage += "\n⏳🛠 Ты что-то майнишь..."
		}
		return fmt.Sprintf(
			"%s\nМесто: %s\nВремя суток: %s\nКоординаты: [%d:%d]%s",
			s.HealthMessage(),
			s.Player.Location.Name(),
			daytime.Get(),
			s.Player.PosX,
			s.Player.PosY,
			doingMessage)
	}
}

func (s *Avatar) InventoryEquipItem(i *PlayerInventory) bool {
	for _, it := range s.Inventory {
		if it.ItemValue == i.ItemValue {
			item := items.GetByID(i.ItemValue)
			if item.Type != items.ITEM_WEAPON &&
				item.Type != items.ITEM_ARMOR &&
				item.Type != items.ITEM_BALLOON &&
				item.Type != items.ITEM_CONSOLE {
				return false
			}
			switch item.Type {
			case items.ITEM_ARMOR:
				if s.Player.EquipA != 0 {
					s.InventoryAddItem(s.Player.EquipA)
				}
				s.InventoryRemoveItem(item.ID)
				s.Player.EquipA = item.ID
				s.Player.UpdateEquipArmor()
			case items.ITEM_WEAPON:
				if s.Player.EquipW != 0 {
					s.InventoryAddItem(s.Player.EquipW)
				}
				s.InventoryRemoveItem(item.ID)
				s.Player.EquipW = item.ID
				s.Player.UpdateEquipWeapon()
			case items.ITEM_BALLOON:
				if s.Player.EquipB != 0 {
					s.InventoryAddItem(s.Player.EquipB)
				}
				s.InventoryRemoveItem(item.ID)
				s.Player.EquipB = item.ID
				s.Player.UpdateEquipBalloon()
			case items.ITEM_CONSOLE:
			}

			return true
		}
	}
	return false
}

//func (s *Avatar) InventoryUnEquipItem(i *PlayerInventory) bool {
//	for index, it := range s.Inventory {
//		if it.ItemValue == i.ItemValue {
//			s.Inventory = append(s.Inventory[:index], s.Inventory[index+1:]...)
//			return true
//		}
//	}
//	return false
//}

func (s *Avatar) Dump() {
	logger.StdOut().Debugf("  ----  ")
	logger.StdOut().Debugf("Player avatar")
	logger.StdOut().Debugf("Player: %+v", s.Player)
	logger.StdOut().Debugf("Wallet: %+v", s.Wallet)
	logger.StdOut().Debugf("Inventory: ")
	for _, inv := range s.Inventory {
		logger.StdOut().Debugf("    Item: %+v", inv)
	}
	logger.StdOut().Debug("  ----  ")
}

func isItemInInventory(its []*inventoryItemsCount, it *items.Item) bool {
	for _, i := range its {
		if i.item.ID == it.ID {
			i.count++
			return true
		}
	}
	return false
}
