package duty

import "time"

const (
	DUTY_CANTEEN = 1
	DUTY_GEO     = 2
	DUTY_PATROL  = 3

	DUTY_CANTEEN_PERIOD = time.Hour
	DUTY_GEO_PERIOD     = time.Hour * 3
	DUTY_PATROL_PERIOD  = time.Hour * 7

	//DUTY_CANTEEN_PERIOD = time.Second * 3
	//DUTY_GEO_PERIOD     = time.Second * 3
	//DUTY_PATROL_PERIOD  = time.Second * 3
)

func GetTimeForDuty(id int) time.Time {
	switch id {
	case DUTY_CANTEEN:
		return time.Now().Add(DUTY_CANTEEN_PERIOD)
	case DUTY_GEO:
		return time.Now().Add(DUTY_GEO_PERIOD)
	case DUTY_PATROL:
		return time.Now().Add(DUTY_PATROL_PERIOD)
	}
	return time.Now()
}
