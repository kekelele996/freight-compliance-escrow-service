package freight

import (
	"fmt"
	"strings"
	"time"
)

func ValidateShipment(s Shipment, cutoff time.Time) error {
	if CanonicalTenant(s.TenantID) == "" || s.ID == "" {
		return fmt.Errorf("%w: identity", ErrValidation)
	}
	if s.DeclaredWeightGrams <= 0 {
		return fmt.Errorf("%w: weight", ErrValidation)
	}
	if s.CreatedAt.After(cutoff) {
		return fmt.Errorf("%w: cutoff", ErrValidation)
	}
	if len(s.Segments) == 0 {
		return fmt.Errorf("%w: segments", ErrValidation)
	}
	for i, x := range s.Segments {
		if strings.TrimSpace(x.From) == "" || x.From == x.To || !x.Arrival.After(x.Departure) {
			return fmt.Errorf("%w: segment %d", ErrValidation, i)
		}
		if i > 0 && s.Segments[i-1].To != x.From {
			return fmt.Errorf("%w: disconnected", ErrValidation)
		}
	}
	return nil
}
func SameOperationalDay(a, b time.Time, loc *time.Location) bool {
	if loc == nil {
		loc = time.UTC
	}
	x, y := a.In(loc), b.In(loc)
	return x.Year() == y.Year() && x.YearDay() == y.YearDay()
}
