package conceal

import (
	"fmt"
	"testing"

	"github.com/shoenig/test/must"
)

func TestUUIDv4(t *testing.T) {
	id := UUIDv4()

	s := fmt.Sprintf("%s", id)
	must.Eq(t, "(redacted)", s)

	real := id.Unveil()
	must.UUIDv4(t, real)
}
