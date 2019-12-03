package items

import "fmt"

type ItemType string
type ItemSubType string
type SocketType int

const (
	ITEM_WEAPON   ItemType = "weapon"
	ITEM_ARMOR    ItemType = "armor"
	ITEM_DEVICE   ItemType = "device"
	ITEM_METAL    ItemType = "metal"
	ITEM_GAS      ItemType = "gas"
	ITEM_JUNK     ItemType = "junk"
	ITEM_DUST     ItemType = "dust"
	ITEM_CRYSTALS ItemType = "crystals"
	ITEM_BALLOON  ItemType = "balloon"
	ITEM_CONSOLE  ItemType = "console"

	ITEM_SUB_BOER    ItemSubType = "boer"
	ITEM_SUB_NULL_T  ItemSubType = "nullt"
	ITEM_SUB_SCANNER ItemSubType = "scanner"

	// sockets
	ITEM_SOCKET_TYPE_DAMAGE SocketType = 1
	ITEM_SOCKET_TYPE_ARMOR  SocketType = 2
	ITEM_SOCKET_TYPE_HP     SocketType = 3
	ITEM_SOCKET_TYPE_MINE   SocketType = 4
	ITEM_SOCKET_TYPE_AIR    SocketType = 5
)

func (s SocketType) String() string {
	switch s {
	case ITEM_SOCKET_TYPE_DAMAGE:
		return "damage"
	case ITEM_SOCKET_TYPE_ARMOR:
		return "armor"
	case ITEM_SOCKET_TYPE_HP:
		return "hp"
	case ITEM_SOCKET_TYPE_MINE:
		return "mine"
	case ITEM_SOCKET_TYPE_AIR:
		return "air"
	}
	return "unknown"
}

func (s SocketType) Icon() string {
	switch s {
	case ITEM_SOCKET_TYPE_DAMAGE:
		return "⚫"
	case ITEM_SOCKET_TYPE_ARMOR:
		return "⚪"
	case ITEM_SOCKET_TYPE_HP:
		return "🔴"
	case ITEM_SOCKET_TYPE_MINE:
		return "🟠"
	case ITEM_SOCKET_TYPE_AIR:
		return "🔵"
	}
	return "🔘"
}

func (s SocketType) Name() string {
	switch s {
	case ITEM_SOCKET_TYPE_DAMAGE:
		return "Урон"
	case ITEM_SOCKET_TYPE_ARMOR:
		return "Броня"
	case ITEM_SOCKET_TYPE_HP:
		return "Очки жизни"
	case ITEM_SOCKET_TYPE_MINE:
		return "Скорость добычи"
	case ITEM_SOCKET_TYPE_AIR:
		return "Обьем баллонов"
	}
	return "unknown"
}

func (s ItemType) String() string {
	return string(s)
}

func (s ItemType) Name() string {
	switch s {
	case ITEM_ARMOR:
		return "Броня"
	case ITEM_WEAPON:
		return "Оружие"
	case ITEM_METAL:
		return "Металл"
	case ITEM_GAS:
		return "Газ"
	case ITEM_JUNK:
		return "Хлам"
	case ITEM_CRYSTALS:
		return "Кристаллы"
	case ITEM_DEVICE:
		return "Оборудование"
	case ITEM_BALLOON:
		return "Баллоны"
	}
	return "ЧТО ЭТО?"
}

func (s ItemSubType) Name() string {
	switch s {
	case ITEM_SUB_BOER:
		return "Бур"
	case ITEM_SUB_NULL_T:
		return "NULL-T"
	case ITEM_SUB_SCANNER:
		return "Сканер"
	}
	return "ЧТО ЭТО?"
}

type Item struct {
	ID      int
	Title   string
	Type    ItemType
	Tier    string
	SubType ItemSubType
	Price   int
	Sell    int

	Damage int
	Armor  int
	Air    int
	Mine   int
	Time   int

	Reputation int
	SocketType SocketType
	Sockets    int
}

type Equipment struct {
	ID         int
	Title      string
	Type       string
	SubType    string
	BuyPrice   int
	SellPrice  int
	Reputation int

	Damage int
	Armor  int
	Air    int
	Mine   int
	Time   int

	SocketType SocketType
	Sockets    int
}

func (s *Item) Icon() string {
	switch s.Type {
	case ITEM_ARMOR:
		return "🛡"
	case ITEM_WEAPON:
		return "⚔️"
	case ITEM_METAL:
		return "🧲"
	case ITEM_GAS:
		return "🛢️"
	case ITEM_JUNK:
		return "🧽"
	case ITEM_DEVICE:
		return "🛠"
	}
	return ""
}

