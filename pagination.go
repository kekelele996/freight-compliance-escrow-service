package freight

import (
	"encoding/base64"
	"fmt"
	"sort"
	"strings"
)

func EncodeCursor(c PageCursor) string {
	return base64.RawURLEncoding.EncodeToString([]byte(c.Tenant + "|" + c.LastID))
}
func DecodeCursor(raw string) (PageCursor, error) {
	b, e := base64.RawURLEncoding.DecodeString(raw)
	if e != nil {
		return PageCursor{}, ErrValidation
	}
	p := strings.Split(string(b), "|")
	if len(p) != 2 || p[0] == "" {
		return PageCursor{}, ErrValidation
	}
	return PageCursor{p[0], p[1]}, nil
}
func PaginateShipments(items []Shipment, tenant, cursor string, limit int) ([]Shipment, string, error) {
	if limit < 1 {
		return nil, "", ErrValidation
	}
	f := []Shipment{}
	for _, s := range items {
		if CanonicalTenant(s.TenantID) == CanonicalTenant(tenant) {
			f = append(f, s)
		}
	}
	sort.Slice(f, func(i, j int) bool { return f[i].ID < f[j].ID })
	start := 0
	if cursor != "" {
		c, e := DecodeCursor(cursor)
		if e != nil {
			return nil, "", e
		}
		if CanonicalTenant(c.Tenant) != CanonicalTenant(tenant) {
			return nil, "", fmt.Errorf("%w: cursor tenant", ErrValidation)
		}
		for start < len(f) && f[start].ID <= c.LastID {
			start++
		}
	}
	end := start + limit
	if end > len(f) {
		end = len(f)
	}
	page := append([]Shipment(nil), f[start:end]...)
	next := ""
	if end < len(f) {
		next = EncodeCursor(PageCursor{CanonicalTenant(tenant), f[end-1].ID})
	}
	return page, next, nil
}
