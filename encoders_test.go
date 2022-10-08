package conceal

import (
	"encoding/json"
	"testing"

	"github.com/shoenig/test/must"
)

func TestText_JSON_marshal(t *testing.T) {
	text := New("abc123")

	result, err := json.Marshal(text)
	must.NoError(t, err)
	must.Eq(t, []byte(`"abc123"`), result)
}

func TestText_JSON_unmarshal(t *testing.T) {
	bs, err := json.Marshal(`abc123`)
	must.NoError(t, err)

	var result Text
	err = json.Unmarshal(bs, &result)
	must.NoError(t, err)
	must.Eq(t, "abc123", result.Unveil())
}
