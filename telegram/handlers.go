package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"tcs/models"
	"tcs/telegram/menu"
)

const alreadyOnDuty = "♟️ Нельзя ходить в наряд когда уже в наряде!"

type handler func(*models.Avatar)

var handlers = make(map[string]handler)

func init() {
	handlers[menu.MENU_STATE] = state
	handlers[menu.MENU_LOCATION] = location

	// base menu
	handlers[menu.MENU_SHOP] = shop
	handlers[menu.MENU_SHOP_ARMOR] = shopArmor
	handlers[menu.MENU_SHOP_DEVICES] = shopDevices
	handlers[menu.MENU_SHOP_WEAPON] = shopWeapon
	handlers[menu.MENU_SHOP_BALLOONS] = shopBalloons
	handlers[menu.MENU_HQ_NORM] = hqNorm
	handlers[menu.MENU_HQ_PUT_NORM] = hqPutNorm
	handlers[menu.MENU_HQ_SELECT_DUTY] = hqDuty
	handlers[menu.MENU_HQ_DUTY_CANTEEN] = hqDutyCanteen
	handlers[menu.MENU_HQ_DUTY_GEO] = hqDutyGeo
	handlers[menu.MENU_HQ_DUTY_PATROL] = hqDutyPatrol

	handlers[menu.MENU_BACK] = back
	handlers[menu.MENU_CANTEEN] = canteen
	handlers[menu.MENU_INVENTORY] = inventory
	handlers[menu.MENU_HQ] = hq
	handlers[menu.MENU_EXIT] = exit
	handlers[menu.MENU_FACTORY] = factory

	// factory
	handlers[menu.MENU_FACTORY_SELLJUNK] = factorySellJunk
	handlers[menu.MENU_FACTORY_SELLRES] = factorySellRes
	handlers[menu.MENU_FACTORY_DISASSEMBLY] = factoryDisassemblyList
	handlers[menu.MENU_FACTORY_CRAFT] = factoryCraftList

	// move handlers
	handlers[menu.MENU_OUTSIDE_GO_NORD] = moveNord
	handlers[menu.MENU_OUTSIDE_GO_SOUTH] = moveSouth
	handlers[menu.MENU_OUTSIDE_GO_EAST] = moveEast
	handlers[menu.MENU_OUTSIDE_GO_WEST] = moveWest
	handlers[menu.MENU_OUTSIDE_ENTER_COORDS] = moveCoords
	handlers[menu.MENU_OUTSIDE_DEVICES] = outsideDevices

	// mine menu
	handlers[menu.MENU_OUTSIDE_MINE_GAS] = mineGas
	handlers[menu.MENU_OUTSIDE_MINE_RESOURCES] = mineResources
	handlers[menu.MENU_OUTSIDE_MINE_RANDOM] = mineRandom

	handlers[menu.MENU_OUTSIDE_NULLT_GO] = nullTransporting
}

func Run(txt string, avatar *models.Avatar) bool {
	method, found := handlers[txt]
	if found {
		if !avatar.IsDuty() {
			method(avatar)
			return true
		}
		if txt != menu.MENU_STATE &&
			txt != menu.MENU_BACK &&
			txt != menu.MENU_INVENTORY {
			msg := tgbotapi.NewMessage(avatar.Player.Chat, avatar.StateMessage())
			msg.ReplyMarkup = menu.StateOnlyMenu
			Send(avatar.ChatID(), msg)
			return true
		}
		method(avatar)
		return true
	}
	return false
}

func state(avatar *models.Avatar) {
	msg := tgbotapi.NewMessage(avatar.Player.Chat, avatar.StateMessage())
	msg.ReplyMarkup = menu.StateMenu
	Send(avatar.ChatID(), msg)
}

func location(avatar *models.Avatar) {
	msg := tgbotapi.NewMessage(avatar.Player.Chat, avatar.LocationMessage())
	msg.ReplyMarkup = menu.MoveMenu
	Send(avatar.ChatID(), msg)
}

func shop(avatar *models.Avatar) {
	msg := tgbotapi.NewMessage(avatar.ChatID(), avatar.ShopMessage())
	msg.ReplyMarkup = menu.ShopMenu
	//avatar.MenuHistory.Push(&menu.BaseMenu)

	Send(avatar.ChatID(), msg)
}

