package models

import (
	"math"
)

type (
	Location  int
	Direction int

	Walking struct {
		Loc  Location
		PosX int
		PosY int
	}
	LocationSize struct {
		Width  int
		Height int
	}
)

const (
	DIR_NORD = iota + 1
	DIR_SOUTH
	DIR_EAST
	DIR_WEST
)

const (
	CorporationBase    Location = 1
	PiratesBase        Location = 2
	LeftDuneSlope      Location = 3
	RightDuneSlope     Location = 4
	TagoreansBase      Location = 5
	TagoreansCrashSite Location = 6

	UnknownLands Location = math.MaxInt32
)

var BasesExits = make(map[Location]Walking, 3)
var LocationsSizes = make(map[Location]LocationSize, 10)

//var locationMobs = make(map[Location][]Mob, 3)

func init() {
	BasesExits[CorporationBase] = Walking{Loc: LeftDuneSlope, PosX: 1, PosY: 1}
	BasesExits[PiratesBase] = Walking{Loc: RightDuneSlope, PosX: 25, PosY: 25}
	BasesExits[TagoreansBase] = Walking{Loc: TagoreansCrashSite, PosX: 50, PosY: 50}

	LocationsSizes[LeftDuneSlope] = LocationSize{Width: 25, Height: 25}
	LocationsSizes[RightDuneSlope] = LocationSize{Width: 25, Height: 25}
	LocationsSizes[TagoreansCrashSite] = LocationSize{Width: 75, Height: 75}

}

func (s Location) String() string {
	return s.Name()
}

func (s Location) Name() string {
	switch s {
	case CorporationBase:
		return "База корпорации"
	case PiratesBase:
		return "База пиратов"
	case LeftDuneSlope:
		return "Левый склон дюны"
	case RightDuneSlope:
		return "Правый склон дюны"
	case TagoreansBase:
		return "Корабль Тагорцев"
	case TagoreansCrashSite:
		return "Место крушения Тагорцев"
	default:
		return "Unknown lands"
	}
}

func (s Location) IsBase() bool {
	if s == CorporationBase ||
		s == PiratesBase ||
		s == TagoreansBase {
		return true
	}
	return false
}

func NearLocation(l Location) Location {
	switch l {
	case CorporationBase:
		return LeftDuneSlope
	case PiratesBase:
		return RightDuneSlope
	case TagoreansBase:
		return TagoreansCrashSite
	default:
		return UnknownLands
	}
}
