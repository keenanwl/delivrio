package ratelookup

import (
	"delivrio.io/shared-utils/pulid"
	"slices"
	"strconv"
)

type LocationZip struct {
	ID               pulid.ID
	Zip              string
	Name             string
	AddressFormatted string
}

// Replace with an `ORDER BY AVG(zip - lookup_zip) ASC`?
func closestXLocations(fromZip string, locations []LocationZip, max int) []LocationZip {

	if max >= len(locations) {
		return locations
	}

	baseID := pulid.ID("***|***")

	locations = append(locations, LocationZip{
		Zip: fromZip,
		ID:  baseID,
	})

	slices.SortFunc(locations, func(i, j LocationZip) int {
		first, err1 := strconv.ParseInt(i.Zip, 10, 64)
		second, err2 := strconv.ParseInt(j.Zip, 10, 64)

		if err1 != nil || err2 != nil {
			if i.Zip < j.Zip {
				return -1
			} else if i.Zip > j.Zip {
				return 1
			}
			return 0
		}
		if first < second {
			return -1
		} else if first > second {
			return 1
		}
		return 0
	})

	found := make([]LocationZip, 0)

	for i, v := range locations {
		if v.ID == baseID {

			takeAbove := max / 2
			takeBelow := max / 2
			if max%2 != 0 {
				takeAbove = ((max - 1) / 2) + 1
				takeBelow = (max - 1) / 2
			}

			if len(locations[:i-1]) < takeBelow {
				takeAbove = takeAbove + (takeBelow - len(locations[:i-1]))
				takeBelow = len(locations[:i-1])
			}

			if len(locations[i+1:]) < takeAbove {
				takeBelow = takeBelow + (takeAbove - len(locations[i+1:]))
				takeAbove = len(locations[i+1:])
			}

			found = append(found, locations[i-takeBelow:i]...)
			found = append(found, locations[i+1:i+1+takeAbove]...)
			continue

		}
	}

	return found

}
