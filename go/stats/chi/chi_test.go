package chi

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCalculateChiTestStatistic(t *testing.T) {

	ts := TestStatistic(1, 4, 1, 1)
	expected := 0.5528571428571429
	require.Equal(t, expected, ts)

	ts = TestStatistic(35, 15, 26, 29)
	// They use heavy + weird rounding
	// https://datascienceplus.com/chi-squared-test-in-r/
	expected = 5.71064864097651
	require.Equal(t, expected, ts)

}
