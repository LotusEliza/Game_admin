package items

import (
	"github.com/finalist736/gokit/logger"
	"tcs/tools"
)

func init() {
	Items = append(Items, Item{
		ID:      1,
		Title:   "ПМ",
		Type:    ITEM_WEAPON,
		Tier:    "⚪",
		Price:   6,
		Sell:    2,
		Damage:  3,
		Sockets: 0,
	})
	Items = append(Items, Item{
		ID:      2,
		Title:   "ТТ",
		Type:    ITEM_WEAPON,
		Tier:    "⚪",
		Price:   50,
		Sell:    3,
		Damage:  5,
		Sockets: 0,
	})
	Items = append(Items, Item{
		ID:         3,
		Title:      "МП-40",
		Type:       ITEM_WEAPON,
		Tier:       "⚪",
		Price:      99,
		Sell:       4,
		Damage:     7,
		SocketType: ITEM_SOCKET_TYPE_DAMAGE,
		Sockets:    1,
	})

	Items = append(Items, Item{
		ID:    4,
		Title: "СКФ-98",
		Type:  ITEM_ARMOR,
		Tier:  "⚪",
		Price: 7,
		Sell:  3,
		Armor: 2,
		Air:   25,
	})
	Items = append(Items, Item{
		ID:    5,
		Title: "СКФ-98а",
		Type:  ITEM_ARMOR,
		Tier:  "⚪",
		Price: 60,
		Sell:  10,
		Armor: 3,
		Air:   30,
	})
	Items = append(Items, Item{
		ID:         6,
		Title:      "СКФ-07",
		Type:       ITEM_ARMOR,
		Tier:       "⚪",
		Price:      100,
		Sell:       18,
		Armor:      5,
		Air:        35,
		SocketType: ITEM_SOCKET_TYPE_ARMOR,
		Sockets:    1,
	})

	Items = append(Items, Item{
		ID:    7,
		Title: "Железо",
		Tier:  "⚪",
		Type:  ITEM_METAL,
		Price: 20,
		Sell:  14,
	})
	Items = append(Items, Item{
		ID:    8,
		Title: "Медь",
		Tier:  "⚪",
		Type:  ITEM_METAL,
		Price: 20,
		Sell:  13,
	})
	Items = append(Items, Item{
		ID:    9,
		Title: "Титан",
		Tier:  "⚪",
		Type:  ITEM_METAL,
		Price: 20,
		Sell:  15,
	})

	Items = append(Items, Item{
		ID:    10,
		Title: "Метан",
		Tier:  "⚪",
		Type:  ITEM_GAS,
		Price: 20,
		Sell:  14,
	})
	Items = append(Items, Item{
		ID:    11,
		Title: "Водород",
		Tier:  "⚪",
		Type:  ITEM_GAS,
		Price: 20,
		Sell:  14,
	})
	Items = append(Items, Item{
		ID:    12,
		Title: "Гелий",
		Tier:  "⚪",
		Type:  ITEM_GAS,
		Price: 20,
		Sell:  14,
	})

	// JUNK
	Items = append(Items, Item{
		ID:    13,
		Title: "Уголь",
		Tier:  "⚫",
		Type:  ITEM_JUNK,
		Price: 4,
		Sell:  1,
	})
	Items = append(Items, Item{
		ID:    14,
		Title: "Кварц",
		Tier:  "⚫",
		Type:  ITEM_JUNK,
		Price: 4,
		Sell:  1,
	})
	Items = append(Items, Item{
		ID:    15,
		Title: "Вода",
		Tier:  "⚫",
		Type:  ITEM_JUNK,
		Price: 4,
		Sell:  1,
	})
	Items = append(Items, Item{
		ID:    16,
		Title: "Гранат",
		Tier:  "⚫",
		Type:  ITEM_JUNK,
		Price: 4,
		Sell:  1,
	})
	Items = append(Items, Item{
		ID:    17,
		Title: "Янтарь",
		Tier:  "⚫",
		Type:  ITEM_JUNK,
		Price: 4,
		Sell:  1,
	})
	Items = append(Items, Item{
		ID:    18,
		Title: "Турмалин",
		Tier:  "⚫",
		Type:  ITEM_JUNK,
		Price: 4,
		Sell:  1,
	})
	Items = append(Items, Item{
		ID:    19,
		Title: "Обломки",
		Tier:  "⚫",
		Type:  ITEM_JUNK,
		Price: 4,
		Sell:  1,
	})

	// devices
	Items = append(Items, Item{
		ID:      20,
		Title:   "Бур-1",
		Type:    ITEM_DEVICE,
		SubType: ITEM_SUB_BOER,
		Tier:    "⚪",
		Price:   50,
		Sell:    30,
		Mine:    3,
		Time:    10,
	})
	Items = append(Items, Item{
		ID:      21,
		Title:   "NULL-Teleport",
		Type:    ITEM_DEVICE,
		SubType: ITEM_SUB_NULL_T,
		Tier:    "⚪",
		Price:   50,
		Sell:    30,
		Mine:    0,
		Time:    30,
	})

	// balloons
	Items = append(Items, Item{
		ID:    22,
		Title: "Баллон Б1",
		Type:  ITEM_BALLOON,
		Tier:  "⚪",
		Price: 30,
		Sell:  30,
		Air:   40,
	})
	Items = append(Items, Item{
		ID:    23,
		Title: "Баллон Б2",
		Type:  ITEM_BALLOON,
		Tier:  "⚪",
		Price: 300,
		Sell:  30,
		Air:   60,
	})
	Items = append(Items, Item{
		ID:    24,
		Title: "Баллон Б50",
		Type:  ITEM_BALLOON,
		Tier:  "⚪",
		Price: 1000,
		Sell:  30,
		Air:   80,
	})
	Items = append(Items, Item{
		ID:      25,
		Title:   "Бур-2",
		Tier:    "🔵",
		Type:    ITEM_DEVICE,
		SubType: ITEM_SUB_BOER,
		Price:   500,
		Sell:    30,
		Mine:    6,
		Time:    8,
	})
	Items = append(Items, Item{
		ID:      26,
		Title:   "Сканер",
		Tier:    "⚪",
		Type:    ITEM_DEVICE,
		SubType: ITEM_SUB_SCANNER,
		Price:   190,
		Sell:    60,
		Mine:    0,
		Time:    10,
	})

}

func GetByID(id int) *Item {
	for _, it := range Items {
		if it.ID == id {
			return &it
		}
	}
	return nil
}

func GetByType(t ItemType) []Item {
	var result []Item
	for _, it := range Items {
		if it.Type == t {
			result = append(result, it)
		}
	}
	return result
}

func Dump() {
	for _, it := range Items {
		logger.StdOut().Debugf("item: %+v", it)
	}
}

func GetJunkFor1Location() []Item {
	var result []Item
	for _, it := range Items {
		if it.ID == 13 || it.ID == 15 || it.ID == 14 || it.ID == 16 {
			result = append(result, it)
		}
	}
	return result
}

var itemsDevicesForCorpBase = []int{20, 25, 21, 26}
var itemsBalloonsForCorpBase = []int{22, 23, 24}

func GetCorpBaseDevices() []Item {
	var result []Item
	for _, it := range itemsDevicesForCorpBase {
		result = append(result, *GetByID(it))
	}
	return result
}

func GetCorpBaseBalloons() []Item {
	var result []Item
	for _, it := range Items {
		if tools.InListInt(it.ID, itemsBalloonsForCorpBase) {
			result = append(result, it)
		}
	}
	return result
}
