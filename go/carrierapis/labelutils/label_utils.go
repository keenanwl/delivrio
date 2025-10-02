package labelutils

import (
	"delivrio.io/go/ent/user"
	"time"
)

func PickupDayToTime(now time.Time, pickupDay user.PickupDay) time.Time {

	// So time is not in past by the time the request
	// makes it to the carrier API
	now = now.Add(time.Minute * 5)

	output := now
	switch pickupDay {
	case user.PickupDayToday:
		break
	case user.PickupDayTomorrow:
		output = output.AddDate(0, 0, 1)
		break
	case user.PickupDayIn_2_Days:
		output = output.AddDate(0, 0, 2)
		break
	case user.PickupDayIn_3_Days:
		output = output.AddDate(0, 0, 3)
		break
	case user.PickupDayIn_4_Days:
		output = output.AddDate(0, 0, 4)
		break
	case user.PickupDayIn_5_Days:
		output = output.AddDate(0, 0, 5)
		break
	}

	return output

}
