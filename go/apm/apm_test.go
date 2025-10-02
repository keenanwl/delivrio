package apm

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// Just trying to prevent foot-guns
// https://github.com/open-telemetry/opentelemetry-go/issues/4476
func Test_NewResource(t *testing.T) {
	_, err := NewResource("123")
	require.NoError(t, err)
}
