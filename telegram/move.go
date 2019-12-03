package telegram

import (
	"fmt"
	"github.com/finalist736/gokit/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"tcs/models"
	"tcs/telegram/menu"
	"time"
)

const moveToCoordinates = "Куда идти?\nВведите координаты в формате - 12:12 или 1 19\nИли любой текст для отмены"
const movementErrorMessage = "Ты сделал пару кругов по коридорам базы и вернулись туда же, дневальный что-то заподозрил..."
const airLessMessage = "🌬️ За кислородом не следил!\n💀 Задохнулся!\nТебя подобрал патруль. На базе еле откачали!"
const wrongWayMessage = "🚷 Дальше нет дороги! Выбери другое направление!"
const cantAction = "🚷 Нельзя выполнить это действие!"
const enterTheBase = "🏰 Возвращаюсь на базу"

const moveTimeout = time.Second * 3
const airTimeout = time.Second * 10

var randomMovingMessages = []string{"🏃‍♂️Выдвигаюсь!", "🏃‍♂️Принял, понял, осознал", "🏃‍♂️Ну, я пошел...", "🏃‍♂️Тебя понял", "🏃‍♂️В дорогу!"}

func moveNord(avatar *models.Avatar) {
	moving(avatar, models.DIR_NORD)
}

func moveSouth(avatar *models.Avatar) {
	moving(avatar, models.DIR_SOUTH)
}
func moveEast(avatar *models.Avatar) {
	moving(avatar, models.DIR_EAST)
}
func moveWest(avatar *models.Avatar) {
	moving(avatar, models.DIR_WEST)
}

func moveCoords(avatar *models.Avatar) {
	avatar.WaitCoordinates = true
	msg := tgbotapi.NewMessage(avatar.Player.Chat, moveToCoordinates)
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	Send(avatar.ChatID(), msg)
}

func tryCoordinates(avatar *models.Avatar, txt string) {
	if strings.ContainsAny(txt, ": ") {
		var parts []string

		parts = strings.Split(txt, ":")
		if len(parts) != 2 {
			parts = strings.Split(txt, " ")
		}
		if len(parts) == 2 {
			// try parse numbers
			var (
				x, y int
				err  error
			)
			x, err = strconv.Atoi(parts[0])
			if err == nil {
				y, err = strconv.Atoi(parts[1])
				if err == nil {
					if avatar.Player.PosX != x || avatar.Player.PosY != y {
						avatar.Destination[0] = x
						avatar.Destination[1] = y
						go movingCoordinatesRoutine(avatar)
						msg := tgbotapi.NewMessage(avatar.Player.Chat, "🏃‍♂ Выдвигаюсь! О любых изменениям буду докладывать!")
						msg.ReplyMarkup = menu.MoveMenu
						Send(avatar.ChatID(), msg)
						avatar.WaitCoordinates = false
						return
					}
				}
			}
		}
	}
	msg := tgbotapi.NewMessage(avatar.Player.Chat, avatar.LocationMessage())
	msg.ReplyMarkup = menu.OutSideMenu
	Send(avatar.ChatID(), msg)
	avatar.WaitCoordinates = false
	return
}

