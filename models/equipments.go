package models

const EquipmentsTableName = "equipments"

type Equipments struct {
	ID         int    `db:"id"`
	Title      string `db:"title"`
	Type       string `db:"type"`
	SubType    string `db:"subtype"`
	BuyPrice   int    `db:"buyprice"`
	SellPrice  int    `db:"sellprice"`
	Reputation int    `db:"reputation"`

	Damage int `db:"damage"`
	Armor  int `db:"armor"`
	Air    int `db:"air"`
	Mine   int `db:"mine"`
	Time   int `db:"time"`

	SocketType int `db:"socketype"`
	Sockets    int `db:"sockets"`
}

var EquipmentsColumns = []string{
	"title",
	"type",
	"subtype",
	"buyprice",
	"sellprice",
	"reputation",
	"damage",
	"armor",
	"air",
	"mine",
	"time",
	"socketype",
	"sockets",
}
