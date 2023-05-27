package conceal

import (
	"crypto/rand"
	"fmt"
)

// UUIDv4 creates a quasi uuid v4 formatted string, useable for secrets.
//
// Not strictly compliant with uuid v4 as these do not contain version bits.
func UUIDv4() *Text {
	b := make([]byte, 16)
	rand.Read(b)

	return New(
		fmt.Sprintf("%x-%x-%x-%x-%x",
			b[0:4],
			b[4:6],
			b[6:8],
			b[8:10],
			b[10:16],
		),
	)
}
