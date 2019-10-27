// Package secrets provides types for protecting secretive text in logs.
package secrets

const (
	redactString        = "<redacted>"
	redactTextGoString  = "secrets.Text{}"
	redactBytesGoString = "secrets.Bytes{}"
)

// New returns a Text that keeps s secret.
func New(s string) Text {
	return Text{
		value: s,
	}
}

// A Text secret contains a string which should not be exposed.
//
// This type overrides String and GoString such that the underling data does
// not get exposed (typically through log statements) accidentally.
//
// To get at the underlying data, use the Secret method.
type Text struct {
	value string
}

// Secret returns the underlying value.
//
// This method should never be called in a context where the value should not
// be exposed, for example in log lines.
func (t Text) Secret() string {
	return t.value
}

// String returns "<redacted>" instead of the underlying value.
func (t Text) String() string {
	return redactString
}

// GoString returns "secrets.Text{}".
func (t Text) GoString() string {
	return redactTextGoString
}

// NewBytes returns a Bytes that keeps b a secret.
//
// A copy of b is created, so that later changes to b have no effect on the
// protected value.
func NewBytes(b []byte) Bytes {
	v := make([]byte, len(b))
	copy(v, b)
	return Bytes{
		value: v,
	}
}

// A Bytes secret contains a byte slice which should not be exposed.
//
// This type overrides String and GoString such that the underlying data does
// not get exposed (typically through log statements) accidentally.
//
// To get at the underlying data, use the Secret method.
type Bytes struct {
	value []byte
}

// Secret returns the underlying value.
//
// This method should never be called in a context where the value should not
// be exposed, for example in log lines.
func (t Bytes) Secret() []byte {
	b := make([]byte, len(t.value))
	copy(b, t.value)
	return b
}

// String returns "<redacted>" instead of the underlying value.
func (t Bytes) String() string {
	return redactString
}

// GoString returns "secrets.Bytes{}".
func (t Bytes) GoString() string {
	return redactBytesGoString
}
