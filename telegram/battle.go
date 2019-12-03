package telegram

import (
	"fmt"
	"github.com/finalist736/gokit/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"math/rand"
	"tcs/items"
	"tcs/models"
	"tcs/telegram/menu"
	"tcs/tools"
)

const playerPunchFirstPercent = 70
const startBattleMsg = "На тебя напал %s! (❤️ %d)\nПридется сражаться\n\n"

func mobBattle(avatar *models.Avatar, mob *models.Mob) {
	logger.StdOut().Debugf("MOB BATTLE: user: %+v; mob: %+v", avatar.Player, mob)
	msg := fmt.Sprintf(startBattleMsg, mob.Title, mob.HP)
	playerPunch := true
	if rand.Intn(100) >= playerPunchFirstPercent {
		playerPunch = false
	}
	mobHp := mob.HP
	playerWeapon := items.GetByID(avatar.Player.EquipW)
	playerArmor := items.GetByID(avatar.Player.EquipA)
	var damage int
	for mobHp > 0 && avatar.Player.HP > 0 {
		if playerPunch {
			playerPunch = false
			damage = playerWeapon.Damage - mob.Armor
			if damage <= 0 {
				damage = 1
			}
			mobHp -= damage

			msg += fmt.Sprintf("👊🏼 Ты нанёс удар: %d\n", damage)

		} else {
			playerPunch = true
			damage = tools.RandMinMax(mob.Damage[0], mob.Damage[1]) - playerArmor.Armor
			if damage <= 0 {
				damage = 1
			}
			avatar.Player.HP -= damage
			msg += fmt.Sprintf("💥 %s нанёс удар: %d\n", mob.Title, damage)
		}
		//logger.StdOut().Debugf("punch result: userHP: %d; mobHP: %d", avatar.Player.HP, mobHp)
	}
	avatar.Player.UpdateHP()

	if mobHp <= 0 {
		msg += "🎖 Ты победил!\n"

		lootCount := tools.RandMinMax(mob.LootCount[0], mob.LootCount[1])
		msg += "\nДобыча:\n"
		var (
			itemID int
			item   *items.Item
		)
		for ; lootCount > 0; lootCount-- {
			itemID = mob.Loot[rand.Intn(len(mob.Loot))]
			item = items.GetByID(itemID)
			msg += fmt.Sprintf("%s\n", item.InventoryItemTitle())
			avatar.InventoryAddItem(itemID)
		}
	} else {
		msg += "💀 Ты проиграл!\nТебя доставили на базу на носилках"
		baseLocation(avatar)
	}
	msg = avatar.HealthMessage() + "\n" + avatar.LocationMiniMessage() + "\n\n" + msg
	msgConfig := tgbotapi.NewMessage(avatar.ChatID(), msg)
	if avatar.Player.Location.IsBase() {
		msgConfig.ReplyMarkup = menu.BaseMenu
		avatar.Player.HP = 10
		avatar.Player.UpdateHP()
	}

	Send(avatar.ChatID(), msgConfig)

}
