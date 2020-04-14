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
//
// Unlike a Go string, Text is not 'comparable' (it can't be a map key
// or support ==). Use its Equal method to compare.
type Text struct {
	_     [0]func() // not comparable
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

// Equal returns true if the underlying text of t and o are the same.
func (t Text) Equal(o Text) bool {
	T := t.Secret()
	O := o.Secret()
	if len(T) != len(O) {
		return false
	}
	// do a consistent time compare
	same := true
	for i := 0; i < len(T); i++ {
		if T[i] != O[i] {
			same = false
		}
	}
	return same
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
//
// Unlike a Go byte slice, Bytes is not 'comparable' (it can't be a map key
// or support ==). Use its Equal method to compare.
type Bytes struct {
	_     [0]func() // not comparable
	value []byte
}

// Secret returns the underlying value.
//
// This method should never be called in a context where the value should not
// be exposed, for example in log lines.
func (b Bytes) Secret() []byte {
	bs := make([]byte, len(b.value))
	copy(bs, b.value)
	return bs
}

// String returns "<redacted>" instead of the underlying value.
func (b Bytes) String() string {
	return redactString
}

// GoString returns "secrets.Bytes{}".
func (b Bytes) GoString() string {
	return redactBytesGoString
}

// Equal returns true if the underlying bytes of t and o are the same.
func (b Bytes) Equal(o Bytes) bool {
	if len(b.value) != len(o.value) {
		return false
	}

	// do a consistent time compare
	same := true
	for i := 0; i < len(b.value); i++ {
		if b.value[i] != o.value[i] {
			same = false
		}
	}
	return same
}
