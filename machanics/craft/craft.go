package craft

import (
	"fmt"
	"tcs/items"
	"tcs/models"
)

func Get(id int) *Craft {
	for _, recs := range receipts {
		for _, rec := range recs {
			if rec.ID == id {
				return rec
			}
		}
	}
	return nil
}

func GetLocationCraftList(l models.Location) ([]*Craft, bool) {
	list, ok := receipts[l]
	return list, ok
}

func (s Craft) Title() string {
	var (
		item         *items.Item
		resourceItem *items.Item
		res          string
	)
	item = items.GetByID(s.Item)
	if item == nil {
		return fmt.Sprintf("%d unknown!", s.Item)
	}
	for _, it := range s.Resources {
		resourceItem = items.GetByID(it.ID)
		res += fmt.Sprintf("    %s - %d\n", resourceItem.InventoryItemTitle(), it.Count)
	}
	txt := fmt.Sprintf("%s /craft_%d\n  Стоимость: 💰%d\n  Время печати: ⏱%d м.\n  Ресурсы:\n%s", item.SimpleItemTitle(), s.ID, s.Credits, int(s.Time.Minutes()), res)
	return txt
}

func IsReceiptInLocation(id int, l models.Location) bool {
	for _, it := range receipts[l] {
		if it.ID == id {
			return true
		}
	}
	return false
}
