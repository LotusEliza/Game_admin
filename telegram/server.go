package telegram

import (
	"github.com/finalist736/gokit/logger"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func Init(tgbotapiKey string) error {
	var (
		err        error
		updChannel tgbotapi.UpdatesChannel

		updConfig tgbotapi.UpdateConfig
		botUser   tgbotapi.User
	)
	bot, err = tgbotapi.NewBotAPI(tgbotapiKey)
	if err != nil {
		return err
	}

	bot, err = tgbotapi.NewBotAPI(tgbotapiKey)
	if err != nil {
		return err
	}

	botUser, err = bot.GetMe()
	if err != nil {
		return err
	}

	logger.StdOut().Infof("auth ok! bot is: %s\n", botUser.FirstName)

	updConfig.Timeout = 60
	updConfig.Limit = 1
	updConfig.Offset = 0

	updChannel, err = bot.GetUpdatesChan(updConfig)
	if err != nil {
		return err
	}

	go readRoutine(updChannel)
	go receiveRoutine()
	go sendRoutine()

	return nil
}

func readRoutine(updChannel tgbotapi.UpdatesChannel) {
	var update tgbotapi.Update
	for {
		select {
		case update = <-updChannel:
			go route(update)
		case <-stopChannel:
			logger.StdOut().Infof("stoping tg bot...")
			stopChannel <- nil
			return
		}
	}
}

func Stop() {
	stopChannel <- nil
	<-stopChannel
}
