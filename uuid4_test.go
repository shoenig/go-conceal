package conceal

import (
	"fmt"
	"testing"

	"github.com/shoenig/test/must"
)

func TestUUIDv4(t *testing.T) {
	id := UUIDv4()

	s := fmt.Sprintf("%s", id) // nolint: staticcheck
	must.Eq(t, "(redacted)", s)

	result := id.Unveil()
	must.UUIDv4(t, result)
}
