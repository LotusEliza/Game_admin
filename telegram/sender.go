package telegram

import (
	"github.com/finalist736/gokit/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"time"
)

func Send(chatID int64, message tgbotapi.Chattable) {
	if bot == nil {
		return
	}
	chatt := chattablePool.Get().(*chattable)
	chatt.chatID = chatID
	chatt.message = message
	senderChannel <- chatt
}

func sendRoutine() {
	var (
		err error

		sendHistory []*history
		i           int
		it          *history

		userBusy bool

		message     *chattable
		sendTicker  = time.After(time.Millisecond)
		now         time.Time
		needRestart bool
	)

	sendHistory = make([]*history, 0, 30)

	for {
		select {
		case now = <-sendTicker:
			mux.Lock()
			// wait for next messages
			if len(sendHistory) == 0 && messageBuffer.Len() == 0 {
				sendTicker = time.After(time.Millisecond * sendPeriod)
				mux.Unlock()
				continue
			}
			//logger.StdOut().Debugf("history: %d", len(sendHistory))
			// clean sendHistory for old sends
			for {
				needRestart = false
				for i, it = range sendHistory {
					if now.Sub(it.when) > time.Second {
						sendHistory = append(sendHistory[:i], sendHistory[i+1:]...)
						historyPool.Put(it)
						needRestart = true
					}
				}
				if !needRestart {
					break
				}
			}

			if messageBuffer.Len() == 0 {
				sendTicker = time.After(time.Millisecond * sendPeriod)
				mux.Unlock()
				continue
			}
			//logger.StdOut().Debugf("buffer: %d", messageBuffer.Len())
			removeList := make([]*chattable, 0, messageBuffer.Len())
			// check for one user send timeout
			for e := messageBuffer.Front(); e != nil; e = e.Next() {

				// too many sends for last second, wait!
				if len(sendHistory) >= 30 {
					break
				}

				message = e.Value.(*chattable)
				// check for sends to such user
				userBusy = false
				for _, it = range sendHistory {
					if it.chatID == message.chatID {
						userBusy = true
						break
					}
				}
				// skip this message
				if userBusy {
					continue
				}

				_, err = bot.Send(message.message)
				if err != nil {
					logger.StdOut().Errorf("bot send error: %s", err)
				}
				//logger.StdOut().Debug("message sent")
				//messageBuffer.Remove(e)
				removeList = append(removeList, message)
				hi := historyPool.Get().(*history)
				hi.chatID = message.chatID
				hi.when = time.Now()
				sendHistory = append(sendHistory, hi)
			}

			for _, it2 := range removeList {
				for e := messageBuffer.Front(); e != nil; e = e.Next() {
					message = e.Value.(*chattable)
					if it2 == message {
						messageBuffer.Remove(e)
						chattablePool.Put(message)
						break
					}
				}
			}
			sendTicker = time.After(time.Millisecond * sendPeriod)
			mux.Unlock()
		}
	}
}

func receiveRoutine() {
	var (
		message *chattable
	)
	for {
		select {
		case message = <-senderChannel:
			mux.Lock()
			messageBuffer.PushBack(message)
			mux.Unlock()
		}
	}
}
