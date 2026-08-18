package freight

import "time"

type SettlementRepository interface {
	Save(Settlement) error
	Load(string) (Settlement, error)
}
type MemorySettlementRepo struct {
	data     map[string]Settlement
	FailSave bool
}

func NewMemorySettlementRepo() *MemorySettlementRepo {
	return &MemorySettlementRepo{data: map[string]Settlement{}}
}
func (r *MemorySettlementRepo) Save(s Settlement) error {
	if r.FailSave {
		return ErrConflict
	}
	r.data[s.ID] = copySettlement(s)
	return nil
}
func (r *MemorySettlementRepo) Load(id string) (Settlement, error) {
	v, ok := r.data[id]
	if !ok {
		return Settlement{}, ErrNotFound
	}
	return copySettlement(v), nil
}
func copySettlement(s Settlement) Settlement {
	s.Lines = append([]ChargeLine(nil), s.Lines...)
	return s
}

type SettlementService struct {
	repo  SettlementRepository
	clock Clock
}

func (s SettlementService) Approve(id string, expected int) (Settlement, error) {
	v, e := s.repo.Load(id)
	if e != nil {
		return Settlement{}, e
	}
	if v.Version != expected {
		return Settlement{}, ErrConflict
	}
	if v.Status != "pending" {
		return Settlement{}, ErrValidation
	}
	v.Status = "approved"
	v.Version++
	v.UpdatedAt = s.clock.Now()
	if e = s.repo.Save(v); e != nil {
		return Settlement{}, e
	}
	return v, nil
}
func NewPendingSettlement(id, tenant, shipment string, total Money, now time.Time) Settlement {
	return Settlement{ID: id, TenantID: tenant, ShipmentID: shipment, Status: "pending", Total: total, Version: 1, UpdatedAt: now}
}
