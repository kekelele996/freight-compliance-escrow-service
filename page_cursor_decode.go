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
	if e = json.Unmarshal(b, &v); e != nil {
		return ScopedCursor{}, e
	}
	v.Tenant = CanonicalTenant(v.Tenant)
	return v, nil
}
