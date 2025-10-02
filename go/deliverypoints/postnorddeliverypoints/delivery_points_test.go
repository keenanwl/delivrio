package postnorddeliverypoints

import (
	"delivrio.io/go/deliverypoints/postnorddeliverypoints/postnordresponse"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDayDistance(t *testing.T) {
	// Define test cases
	var testCases = []struct {
		name     string
		dayOpen  postnordresponse.Day
		dayClose postnordresponse.Day
		want     []postnordresponse.Day
	}{
		{
			name:     "MondayToFriday",
			dayOpen:  postnordresponse.Monday,
			dayClose: postnordresponse.Friday,
			want:     []postnordresponse.Day{postnordresponse.Tuesday, postnordresponse.Wednesday, postnordresponse.Thursday},
		},
		{
			name:     "FridayToMonday",
			dayOpen:  postnordresponse.Friday,
			dayClose: postnordresponse.Monday,
			want:     []postnordresponse.Day{postnordresponse.Saturday, postnordresponse.Sunday},
		},
		{
			name:     "MondayToMonday",
			dayOpen:  postnordresponse.Monday,
			dayClose: postnordresponse.Monday,
			want:     []postnordresponse.Day{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := dayDistance(tc.dayOpen, tc.dayClose)

			require.Equalf(t, tc.want, got, "%s: Incorrect days returned. Got %v, want %v", tc.name, got, tc.want)
		})
	}
}
