// Package conceal provides types for protecting sensitive text in logs.
package conceal

import (
	"slices"
)

const (
	redactString        = "(redacted)"
	redactTextGoString  = "conceal.Text{}"
	redactBytesGoString = "conceal.Bytes{}"
)

func hash(b []byte) int {
	h := 0
	for _, v := range b {
		h = h*31 + int(v)
	}
	return h
}

// New returns a Text that keeps s hidden.
func New(s string) *Text {
	return &Text{
		value: s,
		hash:  hash([]byte(s)),
	}
}

// A Text contains a string which should not be exposed (e.g. in logs).
//
// This type overrides String and GoString such that the underling data does
// not get exposed (typically through log statements) accidentally.
//
// To get at the underlying data, use the Secret method.
//
// Unlike a Go string, Text is not 'comparable' (it can't be a map key
// or support ==). Use its Equal method to compare, and its Hash method
// for use in a hashicorp/go-set.HashMap[Text].
type Text struct {
	_     [0]func() // not comparable
	value string
	hash  int
}

// Unveil returns the underlying value.
//
// This method should never be called in a context where the value should not
// be exposed, for example in log lines.
func (t *Text) Unveil() string {
	return t.value
}

// String returns "(redacted)" instead of the underlying value.
func (t *Text) String() string {
	return redactString
}

// GoString returns "conceal.Text{}".
func (t *Text) GoString() string {
	return redactTextGoString
}

// Equal returns true if the underlying text of t and o are the same.
func (t *Text) Equal(o *Text) bool {
	if t == nil || o == nil {
		return t == o
	}
	return t.value == o.value
}

// Copy creates a deep copy of t.
func (t *Text) Copy() *Text {
	return &Text{
		value: t.value,
		hash:  t.hash,
	}
}

// Hash creates a deterministic hash from the content of t.
//
// Implements hashicorp/go-set/HashFunc[int].
func (t *Text) Hash() int {
	return t.hash
}

// NewBytes returns a Bytes that keeps b a secret.
//
// A copy of b is created, so that later changes to b have no effect on the
// protected value.
func NewBytes(b []byte) *Bytes {
	return &Bytes{
		value: slices.Clone(b),
		hash:  hash(b),
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
	hash  int
}

// Unveil returns the underlying value.
//
// This method should never be called in a context where the value should not
// be exposed, for example in log lines.
func (b *Bytes) Unveil() []byte {
	return slices.Clone(b.value)
}

// String returns "<redacted>" instead of the underlying value.
func (b *Bytes) String() string {
	return redactString
}

// GoString returns "conceal.Bytes{}".
func (b *Bytes) GoString() string {
	return redactBytesGoString
}

// Equal returns true if the underlying bytes of t and o are the same.
func (b *Bytes) Equal(o *Bytes) bool {
	if b == nil || o == nil {
		return b == o
	}
	return slices.Equal(b.value, o.value)
}

// Copy creates a deep copy of b.
func (b *Bytes) Copy() *Bytes {
	return &Bytes{
		value: slices.Clone(b.value),
		hash:  b.hash,
	}
}

// Hash creates a deterministic hash from the content of t.
//
// Implements hashicorp/go-set/HashFunc[int].
func (b *Bytes) Hash() int {
	return b.hash
}
