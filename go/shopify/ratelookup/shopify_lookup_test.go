package ratelookup

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClosestXLocations(t *testing.T) {

	res := closestXLocations("8000", []LocationZip{
		{ID: "L1", Zip: "1000"},
		{ID: "L3", Zip: "3000"},
		{ID: "L10", Zip: "10000"},
		{ID: "L11", Zip: "11000"},
		{ID: "L12", Zip: "12000"},
		{ID: "L4", Zip: "4000"},
		{ID: "L5", Zip: "5000"},
		{ID: "L6", Zip: "6000"},
		{ID: "L7", Zip: "7000"},
		{ID: "L9", Zip: "9000"},
		{ID: "L2", Zip: "2000"},
		{ID: "L8", Zip: "8000"},
	}, 3)

	if res[0].ID != "L7" || res[1].ID != "L8" || res[2].ID != "L9" {
		t.Fatal(res)
	}

	res = closestXLocations("A8000", []LocationZip{
		{ID: "L1", Zip: "1000X"},
		{ID: "L3", Zip: "3000Z"},
		{ID: "L10", Zip: "10000A"},
		{ID: "L11", Zip: "11000P"},
		{ID: "L12", Zip: "12000H"},
		{ID: "L4", Zip: "4000Q"},
		{ID: "L5", Zip: "A5000"},
		{ID: "L6", Zip: "EC1A 1BB"},
		{ID: "L7", Zip: "W1A 0AX"},
		{ID: "L9", Zip: "SW1A 1AA"},
		{ID: "L2", Zip: "SE1 7PB"},
		{ID: "L8", Zip: "NW1W 0NY"},
	}, 3)

	if res[0].ID != "L5" || res[1].ID != "L6" || res[2].ID != "L8" {
		t.Fatal(res)
	}

}

func Test_fastHashIt(t *testing.T) {
	h, err := fastHashIt([]byte("{some long string}"))
	require.NoError(t, err)
	require.Equal(t, "510491756", h)
}
