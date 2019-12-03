package craft

import (
	"tcs/models"
	"time"
)

type CraftResources struct {
	ID    int
	Count int
}

type Craft struct {
	ID        int
	Item      int
	Time      time.Duration
	Resources []CraftResources
	Credits   int
}

var receipts = make(map[models.Location][]*Craft, 10)

func init() {
	receipts[models.CorporationBase] = []*Craft{
		{
			ID:   1,
			Item: 25,
			Time: time.Minute * 1,
			Resources: []CraftResources{
				{
					ID:    8,
					Count: 20,
				},
				{
					ID:    10,
					Count: 5,
				},
			},
			Credits: 10,
		},
		{
			ID:   2,
			Item: 26,
			Time: time.Minute * 1,
			Resources: []CraftResources{
				{
					ID:    8,
					Count: 20,
				},
				{
					ID:    10,
					Count: 5,
				},
			},
			Credits: 10,
		},
	}
	receipts[models.PiratesBase] = receipts[models.CorporationBase]
}