func (s *Item) ShopItemTitle() string {
	switch s.Type {
	case ITEM_WEAPON:
		return fmt.Sprintf("%s %s %s%d 💰%d", s.Tier, s.Title, s.Icon(), s.Damage, s.Price)
	case ITEM_ARMOR:
		return fmt.Sprintf("%s %s %s%d 💰%d", s.Tier, s.Title, s.Icon(), s.Armor, s.Price)
	case ITEM_BALLOON:
		return fmt.Sprintf("%s %s 🌬️%d 💰%d", s.Tier, s.Title, s.Air, s.Price)
	case ITEM_DEVICE:
		switch s.SubType {
		case ITEM_SUB_BOER:
			return fmt.Sprintf("%s %s %s%d ⏱%d 💰%d", s.Tier, s.Title, s.Icon(), s.Mine, s.Time, s.Price)
		case ITEM_SUB_NULL_T:
			return fmt.Sprintf("%s %s  ⏱%d 💰%d", s.Tier, s.Title, s.Time, s.Price)
		case ITEM_SUB_SCANNER:
			return fmt.Sprintf("%s️ %s 💰%d", s.Tier, s.Title, s.Price)
		}

	case ITEM_METAL, ITEM_GAS, ITEM_JUNK:
		return fmt.Sprintf("%s %s 💰%d", s.Tier, s.Title, s.Price)
	}
	return ""
}

func (s *Item) DisassemblyItemTitle() string {
	switch s.Type {
	case ITEM_WEAPON:
		return fmt.Sprintf("%s️ %s %s%d /disass_%d", s.Tier, s.Title, s.Icon(), s.Damage, s.ID)
	case ITEM_ARMOR:
		return fmt.Sprintf("%s️ %s %s%d /disass_%d", s.Tier, s.Title, s.Icon(), s.Armor, s.ID)
	case ITEM_BALLOON:
		return fmt.Sprintf("%s️ %s 🌬️%d /disass_%d", s.Tier, s.Title, s.Air, s.ID)
	case ITEM_DEVICE:
		switch s.SubType {
		case ITEM_SUB_BOER:
			return fmt.Sprintf("%s️ %s %s%d ⏱%d /disass_%d", s.Tier, s.Title, s.Icon(), s.Mine, s.Time, s.ID)
		case ITEM_SUB_NULL_T:
			return fmt.Sprintf("%s️ %s ⏱%d /disass_%d", s.Tier, s.Title, s.Time, s.ID)
		case ITEM_SUB_SCANNER:
			return fmt.Sprintf("%s️ %s /disass_%d", s.Tier, s.Title, s.ID)
		}
	}
	return ""
}

func (s *Item) InventoryItemTitle() string {
	switch s.Type {
	case ITEM_WEAPON:
		return fmt.Sprintf("%s️ %s %s%d /equip_%d", s.Tier, s.Title, s.Icon(), s.Damage, s.ID)
	case ITEM_ARMOR:
		return fmt.Sprintf("%s️ %s %s%d /equip_%d", s.Tier, s.Title, s.Icon(), s.Armor, s.ID)
	case ITEM_BALLOON:
		return fmt.Sprintf("%s️ %s 🌬️%d /equip_%d", s.Tier, s.Title, s.Air, s.ID)
	case ITEM_METAL, ITEM_GAS, ITEM_JUNK:
		return fmt.Sprintf("%s️ %s", s.Tier, s.Title)
	case ITEM_DEVICE:
		switch s.SubType {
		case ITEM_SUB_BOER:
			return fmt.Sprintf("%s️ %s %s%d ⏱%d", s.Tier, s.Title, s.Icon(), s.Mine, s.Time)
		case ITEM_SUB_NULL_T:
			return fmt.Sprintf("%s️ %s ⏱%d", s.Tier, s.Title, s.Time)
		case ITEM_SUB_SCANNER:
			return fmt.Sprintf("%s️ %s", s.Tier, s.Title)
		}
	}
	return ""
}

func (s *Item) SimpleItemTitle() string {
	switch s.Type {
	case ITEM_WEAPON:
		return fmt.Sprintf("%s %s %s%d", s.Tier, s.Title, s.Icon(), s.Damage)
	case ITEM_ARMOR:
		return fmt.Sprintf("%s %s %s%d", s.Tier, s.Title, s.Icon(), s.Armor)
	case ITEM_METAL, ITEM_GAS, ITEM_JUNK:
		return fmt.Sprintf("%s %s", s.Tier, s.Title)
	case ITEM_DEVICE:
		switch s.SubType {
		case ITEM_SUB_BOER:
			return fmt.Sprintf("%s️ %s %s%d ⏱%d", s.Tier, s.Title, s.Icon(), s.Mine, s.Time)
		case ITEM_SUB_NULL_T:
			return fmt.Sprintf("%s️ %s ⏱%d", s.Tier, s.Title, s.Time)
		case ITEM_SUB_SCANNER:
			return fmt.Sprintf("%s️ %s", s.Tier, s.Title)
		}
	case ITEM_BALLOON:
		return fmt.Sprintf("%s %s 🌬️%d", s.Tier, s.Title, s.Air)
	}
	return ""
}

var Items []Item
