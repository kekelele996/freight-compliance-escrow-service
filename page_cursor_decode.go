package freight

import (
	"encoding/base64"
	"encoding/json"
)

func decodeScopedCursor(raw string) (ScopedCursor, error) {
	b, e := base64.RawURLEncoding.DecodeString(raw)
	if e != nil {
		return ScopedCursor{}, ErrValidation
	}
	var v ScopedCursor
	if e = json.Unmarshal(b, &v); e != nil {
		return ScopedCursor{}, ErrValidation
	}
	return v, nil
}
