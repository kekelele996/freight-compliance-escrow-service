package freight

import "sort"

func AllocateCapacity(capacity int64, requests []CapacityRequest) []string {
	r := append([]CapacityRequest(nil), requests...)
	sort.SliceStable(r, func(i, j int) bool {
		if r[i].Priority == r[j].Priority {
			return r[i].ID < r[j].ID
		}
		return r[i].Priority > r[j].Priority
	})
	out := []string{}
	for _, x := range r {
		if x.Weight > 0 && x.Weight <= capacity {
			capacity -= x.Weight
			out = append(out, x.ID)
		}
	}
	return out
}
