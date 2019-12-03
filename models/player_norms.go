package models

import (
	"github.com/finalist736/gokit/database"
	"github.com/finalist736/gokit/logger"
	"github.com/gocraft/dbr"
	"time"
)

// import "time"

type PlayerNorms struct {
	Tg       int    `db:"tg"`
	Date     string `db:"date"`
	Resource int    `db:"resource"`
	Amount   int    `db:"amount"`
}

var PlayerNormsColumns = []string{
	"tg",
	"date",
	"resource",
	"amount",
}

const PlayerNormsTableName = "player_norms"

func PlayerNormLoadCurrent(tg int) *PlayerNorms {
	session := database.GetDefaultSession()
	curDate := time.Now().Format("2006-01-02")
	playerNorm := &PlayerNorms{}
	err := session.Select("*").From(PlayerNormsTableName).Where("tg=? AND date=?", tg, curDate).LoadOne(playerNorm)
	if err != nil {
		if err != dbr.ErrNotFound {
			logger.StdErr().Errorf("loading player norms error: %s", err)
		}
		return nil
	}
	return playerNorm
}

func PlayerNormCreate(pn *PlayerNorms) error {
	session := database.GetDefaultSession()
	_, err := session.
		InsertInto(PlayerNormsTableName).
		Columns(PlayerNormsColumns...).
		Record(pn).
		Exec()
	if err != nil {
		return err
	}
	return nil
}
