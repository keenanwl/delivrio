package bayesian

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProbabilityBBWinsOverA(t *testing.T) {

	bt := NewBinaryTest()
	bt.Add(8500, 1500)
	bt.Add(8500, 1410)

	probs := bt.Probabilities()

	expectedA := 0.9665542433357666
	require.Equal(t, expectedA, probs[0])

	expectedB := 0.03344575666423344
	require.Equal(t, expectedB, probs[1])
}