func moving(avatar *models.Avatar, dir models.Direction) {
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
	if avatar.Player.Location.IsBase() {
		msg := tgbotapi.NewMessage(avatar.Player.Chat, movementErrorMessage)
		msg.ReplyMarkup = menu.BaseMenu
		Send(avatar.ChatID(), msg)
		return
	}

	if avatar.Player.Air == 0 {
		//avatar.Player.HP = 0
		msg := tgbotapi.NewMessage(avatar.Player.Chat, airLessMessage)
		baseLocation(avatar)
		msg.ReplyMarkup = menu.BaseMenu
		Send(avatar.ChatID(), msg)
		avatar.Player.UpdateLocation()
		return
	}

	switch dir {
	case models.DIR_NORD:
		if avatar.Player.PosY == 1 {
			msg := tgbotapi.NewMessage(avatar.Player.Chat, wrongWayMessage)
			msg.ReplyMarkup = menu.OutSideMenu
			Send(avatar.ChatID(), msg)
			return
		}
		avatar.Player.PosY--
	case models.DIR_SOUTH:
		height := models.LocationsSizes[avatar.Player.Location].Height
		if avatar.Player.PosY == height {
			msg := tgbotapi.NewMessage(avatar.Player.Chat, wrongWayMessage)
			msg.ReplyMarkup = menu.OutSideMenu
			Send(avatar.ChatID(), msg)
			return
		}
		avatar.Player.PosY++
	case models.DIR_EAST:
		if avatar.Player.PosX == models.LocationsSizes[avatar.Player.Location].Width {
			msg := tgbotapi.NewMessage(avatar.Player.Chat, wrongWayMessage)
			msg.ReplyMarkup = menu.OutSideMenu
			Send(avatar.ChatID(), msg)
			return
		}
		avatar.Player.PosX++
	case models.DIR_WEST:
		if avatar.Player.PosX == 1 {
			msg := tgbotapi.NewMessage(avatar.Player.Chat, wrongWayMessage)
			msg.ReplyMarkup = menu.OutSideMenu
			Send(avatar.ChatID(), msg)
			return
		}
		avatar.Player.PosX--
	}

	avatar.Player.Air--
	avatar.Player.UpdateAir()
	avatar.Player.UpdateLocation()

	avatar.SetMoving()
	//TODO changing location after timer to prevent cheating!
	go movingRoutine(avatar)

	msg := tgbotapi.NewMessage(avatar.Player.Chat, randomMovingMessages[rand.Intn(len(randomMovingMessages))])
	msg.ReplyMarkup = menu.MoveMenu
	Send(avatar.ChatID(), msg)
	switch avatar.Player.Faction {
	case 1:
		if avatar.Player.PosX == 1 && avatar.Player.PosY == 1 {
			Send(avatar.ChatID(), tgbotapi.NewMessage(avatar.ChatID(), enterTheBase))
			avatar.Player.Location = models.CorporationBase
		}
	case 2:
		if avatar.Player.PosX == 25 && avatar.Player.PosY == 25 {
			Send(avatar.ChatID(), tgbotapi.NewMessage(avatar.ChatID(), enterTheBase))
			avatar.Player.Location = models.PiratesBase
		}
	}

}

func determineCoordinates(avatar *models.Avatar) {
	var diffX, diffY int

	diffX = int(math.Abs(float64(avatar.Player.PosX - avatar.Destination[0])))
	diffY = int(math.Abs(float64(avatar.Player.PosY - avatar.Destination[1])))

	if diffX > diffY {
		if avatar.Player.PosX > avatar.Destination[0] {
			avatar.Player.PosX--
		} else {
			avatar.Player.PosX++
		}
	} else {
		if avatar.Player.PosY > avatar.Destination[1] {
			avatar.Player.PosY--
		} else {
			avatar.Player.PosY++
		}
	}
}

