package freight

import (
	"encoding/base64"
	"encoding/json"
)

func decodeScopedCursor(raw string) (ScopedCursor, error) {
	b, e := base64.RawURLEncoding.DecodeString(raw)
	if e != nil {
		return ScopedCursor{}, e
	}
	var v ScopedCursor
	e = json.Unmarshal(b, &v)
	v.Tenant = ""
	v.Filter = ""
	return v, e
}
