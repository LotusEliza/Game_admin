package models

import (
	"github.com/finalist736/gokit/database"
	"github.com/finalist736/gokit/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/gocraft/dbr"
	"strconv"
	"strings"
	"tcs/items"
	"time"
)

type Player struct {
	Tg         int       `db:"tg"`
	Chat       int64     `db:"chat"`
	Name       string    `db:"name"`
	Story      int       `db:"story"`
	Faction    int       `db:"faction"`
	Referrer   int       `db:"referrer"`
	Location   Location  `db:"location"`
	PosX       int       `db:"posx"`
	PosY       int       `db:"posy"`
	HP         int       `db:"hp"`
	EquipW     int       `db:"equipw"`
	EquipA     int       `db:"equipa"`
	EquipB     int       `db:"equipb"`
	EquipC     int       `db:"equipc"`
	Air        int       `db:"air"`
	Registered time.Time `db:"registered"`
	LastActive time.Time `db:"lastactive"`
}

var PlayerColumns = []string{
	"tg",
	"chat",
	"name",
	"story",
	"faction",
	"referrer",
	"location",
	"posx",
	"posy",
	"hp",
	"equipw",
	"equipa",
	"equipb",
	"equipc",
	"air",
	"registered",
	"lastactive",
}

const PlayerTableName = "players"

func NewPlayer() *Player {
	return new(Player)
}

func PlayerLoad(tg int) (*Player, bool) {
	player := NewPlayer()
	session := database.GetDefaultSession()
	err := session.
		Select("*").
		From(PlayerTableName).
		Where("tg=?", tg).
		LoadOne(player)
	if err != nil {
		if err != dbr.ErrNotFound {
			logger.StdErr().Errorf("select player error: %s", err)
		}
		return nil, false
	}
	return player, true
}

func (s *Player) IsAirFull() bool {
	item := items.GetByID(s.EquipB)
	if item == nil {
		return true
	}
	if s.Air >= item.Air {
		return true
	}
	return false
}

func (s *Player) Register(message *tgbotapi.Message, chat *tgbotapi.Chat) (bool, string) {
	session := database.GetDefaultSession()
	s.Tg = message.From.ID
	s.Chat = chat.ID
	s.Faction = 0
	s.Story = 1
	s.Name = message.From.UserName
	s.Registered = time.Now()
	s.LastActive = s.Registered
	s.Location = 0
	s.HP = 10
	s.EquipA = 4
	s.EquipW = 1
	s.Air = 40
	s.EquipB = 22

	var err error
	var referrerName string
	//var referrerID int

	// check referrer
	if strings.Contains(message.Text, "/start") {
		parts := strings.Split(message.Text, " ")
		if len(parts) == 2 {
			s.Referrer, err = strconv.Atoi(parts[1])
			if err == nil {
				referrer, exists := PlayerLoad(s.Referrer)
				if exists {
					referrerName = referrer.Name
					s.Referrer = referrer.Tg
				} else {
					logger.StdOut().Debugf(
						"unknown referrerID str: %s; int: %d",
						parts[1],
						s.Referrer)
					s.Referrer = 0
				}
			}
		}
	}

	tx, err := session.Begin()

	if err != nil {
		logger.StdErr().Errorf("tx begin error: %s", err)
		return false, referrerName
	}
	defer tx.RollbackUnlessCommitted()
	_, err = tx.
		InsertInto("players").
		Columns(PlayerColumns...).
		Record(s).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("insert player error: %s", err)
		return false, referrerName
	}

	wallet := NewPlayerWallet()
	wallet.Tg = s.Tg
	wallet.Credits = 40
	wallet.Gold = 10

	_, err = tx.
		InsertInto("player_wallet").
		Columns(playerWalletColumns...).
		Record(wallet).Exec()
	if err != nil {
		logger.StdErr().Errorf("insert player_wallet error: %s", err)
		return false, referrerName
	}

	bur1 := NewPlayerInventory()
	bur1.Tg = s.Tg
	bur1.ItemType = items.ITEM_DEVICE.String()
	bur1.ItemValue = 20
	_, err = tx.InsertInto(PlayerInventoryTableName).
		Columns(PlayerInventoryColumns...).
		Record(bur1).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("insert player_inventory error: %s", err)
		return false, referrerName
	}

	err = tx.Commit()
	if err != nil {
		logger.StdErr().Errorf("tx commit error: %s", err)
		return false, referrerName
	}
	return true, referrerName
}

func (s *Player) UpdateLastActive() {
	session := database.GetDefaultSession()
	s.LastActive = time.Now()
	_, err := session.
		Update("players").
		Set("lastactive", s.LastActive).
		Where("tg=?", s.Tg).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("update lastactive player error: %s", err)
	}
}

func (s *Player) UpdateFaction() {
	session := database.GetDefaultSession()
	_, err := session.Update("players").
		Set("faction", s.Faction).
		Where("tg=?", s.Tg).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("update faction player error: %s", err)
	}
}

func (s *Player) UpdateHP() {
	session := database.GetDefaultSession()
	_, err := session.Update("players").
		Set("hp", s.HP).
		Where("tg=?", s.Tg).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("update location player error: %s", err)
	}
}

func (s *Player) UpdateLocation() {
	session := database.GetDefaultSession()
	_, err := session.Update("players").
		Set("location", s.Location).
		Set("posx", s.PosX).
		Set("posy", s.PosY).
		Where("tg=?", s.Tg).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("update location player error: %s", err)
	}
}

func (s *Player) UpdateAir() {
	session := database.GetDefaultSession()
	_, err := session.Update("players").
		Set("air", s.Air).
		Where("tg=?", s.Tg).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("update air player error: %s", err)
	}
}

func (s *Player) UpdateStory() {
	session := database.GetDefaultSession()
	_, err := session.Update("players").
		Set("story", s.Story).
		Where("tg=?", s.Tg).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("update story player error: %s", err)
	}
}

func (s *Player) UpdateEquipArmor() {
	session := database.GetDefaultSession()
	_, err := session.Update("players").
		Set("equipa", s.EquipA).
		Where("tg=?", s.Tg).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("update equipa player error: %s", err)
	}
}

func (s *Player) UpdateEquipWeapon() {
	session := database.GetDefaultSession()
	_, err := session.Update("players").
		Set("equipw", s.EquipW).
		Where("tg=?", s.Tg).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("update equipw player error: %s", err)
	}
}

func (s *Player) UpdateEquipBalloon() {
	session := database.GetDefaultSession()
	_, err := session.Update("players").
		Set("equipb", s.EquipB).
		Where("tg=?", s.Tg).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("update equipw player error: %s", err)
	}
}

func (s *Player) UpdateEquipConsole() {
	session := database.GetDefaultSession()
	_, err := session.Update("players").
		Set("equipc", s.EquipC).
		Where("tg=?", s.Tg).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("update equipw player error: %s", err)
	}
}

func (s *Player) FactionName() string {
	switch s.Faction {
	case 0:
		return "Без фракции"
	case 1:
		return "Корпорация"
	case 2:
		return "Пираты"
	default:
		return "Unknown"
	}
}
