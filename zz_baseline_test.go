package freight

import (
	"testing"
	"time"
)

func TestBaselineBookingAndTariff(t *testing.T) {
	now := time.Date(2026, 8, 18, 10, 0, 0, 0, time.UTC)
	clock := FixedClock{now}
	svc := BookingService{NewShipmentStore(), NewIdempotencyStore(clock), NewHoldBook(), clock}
	s := Shipment{ID: "s1", DeclaredWeightGrams: 10, DeclaredValue: MustMoney("USD", 100), CreatedAt: now, Segments: []Segment{{Carrier: "acme", From: "A", To: "B", Departure: now, Arrival: now.Add(time.Hour)}}}
	if _, e := svc.Create("north", "x", s); e != nil {
		t.Fatal(e)
	}
	q, e := (Tariff{Currency: "USD", Tiers: []TariffTier{{10, 2}}, OversizeRateCents: 3}).Quote(12)
	if e != nil || q.Cents != 26 {
		t.Fatal(q, e)
	}
}
