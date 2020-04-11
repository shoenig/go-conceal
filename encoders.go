package secrets

// MarshalText implements the encoding.TextMarshaler interface, which is also
// compatible with the encoding/json marshaler.
func (t Text) MarshalText() ([]byte, error) {
	return []byte(t.Secret()), nil
}

// UnmarshalText implments the encoding.TextUnmarshaler interface, which is
// also compatible with the encoding/json unmarhaler.
func (t *Text) UnmarshalText(text []byte) error {
	t.value = string(text)
	return nil
}
