package models

import (
	"github.com/finalist736/gokit/logger"
	"math/rand"
	"tcs/machanics/daytime"
)

type (
	Mob struct {
		ID          int
		Title       string
		HP          int
		Damage      [2]int
		Armor       int
		Probability int
		Loot        []int
		LootCount   [2]int
	}
)

var locationsMobs map[daytime.TimePart]map[Location][]Mob

func init() {
	// mobs
	locationsMobs = make(map[daytime.TimePart]map[Location][]Mob, 2)
	locationsMobs[daytime.Day] = make(map[Location][]Mob, 10)
	locationsMobs[daytime.Night] = make(map[Location][]Mob, 10)

	locationsMobs[daytime.Day][LeftDuneSlope] = append(locationsMobs[daytime.Day][LeftDuneSlope],
		Mob{
			ID:          0,
			Probability: 200,
		},
		Mob{
			ID:          1,
			Title:       "🧿 Какая-то аномалия",
			HP:          5,
			Damage:      [2]int{2, 3},
			Armor:       1,
			Probability: 10,
			Loot:        []int{13, 14, 15, 16},
			LootCount:   [2]int{1, 3},
		},
		Mob{
			ID:          2,
			Title:       "🦟 Мухозоид",
			HP:          5,
			Damage:      [2]int{2, 4},
			Armor:       1,
			Probability: 5,
			Loot:        []int{17, 18, 19},
			LootCount:   [2]int{2, 3},
		},
		Mob{
			ID:          3,
			Title:       "🦛 Крупный",
			HP:          5,
			Damage:      [2]int{2, 4},
			Armor:       1,
			Probability: 2,
			Loot:        []int{15, 17, 18, 19},
			LootCount:   [2]int{2, 4},
		},
	)

	locationsMobs[daytime.Night][LeftDuneSlope] = append(locationsMobs[daytime.Night][LeftDuneSlope],
		Mob{
			ID:          0,
			Probability: 100,
		},
		Mob{
			ID:          1,
			Title:       "🧿 Серая аномалия",
			Damage:      [2]int{2, 5},
			HP:          7,
			Armor:       2,
			Probability: 7,
			Loot:        []int{13, 14, 15, 16},
			LootCount:   [2]int{3, 4},
		},
		Mob{
			ID:          2,
			Title:       "🦟 Ночной мухозоид",
			HP:          7,
			Damage:      [2]int{2, 5},
			Armor:       2,
			Probability: 3,
			Loot:        []int{17, 18, 19},
			LootCount:   [2]int{2, 3},
		},
		Mob{
			ID:          3,
			Title:       "🦛 Тёмный",
			Damage:      [2]int{2, 5},
			HP:          7,
			Armor:       2,
			Probability: 1,
			Loot:        []int{15, 17, 18, 19},
			LootCount:   [2]int{2, 4},
		},
	)

	locationsMobs[daytime.Day][RightDuneSlope] = append(locationsMobs[daytime.Day][RightDuneSlope],
		Mob{
			ID:          0,
			Probability: 200,
		},
		Mob{
			ID:          1,
			Title:       "🧿 Какая-то аномалия",
			HP:          5,
			Damage:      [2]int{2, 3},
			Armor:       1,
			Probability: 10,
			Loot:        []int{13, 14, 15, 16},
			LootCount:   [2]int{1, 3},
		},
		Mob{
			ID:          2,
			Title:       "🦟 Мухозоид",
			HP:          5,
			Damage:      [2]int{2, 3},
			Armor:       1,
			Probability: 5,
			Loot:        []int{17, 18, 19},
			LootCount:   [2]int{2, 3},
		},
		Mob{
			ID:          3,
			Title:       "🦛 Крупный",
			HP:          5,
			Damage:      [2]int{2, 3},
			Armor:       1,
			Probability: 2,
			Loot:        []int{15, 17, 18, 19},
			LootCount:   [2]int{2, 4},
		},
	)

	locationsMobs[daytime.Night][RightDuneSlope] = append(locationsMobs[daytime.Night][RightDuneSlope],
		Mob{
			ID:          0,
			Probability: 100,
		},
		Mob{
			ID:          1,
			Title:       "🧿 Серая аномалия",
			Damage:      [2]int{3, 5},
			HP:          7,
			Armor:       2,
			Probability: 7,
			Loot:        []int{13, 14, 15, 16},
			LootCount:   [2]int{3, 4},
		},
		Mob{
			ID:          2,
			Title:       "🦟 Ночной мухозоид",
			HP:          7,
			Damage:      [2]int{3, 5},
			Armor:       2,
			Probability: 3,
			Loot:        []int{17, 18, 19},
			LootCount:   [2]int{2, 3},
		},
		Mob{
			ID:          3,
			Title:       "🦛 Тёмный",
			Damage:      [2]int{3, 5},
			HP:          7,
			Armor:       2,
			Probability: 1,
			Loot:        []int{15, 17, 18, 19},
			LootCount:   [2]int{2, 4},
		},
	)

	//fmt.Printf("%+v\n", locationsMobs)
}

func GetCurrentMob(l Location) (*Mob, bool) {
	if l.IsBase() {
		return nil, false
	}
	mobs := locationsMobs[daytime.Get()][l]
	var totalWeight int
	for _, mob := range mobs {
		totalWeight += mob.Probability
	}
	if totalWeight == 0 {
		return nil, false
	}
	logger.StdOut().Debugf("mobs total weight: %d", totalWeight)
	prob := rand.Intn(totalWeight)
	logger.StdOut().Debugf("mobs prob: %d", prob)
	totalWeight = 0
	for _, mob := range mobs {
		totalWeight += mob.Probability
		if prob < totalWeight {
			if mob.ID == 0 {
				break
			}
			return &mob, true
		}
	}
	return nil, false
}
