package conceal

// MarshalText implements the encoding.TextMarshaller interface, which is also
// compatible with the encoding/json marshaller.
func (t *Text) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface, which is
// also compatible with the encoding/json un-marshaller.
func (t *Text) UnmarshalText(text []byte) error {
	t.value = string(text)
	return nil
}
