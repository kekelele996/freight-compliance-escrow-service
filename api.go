package freight

import "fmt"

type CreateShipmentRequest struct {
	Tenant, IdempotencyKey string
	Shipment               Shipment
}
type API struct{ Booking BookingService }

func (a API) CreateShipment(r CreateShipmentRequest) (string, error) {
	if r.IdempotencyKey == "" {
		return "", fmt.Errorf("%w: idempotency key", ErrValidation)
	}
	return a.Booking.Create(r.Tenant, r.IdempotencyKey, r.Shipment)
}
