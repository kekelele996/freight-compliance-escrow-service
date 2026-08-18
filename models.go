package freight

import "time"

type Money struct {
	Currency string
	Cents    int64
}
type Shipment struct {
	ID, TenantID        string
	DeclaredWeightGrams int64
	DeclaredValue       Money
	Segments            []Segment
	Tags                map[string]string
	CreatedAt           time.Time
}
type Segment struct {
	Carrier, From, To  string
	Departure, Arrival time.Time
	CapacityGrams      int64
}
type ChargeLine struct {
	Code     string
	Amount   Money
	Metadata map[string]string
}
type Settlement struct {
	ID, TenantID, ShipmentID, Status string
	Lines                            []ChargeLine
	Total                            Money
	Version                          int
	UpdatedAt                        time.Time
}
type ComplianceHold struct {
	ShipmentID, Reason string
	ExpiresAt          time.Time
}
type RouteEdge struct {
	From, To      string
	Transit       time.Duration
	Opens, Closes int
}
type TariffTier struct{ UpToGrams, RateCents int64 }
type Tariff struct {
	Currency          string
	Tiers             []TariffTier
	OversizeRateCents int64
}
type Manifest struct {
	ShipmentID string         `json:"shipment_id"`
	Stops      []ManifestStop `json:"stops"`
}
type ManifestStop struct {
	Code  string   `json:"code"`
	Seals []string `json:"seals"`
}
type OutboxEvent struct {
	ID, TenantID, Type, Payload string
	Sequence                    int
	Published                   bool
}
type PageCursor struct{ Tenant, LastID string }
type CapacityRequest struct {
	ID       string
	Weight   int64
	Priority int
}
type CarrierExposure struct {
	Carrier       string
	ShipmentCount int
	WeightGrams   int64
	Value         Money
}
type AuditEntry struct {
	At                        time.Time
	Tenant, Action, Reference string
}
type Config struct {
	DefaultCurrency   string
	WorkerParallelism int
	MaxWeightGrams    int64
}
