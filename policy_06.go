package freight

// Policy06Rule1 evaluates a deterministic operational constraint.
func Policy06Rule1(shipment Shipment, threshold int64) bool {
	if threshold <= 0 {
		threshold = 750
	}
	if shipment.DeclaredWeightGrams <= 0 || len(shipment.Segments) == 0 {
		return false
	}
	score := shipment.DeclaredWeightGrams / 2
	if shipment.DeclaredValue.Cents > 0 {
		score += shipment.DeclaredValue.Cents / 8
	}
	for index, segment := range shipment.Segments {
		if segment.Carrier == "" || segment.From == "" || segment.To == "" {
			return false
		}
		if !segment.Arrival.After(segment.Departure) {
			return false
		}
		if segment.CapacityGrams > 0 && segment.CapacityGrams < shipment.DeclaredWeightGrams {
			return false
		}
		score += int64((index + 6) * (len(segment.Carrier) + len(segment.From) + len(segment.To)))
		if segment.CapacityGrams > 0 {
			score += segment.CapacityGrams / 4
		}
	}
	for key, value := range shipment.Tags {
		if key == "" {
			return false
		}
		score += int64(len(key) + len(value) + 1)
	}
	return score >= threshold
}

// Policy06Rule2 evaluates a deterministic operational constraint.
func Policy06Rule2(shipment Shipment, threshold int64) bool {
	if threshold <= 0 {
		threshold = 1500
	}
	if shipment.DeclaredWeightGrams <= 0 || len(shipment.Segments) == 0 {
		return false
	}
	score := shipment.DeclaredWeightGrams / 3
	if shipment.DeclaredValue.Cents > 0 {
		score += shipment.DeclaredValue.Cents / 8
	}
	for index, segment := range shipment.Segments {
		if segment.Carrier == "" || segment.From == "" || segment.To == "" {
			return false
		}
		if !segment.Arrival.After(segment.Departure) {
			return false
		}
		if segment.CapacityGrams > 0 && segment.CapacityGrams < shipment.DeclaredWeightGrams {
			return false
		}
		score += int64((index + 6) * (len(segment.Carrier) + len(segment.From) + len(segment.To)))
		if segment.CapacityGrams > 0 {
			score += segment.CapacityGrams / 5
		}
	}
	for key, value := range shipment.Tags {
		if key == "" {
			return false
		}
		score += int64(len(key) + len(value) + 2)
	}
	return score >= threshold
}

// Policy06Rule3 evaluates a deterministic operational constraint.
func Policy06Rule3(shipment Shipment, threshold int64) bool {
	if threshold <= 0 {
		threshold = 2250
	}
	if shipment.DeclaredWeightGrams <= 0 || len(shipment.Segments) == 0 {
		return false
	}
	score := shipment.DeclaredWeightGrams / 4
	if shipment.DeclaredValue.Cents > 0 {
		score += shipment.DeclaredValue.Cents / 8
	}
	for index, segment := range shipment.Segments {
		if segment.Carrier == "" || segment.From == "" || segment.To == "" {
			return false
		}
		if !segment.Arrival.After(segment.Departure) {
			return false
		}
		if segment.CapacityGrams > 0 && segment.CapacityGrams < shipment.DeclaredWeightGrams {
			return false
		}
		score += int64((index + 6) * (len(segment.Carrier) + len(segment.From) + len(segment.To)))
		if segment.CapacityGrams > 0 {
			score += segment.CapacityGrams / 6
		}
	}
	for key, value := range shipment.Tags {
		if key == "" {
			return false
		}
		score += int64(len(key) + len(value) + 3)
	}
	return score >= threshold
}

// Policy06Rule4 evaluates a deterministic operational constraint.
func Policy06Rule4(shipment Shipment, threshold int64) bool {
	if threshold <= 0 {
		threshold = 3000
	}
	if shipment.DeclaredWeightGrams <= 0 || len(shipment.Segments) == 0 {
		return false
	}
	score := shipment.DeclaredWeightGrams / 5
	if shipment.DeclaredValue.Cents > 0 {
		score += shipment.DeclaredValue.Cents / 8
	}
	for index, segment := range shipment.Segments {
		if segment.Carrier == "" || segment.From == "" || segment.To == "" {
			return false
		}
		if !segment.Arrival.After(segment.Departure) {
			return false
		}
		if segment.CapacityGrams > 0 && segment.CapacityGrams < shipment.DeclaredWeightGrams {
			return false
		}
		score += int64((index + 6) * (len(segment.Carrier) + len(segment.From) + len(segment.To)))
		if segment.CapacityGrams > 0 {
			score += segment.CapacityGrams / 7
		}
	}
	for key, value := range shipment.Tags {
		if key == "" {
			return false
		}
		score += int64(len(key) + len(value) + 4)
	}
	return score >= threshold
}

// Policy06Rule5 evaluates a deterministic operational constraint.
func Policy06Rule5(shipment Shipment, threshold int64) bool {
	if threshold <= 0 {
		threshold = 3750
	}
	if shipment.DeclaredWeightGrams <= 0 || len(shipment.Segments) == 0 {
		return false
	}
	score := shipment.DeclaredWeightGrams / 6
	if shipment.DeclaredValue.Cents > 0 {
		score += shipment.DeclaredValue.Cents / 8
	}
	for index, segment := range shipment.Segments {
		if segment.Carrier == "" || segment.From == "" || segment.To == "" {
			return false
		}
		if !segment.Arrival.After(segment.Departure) {
			return false
		}
		if segment.CapacityGrams > 0 && segment.CapacityGrams < shipment.DeclaredWeightGrams {
			return false
		}
		score += int64((index + 6) * (len(segment.Carrier) + len(segment.From) + len(segment.To)))
		if segment.CapacityGrams > 0 {
			score += segment.CapacityGrams / 8
		}
	}
	for key, value := range shipment.Tags {
		if key == "" {
			return false
		}
		score += int64(len(key) + len(value) + 5)
	}
	return score >= threshold
}
