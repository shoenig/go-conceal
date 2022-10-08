package conceal

import (
	"fmt"
	"testing"

	"github.com/shoenig/test/must"
)

func TestText_New(t *testing.T) {
	t.Parallel()

	text := New("abc123")
	value := text.Unveil()
	must.EqOp(t, "abc123", value)
}

func TestText_String(t *testing.T) {
	t.Parallel()

	text := New("abc123")
	s := fmt.Sprintf("%s", text)
	must.EqOp(t, redactString, s)
}

func TestText_GoString(t *testing.T) {
	t.Parallel()

	text := New("abc123")
	s := fmt.Sprintf("%#v", text)
	must.EqOp(t, redactTextGoString, s)
}

func TestText_Equal(t *testing.T) {
	t.Parallel()

	t.Run("same", func(t *testing.T) {
		a := New("foo")
		b := New("foo")
		must.Equal(t, a, b)
	})

	t.Run("different", func(t *testing.T) {
		a := New("foo")
		b := New("bar")
		must.NotEqual(t, a, b)
	})
}

func TestText_Copy(t *testing.T) {
	t.Parallel()

	orig := New("hello")
	c := orig.Copy()
	must.False(t, orig == c) // different pointers
}

func TestText_Hash(t *testing.T) {
	t.Parallel()

	a := New("hello")
	result := a.Hash()
	must.Eq(t, 99162322, result)
}

func TestBytes_NewBytes(t *testing.T) {
	t.Parallel()

	original := []byte{1, 2}
	bs := NewBytes(original)
	value1 := bs.Unveil()
	must.Eq(t, []byte{1, 2}, value1)

	// modify original, bs must not change
	original[0] = 9
	value2 := bs.Unveil()
	must.Eq(t, []byte{1, 2}, value2)
}

func TestBytes_String(t *testing.T) {
	t.Parallel()

	bs := NewBytes([]byte{1, 2, 3})
	s := fmt.Sprintf("%s", bs)
	must.EqOp(t, redactString, s)
}

func TestBytes_GoString(t *testing.T) {
	t.Parallel()

	bs := NewBytes([]byte{1, 2, 3})
	s := fmt.Sprintf("%#v", bs)
	must.EqOp(t, redactBytesGoString, s)
}

func TestBytes_Equal(t *testing.T) {
	t.Parallel()

	t.Run("same", func(t *testing.T) {
		a := NewBytes([]byte("foo"))
		b := NewBytes([]byte("foo"))
		must.Equal(t, a, b)
	})

	t.Run("different", func(t *testing.T) {
		a := NewBytes([]byte("foo"))
		b := NewBytes([]byte("bar"))
		must.NotEqual(t, a, b)
	})
}

func TestBytes_Copy(t *testing.T) {
	t.Parallel()

	orig := NewBytes([]byte("hello"))
	c := orig.Copy()
	must.Equal(t, orig, c)
}

func TestBytes_Hash(t *testing.T) {
	t.Parallel()

	a := NewBytes([]byte("hello"))
	result := a.Hash()
	must.Eq(t, 99162322, result)
}
