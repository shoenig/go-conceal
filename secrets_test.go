package secrets

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Text(t *testing.T) {
	text := New("abc123")
	value := text.Secret()
	require.Equal(t, "abc123", value)
}

func Test_Text_String(t *testing.T) {
	text := New("abc123")
	s := fmt.Sprintf("%s", text)
	require.Equal(t, redactString, s)
}

func Test_Text_GoString(t *testing.T) {
	text := New("abc123")
	s := fmt.Sprintf("%#v", text)
	require.Equal(t, redactTextGoString, s)
}

func Test_Bytes(t *testing.T) {
	original := []byte{1, 2}
	bs := NewBytes(original)

	value1 := bs.Secret()
	require.Equal(t, []byte{1, 2}, value1)

	// modify original, bs must not change
	original[0] = 9
	value2 := bs.Secret()
	require.Equal(t, []byte{1, 2}, value2)
}

func Test_Bytes_String(t *testing.T) {
	bs := NewBytes([]byte{1, 2, 3})
	s := fmt.Sprintf("%s", bs)
	require.Equal(t, redactString, s)
}

func Test_Bytes_GoString(t *testing.T) {
	bs := NewBytes([]byte{1, 2, 3})
	s := fmt.Sprintf("%#v", bs)
	require.Equal(t, redactBytesGoString, s)
}
