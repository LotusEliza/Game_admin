package telegram

import (
	"fmt"
	"github.com/finalist736/gokit/cache/ramcache"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"tcs"
	"tcs/machanics/story"
	"tcs/models"
	"tcs/telegram/menu"
	"time"
)

var cache = ramcache.New(time.Hour*24, time.Hour)

func avatarCacheName(tg int) string {
	return fmt.Sprintf("avatar_%d", tg)
}

func route(update tgbotapi.Update) {
	//logger.StdOut().Debugf("update: %+v", update)
	//logger.StdOut().Debugf("message: %+v", update.Message)
	//logger.StdOut().Debugf("entities: %+v", update.Message.Entities)

	if update.Message != nil {

		if update.Message.From == nil {
			Send(
				update.Message.Chat.ID,
				tgbotapi.NewMessage(
					update.Message.Chat.ID,
					"bot can work with private chat only!"))
			return
		}

		var (
			avatarData interface{}
			avatar     *models.Avatar
			exists     bool
			loaded     bool
		)

		avatarData, exists = cache.Get(avatarCacheName(update.Message.From.ID))
		if exists {
			avatar = avatarData.(*models.Avatar)
			cache.Set(avatarCacheName(avatar.Player.Tg), avatar)
		} else {
			avatar, loaded = models.LoadAvatar(update.Message.From.ID)
			if !loaded {
				player := models.NewPlayer()
				ok, referrerName := player.Register(update.Message, update.Message.Chat)
				if !ok {
					Send(
						update.Message.Chat.ID,
						tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка регистрации... Перезапустите бота!"))
				} else {
					// need to send referrer name!
					if referrerName != "" {
						refMsg := tgbotapi.NewMessage(
							update.Message.Chat.ID,
							fmt.Sprintf(
								"Вас пригласил %s, он получит награду",
								referrerName))
						Send(update.Message.Chat.ID, refMsg)
					}
					storyMessage := story.GetCurrentStory(player, update.Message.Chat.ID)
					if storyMessage != nil {
						Send(update.Message.Chat.ID, storyMessage)
					}
				}
				return
			} else {
				if avatar.Player.Location.IsBase() && !avatar.Player.IsAirFull() {
					go airRoutine(avatar)
				}
				//avatar.Dump()
				cache.Set(avatarCacheName(avatar.Player.Tg), avatar)
			}
		}

		//logger.StdOut().Debugf("update from: %+v", avatar.Player)
		// TODO uncomment updateLastActive later!
		avatar.Player.UpdateLastActive()

		// check for story
		if avatar.Player.Faction == 0 {
			// check for correct answer
			var showBaseMenu bool = false
			if update.Message.Text != "" {
				nextID := story.CheckCorrectAnswer(update.Message.Text, avatar.Player)
				if nextID == -1 {
					storyMessage := story.GetCurrentStory(avatar.Player, update.Message.Chat.ID)
					if storyMessage != nil {
						Send(update.Message.Chat.ID, storyMessage)
						return
					}
				}
				avatar.Player.Story = nextID
				avatar.Player.UpdateStory()
				// check end of story
				if story.CheckEndOfStory(avatar.Player) {
					// set additional quests
					showBaseMenu = true
				}
			}
			storyMessage := story.GetCurrentStory(avatar.Player, update.Message.Chat.ID)
			if storyMessage != nil {
				Send(update.Message.Chat.ID, storyMessage)
				if showBaseMenu {
					baseMsg := tgbotapi.NewMessage(avatar.Player.Chat, "Ты находишься на базе")
					baseMsg.ReplyMarkup = menu.BaseMenu
					Send(avatar.Player.Chat, baseMsg)
				}
				return
			}
			return
		}
		avatar.Lock()
		if update.Message.IsCommand() {
			// command
			commands(avatar, &update)
		} else {
			found := Run(update.Message.Text, avatar)
			if !found {
				// try devices messages!
				msg, ok := tryOutsideDevices(avatar, update.Message.Text)
				if ok {
					Send(avatar.ChatID(), msg)
					avatar.Unlock()
					return
				}
				// try coordinates
				if avatar.WaitCoordinates {
					tryCoordinates(avatar, update.Message.Text)
					avatar.Unlock()
					return
				}

				// send help message
				sendHelpMessage(avatar)
			}
		}
		avatar.Unlock()

	}

	//Send(update.Message.Chat.ID, tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text))
}

func sendHelpMessage(avatar *models.Avatar) {
	msg := tgbotapi.NewMessage(
		avatar.Player.Chat,
		fmt.Sprintf(
			`*Tagoreans Crash Site* _v%s_
Telegram Bot MMORPG
_Автор вдохновлялся произведением А. Б. Стругацких "Жук в муравейнике".
Данная игра является фанфиком этого произведения_
[Читать книгу](http://www.rusf.ru/abs/books/zhvm00.htm)
[Статья в Wikipedia](https://ru.wikipedia.org/wiki/Жук_в_муравейнике)
`,
			strategy.Version))
	if avatar.IsDuty() {
		msg.ReplyMarkup = menu.StateOnlyMenu
	} else {
		if avatar.Player.Location.IsBase() {
			msg.ReplyMarkup = menu.BaseMenu
		} else {
			msg.ReplyMarkup = menu.OutSideMenu
		}
	}
	msg.ParseMode = "Markdown"
	Send(avatar.Player.Chat, msg)
}
