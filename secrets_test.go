package secrets

import (
	"fmt"
	"testing"

	"github.com/shoenig/test/must"
)

func TestText_New(t *testing.T) {
	t.Parallel()

	text := New("abc123")
	value := text.Secret()
	must.Eq(t, "abc123", value)
}

func TestText_String(t *testing.T) {
	t.Parallel()

	text := New("abc123")
	s := fmt.Sprintf("%s", text)
	must.Eq(t, redactString, s)
}

func TestText_GoString(t *testing.T) {
	t.Parallel()

	text := New("abc123")
	s := fmt.Sprintf("%#v", text)
	must.Eq(t, redactTextGoString, s)
}

func TestText_Equal(t *testing.T) {
	t.Parallel()

	t.Run("same", func(t *testing.T) {
		a := New("foo")
		b := New("foo")
		same := a.Equal(b)
		must.True(t, same)
	})

	t.Run("different", func(t *testing.T) {
		foo := New("foo")
		bar := New("bar")
		same := foo.Equal(bar)
		must.False(t, same)
	})
}

func TestBytes_NewBytes(t *testing.T) {
	t.Parallel()

	original := []byte{1, 2}
	bs := NewBytes(original)
	value1 := bs.Secret()
	must.Eq(t, []byte{1, 2}, value1)

	// modify original, bs must not change
	original[0] = 9
	value2 := bs.Secret()
	must.Eq(t, []byte{1, 2}, value2)
}

func TestBytes_String(t *testing.T) {
	t.Parallel()

	bs := NewBytes([]byte{1, 2, 3})
	s := fmt.Sprintf("%s", bs)
	must.Eq(t, redactString, s)
}

func TestBytes_GoString(t *testing.T) {
	t.Parallel()

	bs := NewBytes([]byte{1, 2, 3})
	s := fmt.Sprintf("%#v", bs)
	must.Eq(t, redactBytesGoString, s)
}

func TestBytes_Equal(t *testing.T) {
	t.Parallel()

	t.Run("same", func(t *testing.T) {
		a := NewBytes([]byte("foo"))
		b := NewBytes([]byte("foo"))
		same := a.Equal(b)
		must.True(t, same)
	})

	t.Run("different", func(t *testing.T) {
		foo := NewBytes([]byte("foo"))
		bar := NewBytes([]byte("bar"))
		same := foo.Equal(bar)
		must.False(t, same)
	})
}
