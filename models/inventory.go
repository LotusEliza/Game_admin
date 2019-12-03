package models

import (
	"github.com/finalist736/gokit/database"
	"github.com/finalist736/gokit/logger"
)

type PlayerInventory struct {
	Tg        int    `db:"tg"`
	ItemType  string `db:"itemtype"`
	ItemValue int    `db:"itemvalue"`
}

func NewPlayerInventory() *PlayerInventory {
	return new(PlayerInventory)
}

var PlayerInventoryColumns = []string{
	"tg",
	"itemtype",
	"itemvalue",
}

const PlayerInventoryTableName = "player_inventory"

func InventoryLoadByPlayerID(tg int) []*PlayerInventory {
	var inv []*PlayerInventory
	session := database.GetDefaultSession()
	_, err := session.Select("*").
		From(PlayerInventoryTableName).
		Where("tg=?", tg).
		OrderBy("itemtype").
		OrderBy("itemvalue").
		Load(&inv)
	if err != nil {
		logger.StdErr().Errorf("select player_inventory error: %s", err)
	}
	return inv
}

func (s *PlayerInventory) CreateItem() bool {
	session := database.GetDefaultSession()
	_, err := session.
		InsertInto(PlayerInventoryTableName).
		Columns(PlayerInventoryColumns...).
		Record(s).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("insert player_inventory error: %s", err)
		return false
	}
	return true
}

func (s *PlayerInventory) RemoveItem() bool {
	session := database.GetDefaultSession()
	_, err := session.
		DeleteFrom(PlayerInventoryTableName).
		Where("tg=? AND itemtype=? AND itemvalue=?", s.Tg, s.ItemType, s.ItemValue).
		Limit(1).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("delete player_inventory error: %s", err)
		return false
	}
	return true
}
