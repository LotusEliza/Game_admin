package menu

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

var BaseMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(MENU_STATE),
		tgbotapi.NewKeyboardButton(MENU_SHOP),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(MENU_HQ),
		tgbotapi.NewKeyboardButton(MENU_CANTEEN),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(MENU_FACTORY),
		tgbotapi.NewKeyboardButton(MENU_EXIT),
	),
)

var HeadQuarterMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(MENU_HQ_NORM),
		tgbotapi.NewKeyboardButton(MENU_HQ_SELECT_DUTY),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(MENU_BACK),
	),
)

var HeadQuarterNormMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(MENU_HQ_PUT_NORM),
		tgbotapi.NewKeyboardButton(MENU_BACK),
	),
)

var HeadQuarterDutyMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(MENU_HQ_DUTY_CANTEEN),
		tgbotapi.NewKeyboardButton(MENU_HQ_DUTY_GEO),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(MENU_HQ_DUTY_PATROL),
		tgbotapi.NewKeyboardButton(MENU_BACK),
	),
)

var StateMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(MENU_INVENTORY),
		tgbotapi.NewKeyboardButton(MENU_BACK),
	),
)

var OutSideMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(MENU_OUTSIDE_GO_NORD),
		tgbotapi.NewKeyboardButton(MENU_OUTSIDE_GO_SOUTH),
		tgbotapi.NewKeyboardButton(MENU_OUTSIDE_GO_EAST),
		tgbotapi.NewKeyboardButton(MENU_OUTSIDE_GO_WEST),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(MENU_OUTSIDE_ENTER_COORDS),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(MENU_OUTSIDE_DEVICES),
		tgbotapi.NewKeyboardButton(MENU_STATE),
	),
)

var ChoiceResourceMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(MENU_OUTSIDE_MINE_GAS),
		tgbotapi.NewKeyboardButton(MENU_OUTSIDE_MINE_RESOURCES),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(MENU_OUTSIDE_MINE_RANDOM),
		tgbotapi.NewKeyboardButton(MENU_BACK),
	),
)

var MoveMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(MENU_STATE),
		tgbotapi.NewKeyboardButton(MENU_LOCATION),
	),
)

var StateOnlyMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(MENU_STATE),
	),
)

var ShopMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(MENU_SHOP_WEAPON),
		tgbotapi.NewKeyboardButton(MENU_SHOP_ARMOR),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(MENU_SHOP_DEVICES),
		tgbotapi.NewKeyboardButton(MENU_SHOP_BALLOONS),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(MENU_BACK),
	),
)

var FactoryMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(MENU_FACTORY_SELLJUNK),
		tgbotapi.NewKeyboardButton(MENU_FACTORY_DISASSEMBLY),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(MENU_FACTORY_SELLRES),
		tgbotapi.NewKeyboardButton(MENU_FACTORY_CRAFT),
	),
	tgbotapi.NewKeyboardButtonRow(
		//tgbotapi.NewKeyboardButton(MENU_FACTORY_SELLRES),
		tgbotapi.NewKeyboardButton(MENU_BACK),
	),
)

var ShopBackMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(MENU_SHOP),
		tgbotapi.NewKeyboardButton(MENU_BACK),
	),
)

//
//var BackMenu = tgbotapi.NewReplyKeyboard(
//	tgbotapi.NewKeyboardButtonRow(
//		tgbotapi.NewKeyboardButton(MENU_BACK),
//	),
//)
//
//var CanteenMenu = tgbotapi.NewReplyKeyboard(
//	tgbotapi.NewKeyboardButtonRow(
//		tgbotapi.NewKeyboardButton(MENU_BACK),
//	),
//)

var SimpleBackMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(MENU_BACK),
	),
)

var NULLtMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(MENU_OUTSIDE_NULLT_GO),
		tgbotapi.NewKeyboardButton(MENU_BACK),
	),
)
