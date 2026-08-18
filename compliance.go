package freight

import "time"

type HoldBook struct{ holds map[string]ComplianceHold }

func NewHoldBook() *HoldBook               { return &HoldBook{holds: map[string]ComplianceHold{}} }
func (h *HoldBook) Place(v ComplianceHold) { h.holds[v.ShipmentID] = v }
func (h *HoldBook) Active(id string, now time.Time) (ComplianceHold, bool) {
	v, ok := h.holds[id]
	if !ok || !v.ExpiresAt.After(now) {
		return ComplianceHold{}, false
	}
	return v, true
}
func (h *HoldBook) ReleaseExpired(now time.Time) {
	for id, v := range h.holds {
		if !v.ExpiresAt.After(now) {
			delete(h.holds, id)
		}
	}
}
