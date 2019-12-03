package models

import (
	"github.com/finalist736/gokit/database"
	"github.com/finalist736/gokit/logger"
	"github.com/gocraft/dbr"
)

type PlayerWallet struct {
	Tg         int `db:"tg"`
	Credits    int `db:"credits"`
	Gold       int `db:"gold"`
	Reputation int `db:"reputation"`
}

var playerWalletColumns = []string{
	"tg",
	"credits",
	"gold",
}

const PlayerWalletTableName = "player_wallet"

func NewPlayerWallet() *PlayerWallet {
	return new(PlayerWallet)
}

func WalletLoadByPlayerID(id int) *PlayerWallet {
	session := database.GetDefaultSession()
	wallet := NewPlayerWallet()
	err := session.
		Select("*").
		From("player_wallet").
		Where("tg=?", id).
		LoadOne(wallet)
	if err != nil {
		if err != dbr.ErrNotFound {
			logger.StdErr().Errorf("select player_wallet error: %s", err)
		}
	}
	return wallet
}

func (s *PlayerWallet) UpdateMoney() {
	session := database.GetDefaultSession()
	_, err := session.Update("player_wallet").
		Set("credits", s.Credits).
		Set("gold", s.Gold).
		Set("reputation", s.Reputation).
		Where("tg=?", s.Tg).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("update credits player_wallet error: %s", err)
	}
}
