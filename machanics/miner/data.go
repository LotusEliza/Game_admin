package miner

import (
	"fmt"
	"github.com/finalist736/gokit/logger"
	"math"
	"math/rand"
	"sync"
	"tcs/items"
	"tcs/models"
	"tcs/tools"
	"time"
)

const (
	GenerateTimeout = time.Minute * 120
)

type (
	MineSpot struct {
		Coordinates [2]int
		Type        items.ItemType
		Items       []int
	}
	SpotConfig struct {
		Count int
		Type  items.ItemType
		Items []int
	}
)

var (
	mux          sync.RWMutex
	spots        map[models.Location][]MineSpot
	defaultPoint = &MineSpot{
		Coordinates: [2]int{0, 0},
		Type:        "",
		Items:       nil,
	}

	locationSpotsConfig map[models.Location][]SpotConfig
)

func Init() {
	spots = make(map[models.Location][]MineSpot, 2)

	locationSpotsConfig = make(map[models.Location][]SpotConfig, 3)
	locationSpotsConfig[models.LeftDuneSlope] = append(
		locationSpotsConfig[models.LeftDuneSlope],
		SpotConfig{
			Count: 2,
			Type:  items.ITEM_METAL,
			Items: []int{8},
		},
		SpotConfig{
			Count: 2,
			Type:  items.ITEM_GAS,
			Items: []int{10},
		},
	)

	locationSpotsConfig[models.RightDuneSlope] = append(
		locationSpotsConfig[models.RightDuneSlope],
		SpotConfig{
			Count: 2,
			Type:  items.ITEM_METAL,
			Items: []int{8},
		},
		SpotConfig{
			Count: 2,
			Type:  items.ITEM_GAS,
			Items: []int{10},
		},
	)
	logger.StdOut().Debugf("spots config: %+v", locationSpotsConfig)
	go landFillingRoutine()
}

func landFillingRoutine() {
	generateSpots()
	for {
		<-time.After(GenerateTimeout)
		generateSpots()
	}
}

func generateSpots() {
	mux.Lock()
	defer mux.Unlock()

	var (
		x, y            int
		ox, oy          int
		dx, dy          int
		positionCorrect bool
	)

	for loc, cfg := range locationSpotsConfig {
		spots[loc] = make([]MineSpot, 0, len(cfg))
		for _, res := range cfg {
			for i := 0; i < res.Count; i++ {

				for {
					x = tools.RandMinMax(3, models.LocationsSizes[loc].Width-2)
					y = tools.RandMinMax(3, models.LocationsSizes[loc].Height-2)
					positionCorrect = true

					for _, it := range spots[loc] {
						ox = it.Coordinates[0]
						oy = it.Coordinates[1]
						dx = int(math.Abs(float64(ox - x)))
						dy = int(math.Abs(float64(oy - y)))
						if dx < 5 && dy < 5 {
							positionCorrect = false
							break
						}
					}
					if positionCorrect {
						break
					}
				}

				spot := MineSpot{}
				spot.Coordinates[0] = x
				spot.Coordinates[1] = y
				spot.Type = res.Type
				spot.Items = make([]int, len(res.Items))
				copy(spot.Items, res.Items)

				spots[loc] = append(spots[loc], spot)
			}
		}
	}

	logger.StdOut().Debugf("mine spots: %+v", spots)

}

func Get(l models.Location, t items.ItemType, maxCount, x, y int) (cnt, item int) {
	mux.RLock()
	defer mux.RUnlock()
	mspts, ok := spots[l]
	if !ok {
		return
	}
	if t != items.ITEM_METAL && t != items.ITEM_GAS && t != items.ITEM_JUNK {
		return
	}
	if t == items.ITEM_JUNK {
		maxCount -= int(math.Round(float64(maxCount) * 0.3))
	}

	for _, spt := range mspts {
		if t != items.ITEM_JUNK && (t != spt.Type) {
			continue
		}
		if (x < spt.Coordinates[0]-2 || x > spt.Coordinates[0]+2) ||
			(y < spt.Coordinates[1]-2 || y > spt.Coordinates[1]+2) {
			continue
		}
		return maxCount, spt.Items[rand.Intn(len(spt.Items))]
	}
	return
}

func GetAllSpots(l models.Location) string {
	mux.RLock()
	defer mux.RUnlock()
	l = models.NearLocation(l)
	mspts, ok := spots[l]
	if !ok {
		return "Нет мест для майнинга"
	}
	txt := "Рыбные места:\n"
	for _, spt := range mspts {

		txt += fmt.Sprintf(
			"%s майнится в координатах [%d:%d]\n",
			spt.Type.Name(),
			spt.Coordinates[0],
			spt.Coordinates[1])
	}
	txt += "\nЭти места меняются каждые два часа"
	return txt
}

func GetNearestSpot(l models.Location, x, y int) (bool, *MineSpot) {
	mux.RLock()
	defer mux.RUnlock()
	mspts, ok := spots[l]
	if !ok {
		return false, defaultPoint
	}
	var (
		dist float64
		spot MineSpot
		min  = math.MaxFloat64
	)

	for _, it := range mspts {
		dist = tools.Distance(x, y, it.Coordinates[0], it.Coordinates[1])
		if min > dist {
			min = dist
			spot = it
		}
	}
	return true, &spot
}
