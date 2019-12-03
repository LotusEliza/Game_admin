package models

import (
	"github.com/finalist736/gokit/database"
	"github.com/finalist736/gokit/logger"
	"github.com/gocraft/dbr"
	"tcs/machanics/duty"
	"time"
)

type PlayerDuty struct {
	Tg        int       `db:"tg"`
	TimeStart time.Time `db:"timestart"`
	TimeEnd   time.Time `db:"timeend"`
	Duty      int       `db:"duty"`
}

const PlayerDutyTableName = "player_duty"

var PlayerDutyColumns = []string{
	"tg",
	"timestart",
	"timeend",
	"duty",
}

func NewPlayerDuty() *PlayerDuty {
	return &PlayerDuty{}
}

func DutyLoadLast(tg int) (*PlayerDuty, error) {
	pld := NewPlayerDuty()
	session := database.GetDefaultSession()
	err := session.Select("*").From(PlayerDutyTableName).Where("tg=?", tg).OrderBy("timestart desc").LoadOne(pld)
	if err != nil {
		if err != dbr.ErrNotFound {
			logger.StdErr().Errorf("DutyLoadLast select player_duty error: %s", err)
		}
		return nil, err
	}
	return pld, nil
}

func CreateDuty(tg, d int) *PlayerDuty {
	pld := NewPlayerDuty()

	pld.Tg = tg
	pld.TimeStart = time.Now()
	pld.TimeEnd = duty.GetTimeForDuty(d)
	pld.Duty = d

	session := database.GetDefaultSession()
	_, err := session.InsertInto(PlayerDutyTableName).Columns(PlayerDutyColumns...).Record(pld).Exec()
	if err != nil {
		logger.StdErr().Errorf("CreateDuty insert player_duty error: %s", err)
		return nil
	}
	return pld
}

func (s *PlayerDuty) Message() string {
	var msg = ""
	switch s.Duty {
	case duty.DUTY_CANTEEN:
		msg = "🍜 в столовой"
	case duty.DUTY_GEO:
		msg = "🛠 в геологической рейде"
	case duty.DUTY_PATROL:
		msg = "🚶‍♂ в патруле"
	}
	return msg
}
