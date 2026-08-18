package freight

import "sort"

func (t Tariff) Quote(weight int64) (Money, error) {
	if weight < 0 {
		return Money{}, ErrInvalidAmount
	}
	tiers := append([]TariffTier(nil), t.Tiers...)
	sort.Slice(tiers, func(i, j int) bool { return tiers[i].UpToGrams < tiers[j].UpToGrams })
	var total, used int64
	for _, tier := range tiers {
		if weight <= used {
			break
		}
		take := weight - used
		span := tier.UpToGrams - used
		if take > span {
			take = span
		}
		total += take * tier.RateCents
		used += take
	}
	if weight > used {
		total += (weight - used) * t.OversizeRateCents
	}
	return NewMoney(t.Currency, total)
}
func (t Tariff) ApplyContractDiscount(base Money, bp int64, minimum Money) (Money, error) {
	if base.Currency != minimum.Currency {
		return Money{}, ErrCurrencyMismatch
	}
	d, e := base.MultiplyBasisPoints(bp)
	if e != nil {
		return Money{}, e
	}
	after, e := base.Sub(d)
	if e != nil {
		return Money{}, e
	}
	if after.Cents < minimum.Cents {
		return minimum, nil
	}
	return after, nil
}
