package freight

import (
	"fmt"
	"time"
)

type BookingService struct {
	Shipments   *ShipmentStore
	Idempotency *IdempotencyStore
	Holds       *HoldBook
	Clock       Clock
}

func (s BookingService) Create(tenant, key string, shipment Shipment) (string, error) {
	fp := ShipmentReference(tenant, shipment.ID)
	old, e := s.Idempotency.Reserve(tenant, key, fp)
	if e != nil {
		return "", e
	}
	if old != "" {
		return old, nil
	}
	shipment.TenantID = CanonicalTenant(tenant)
	if e = ValidateShipment(shipment, s.Clock.Now().Add(5*time.Minute)); e != nil {
		return "", e
	}
	if h, ok := s.Holds.Active(shipment.ID, s.Clock.Now()); ok {
		return "", fmt.Errorf("%w: %s", ErrWindowClosed, h.Reason)
	}
	s.Shipments.Put(shipment)
	response := ShipmentReference(tenant, shipment.ID)
	s.Idempotency.Commit(tenant, key, response)
	return response, nil
}
