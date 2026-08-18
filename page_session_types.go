package freight

import (
	"encoding/base64"
	"encoding/json"
)

type ScopedCursor struct{ Tenant, Filter, Direction, LastID string }
type PageCheckpoint struct {
	Tenant, Filter, Direction, LastID string
	Version                           int
}
type PageSession struct{ Checkpoint PageCheckpoint }

func EncodeScopedCursor(v ScopedCursor) string {
	b, _ := json.Marshal(v)
	return base64.RawURLEncoding.EncodeToString(b)
}