func shopArmor(avatar *models.Avatar) {
	msg := tgbotapi.NewMessage(avatar.ChatID(), avatar.ShopArmorMessage())
	msg.ReplyMarkup = menu.ShopBackMenu
	//avatar.MenuHistory.Push(&menu.ShopMenu)

	Send(avatar.ChatID(), msg)
}

func shopDevices(avatar *models.Avatar) {
	msg := tgbotapi.NewMessage(avatar.ChatID(), avatar.ShopDevicesMessage())
	msg.ReplyMarkup = menu.ShopBackMenu
	Send(avatar.ChatID(), msg)
}

func shopBalloons(avatar *models.Avatar) {
	msg := tgbotapi.NewMessage(avatar.ChatID(), avatar.ShopBalloonsMessage())
	msg.ReplyMarkup = menu.ShopBackMenu
	Send(avatar.ChatID(), msg)
}

func shopWeapon(avatar *models.Avatar) {
	msg := tgbotapi.NewMessage(avatar.ChatID(), avatar.ShopWeaponMessage())
	msg.ReplyMarkup = menu.ShopBackMenu
	//avatar.MenuHistory.Push(&menu.ShopMenu)

	Send(avatar.ChatID(), msg)
}

func inventory(avatar *models.Avatar) {
	txt := avatar.InventoryMessage()
	//logger.StdOut().Debugf("INV MESS: %s", txt)
	msg := tgbotapi.NewMessage(avatar.ChatID(), txt)
	msg.ReplyMarkup = menu.SimpleBackMenu
	//msg.ParseMode = "Markdown"
	Send(avatar.ChatID(), msg)
}

func exit(avatar *models.Avatar) {
	if avatar.Player.Air == 0 || avatar.Player.EquipB == 0 {
		Send(avatar.ChatID(), tgbotapi.NewMessage(avatar.ChatID(), "👮 Дежурный по кессону грудью загородил выход: без кислорода не пущу!!!"))
		return
	}

	if !exitToPlanet(avatar) {
		msg := tgbotapi.NewMessage(avatar.ChatID(), "🌏 Ты уже находишься на поверхности планеты!")
		msg.ReplyMarkup = menu.OutSideMenu
		//avatar.MenuHistory.Push(&menu.StateMenu)
		Send(avatar.ChatID(), msg)
		return
	}
	avatar.Player.UpdateLocation()
	avatar.SetMoving()
	go movingRoutine(avatar)
	msg := tgbotapi.NewMessage(avatar.ChatID(), "🚪 Проход через кессон занимает время\n🌬️ Следите за уровнем кислорода в баллонах!\nРасход: 1 🌬️ на 1 клетку\nИспользование оборудования: 1 🌬️ за одно действие")
	msg.ReplyMarkup = menu.MoveMenu
	Send(avatar.ChatID(), msg)
}

func back(avatar *models.Avatar) {

	//if avatar.MenuHistory.Len() == 0 {
	//	logger.StdOut().Debugf("showing default menu")
	//	// show default menu
	switch avatar.Player.Location {
	case 1, 2:
		// Default base menu
		msg := tgbotapi.NewMessage(avatar.ChatID(), avatar.LocationMessage())
		if avatar.IsDuty() {
			msg.ReplyMarkup = menu.StateOnlyMenu
		} else {
			msg.ReplyMarkup = menu.BaseMenu
		}
		Send(avatar.ChatID(), msg)
		return
	default:
		// default outside menu
		msg := tgbotapi.NewMessage(avatar.ChatID(), avatar.LocationMessage())
		if avatar.IsMining {
			msg.ReplyMarkup = menu.StateOnlyMenu
		} else if avatar.IsMoving() {
			msg.ReplyMarkup = menu.MoveMenu
		} else {
			msg.ReplyMarkup = menu.OutSideMenu
		}
		//avatar.MenuHistory.Push(&menu.OutSideMenu)
		Send(avatar.ChatID(), msg)
		return
	}
	//} else {
	//	logger.StdOut().Debugf("showing previous menu")
	//	prev := avatar.MenuHistory.Pop().(*tgbotapi.ReplyKeyboardMarkup)
	//	message := "some message"
	//	switch prev {
	//	case &menu.ShopMenu:
	//		message = avatar.ShopMessage()
	//	default:
	//		message = avatar.StateMessage()
	//	}
	//	msg := tgbotapi.NewMessage(avatar.ChatID(), message)
	//	msg.ReplyMarkup = prev
	//	//avatar.MenuHistory.Push(prev)
	//	Send(avatar.ChatID(), msg)
	//}
}
