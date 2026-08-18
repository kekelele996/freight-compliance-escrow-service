package freight

import "sort"

func SummarizeCarrierExposure(items []Shipment, currency string) []CarrierExposure {
	m := map[string]CarrierExposure{}
	for _, s := range items {
		for _, x := range s.Segments {
			v := m[x.Carrier]
			v.Carrier = x.Carrier
			v.ShipmentCount++
			v.WeightGrams += s.DeclaredWeightGrams
			if s.DeclaredValue.Currency == currency {
				v.Value.Currency = currency
				v.Value.Cents += s.DeclaredValue.Cents
			}
			m[x.Carrier] = v
		}
	}
	o := make([]CarrierExposure, 0, len(m))
	for _, v := range m {
		o = append(o, v)
	}
	sort.Slice(o, func(i, j int) bool { return o[i].Carrier < o[j].Carrier })
	return o
}
