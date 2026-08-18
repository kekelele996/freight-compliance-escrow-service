package freight

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func CanonicalTenant(v string) string { return strings.ToLower(strings.TrimSpace(v)) }
func ShipmentReference(tenant, id string) string {
	sum := sha256.Sum256([]byte(strings.ToUpper(strings.TrimSpace(tenant)) + ":" + strings.TrimSpace(id)))
	return "SHP-" + strings.ToUpper(hex.EncodeToString(sum[:6]))
}
