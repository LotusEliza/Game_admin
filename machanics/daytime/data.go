package daytime

import (
	"github.com/finalist736/gokit/logger"
	"sync"
	"time"
)

type (
	TimePart bool
	dayTime  struct {
		sync.RWMutex
		part TimePart
	}
)

const (
	Day   TimePart = true
	Night TimePart = false

	dayTimePeriod = time.Hour * 2
)

var (
	defaultDayTime *dayTime

	dayTimes = []int{0, 4, 8, 12, 16, 20}
)

func Init() {
	defaultDayTime = new(dayTime)
	go defaultDayTime.routine()
}

func (s TimePart) String() string {
	switch s {
	case Day:
		return "День"
	case Night:
		return "Ночь"
	}
	return "Не понятно что за время?"
}

func (s *dayTime) get() TimePart {
	s.RLock()
	defer s.RUnlock()
	return s.part
}

func (s *dayTime) routine() {

	currentPeriod := time.Nanosecond
	now := time.Now()
	hour := now.Hour()
	if hour%2 == 0 {
		hour += 2
		if !isDayTime(hour) {
			s.part = Day
		} else {
			s.part = Night
		}
	} else {
		hour++
		if !isDayTime(hour) {
			s.part = Day
		} else {
			s.part = Night
		}
	}
	next := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		hour,
		0,
		0,
		0,
		now.Location())
	currentPeriod = next.Sub(now) + (time.Millisecond * 100)
	logger.StdOut().Debugf("next2: %s; dayTime: %s", currentPeriod.String(), s.part)

	for {
		now = <-time.After(currentPeriod)
		currentPeriod = dayTimePeriod
		s.Lock()
		if isDayTime(now.Hour()) {
			s.part = Day
		} else {
			s.part = Night
		}
		s.Unlock()
		logger.StdOut().Debugf("now: %s; dayTime: %s", now.String(), s.part)
	}
}

func Get() TimePart {
	return defaultDayTime.get()
}

func isDayTime(t int) bool {
	for _, it := range dayTimes {
		if it == t {
			return true
		}
	}
	return false
}
