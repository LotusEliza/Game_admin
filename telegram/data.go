package telegram

import (
	"container/list"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"sync"
	"time"
)

type (
	chattable struct {
		chatID  int64
		message tgbotapi.Chattable
	}
	history struct {
		chatID int64
		when   time.Time
	}
)

const (
	sendBuffer = 10000
	sendPeriod = 100
)

var (
	stopChannel   chan interface{}
	senderChannel chan *chattable
	bot           *tgbotapi.BotAPI
	mux           sync.Mutex
	messageBuffer = list.New()

	chattablePool = sync.Pool{
		New: func() interface{} {
			//logger.StdOut().Debug("chattable created")
			return new(chattable)
		},
	}

	historyPool = sync.Pool{
		New: func() interface{} {
			//logger.StdOut().Debug("history created")
			return new(history)
		},
	}
)

func init() {
	stopChannel = make(chan interface{})
	senderChannel = make(chan *chattable, sendBuffer)
}