func movingCoordinatesRoutine(avatar *models.Avatar) {

	avatar.Lock()
	logger.StdOut().Debugf("moving to [%d:%d] from [%d:%d]",
		avatar.Destination[0],
		avatar.Destination[1],
		avatar.Player.PosX,
		avatar.Player.PosY)
	for avatar.Player.PosX != avatar.Destination[0] ||
		avatar.Player.PosY != avatar.Destination[1] {

		determineCoordinates(avatar)
		avatar.Player.Air--

		avatar.Player.UpdateAir()
		avatar.Player.UpdateLocation()

		if avatar.Player.Air == 0 {
			//avatar.Player.HP = 0
			msg := tgbotapi.NewMessage(avatar.Player.Chat, airLessMessage)
			baseLocation(avatar)
			msg.ReplyMarkup = menu.BaseMenu
			Send(avatar.ChatID(), msg)
			logger.StdOut().Debug("dead with air")
			avatar.Player.UpdateLocation()
			avatar.Unlock()
			return
		}
		avatar.SetMoving()
		avatar.Unlock()
		<-time.After(moveTimeout)
		avatar.Lock()
		//msg := tgbotapi.NewMessage(avatar.ChatID(), avatar.LocationMessage())
		//Send(avatar.ChatID(), msg)
		mob, exists := models.GetCurrentMob(avatar.Player.Location)
		if exists {
			logger.StdOut().Debugf("mob: %+v", mob)
			mobBattle(avatar, mob)
			if avatar.Player.Location.IsBase() {
				//avatar.Unlock()
				logger.StdOut().Debug("dead with mob")
				break
			}
			msg := tgbotapi.NewMessage(avatar.ChatID(), "🏃‍♂ Ты продолжил путь")
			msg.ReplyMarkup = menu.MoveMenu
			Send(avatar.ChatID(), msg)
		}

	}
	avatar.SetUnMoving()

	switch avatar.Player.Faction {
	case 1:
		if avatar.Player.PosX == 1 && avatar.Player.PosY == 1 {
			Send(avatar.ChatID(), tgbotapi.NewMessage(avatar.ChatID(), enterTheBase))
			avatar.Player.Location = models.CorporationBase
		}
	case 2:
		if avatar.Player.PosX == 25 && avatar.Player.PosY == 25 {
			Send(avatar.ChatID(), tgbotapi.NewMessage(avatar.ChatID(), enterTheBase))
			avatar.Player.Location = models.PiratesBase
		}
	}

	logger.StdOut().Debug("moving end")
	msg := tgbotapi.NewMessage(avatar.ChatID(), avatar.LocationMessage())
	if avatar.Player.Location.IsBase() {
		msg.ReplyMarkup = menu.BaseMenu
		baseLocation(avatar)
	} else {
		msg.ReplyMarkup = menu.OutSideMenu
	}
	Send(avatar.ChatID(), msg)
	avatar.Unlock()
}

func movingRoutine(avatar *models.Avatar) {
	<-time.After(moveTimeout)
	avatar.Lock()
	defer avatar.Unlock()
	avatar.SetUnMoving()
	msg := tgbotapi.NewMessage(avatar.Player.Chat, avatar.LocationMessage())
	if avatar.Player.Location.IsBase() {
		baseLocation(avatar)
		msg.ReplyMarkup = menu.BaseMenu
		Send(avatar.ChatID(), msg)
	} else {
		msg.ReplyMarkup = menu.OutSideMenu
		Send(avatar.ChatID(), msg)
		mob, exists := models.GetCurrentMob(avatar.Player.Location)
		if exists {
			logger.StdOut().Debugf("mob: %+v", mob)
			mobBattle(avatar, mob)
		}
	}
}

func airRoutine(avatar *models.Avatar) {
	logger.StdOut().Debug("airRoutine")
	avatar.Lock()
	if avatar.IsFilling {
		logger.StdOut().Debug("airRoutine already filling")
		avatar.Unlock()
		return
	}
	if avatar.Player.IsAirFull() {
		avatar.IsFilling = false
		if avatar.Player.Location.IsBase() {
			Send(avatar.ChatID(), tgbotapi.NewMessage(avatar.ChatID(), fmt.Sprintf("🌬️ %d\nБаллоны заполнены!", avatar.Player.Air)))
		}
		logger.StdOut().Debug("airRoutine full ballons")
		avatar.Unlock()
		return
	}
	logger.StdOut().Debug("airRoutine filling.....")
	avatar.IsFilling = true
	avatar.Unlock()
	<-time.After(airTimeout)
	logger.StdOut().Debug("airRoutine tick timer")
	avatar.Lock()
	if !avatar.Player.Location.IsBase() {
		logger.StdOut().Debug("airRoutine not base, go home!")
		avatar.IsFilling = false
		avatar.Unlock()
		return
	}
	avatar.Player.Air++
	avatar.Player.UpdateAir()
	avatar.IsFilling = false
	avatar.Unlock()
	logger.StdOut().Debug("airRoutine try run next routine")
	go airRoutine(avatar)
}

func baseLocation(avatar *models.Avatar) {
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

	avatar.Player.UpdateLocation()
	go airRoutine(avatar)
}

func exitToPlanet(avatar *models.Avatar) bool {
	if avatar.Player.Location.IsBase() {
		w, ok := models.BasesExits[avatar.Player.Location]
		if !ok {
			return false
		}
		avatar.Player.Location = w.Loc
		avatar.Player.PosX = w.PosX
		avatar.Player.PosY = w.PosY
		return true
	}
	return false
}
