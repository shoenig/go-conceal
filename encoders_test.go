package secrets

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestText_JSON_marshal(t *testing.T) {
	text := New("abc123")

	result, err := json.Marshal(text)
	require.NoError(t, err)
	require.Equal(t, []byte(`"abc123"`), result)
}

func TestText_JSON_unmarshal(t *testing.T) {
	bs, err := json.Marshal(`abc123`)
	require.NoError(t, err)

	var result Text
	err = json.Unmarshal(bs, &result)
	require.NoError(t, err)
	require.Equal(t, "abc123", result.Secret())
}
