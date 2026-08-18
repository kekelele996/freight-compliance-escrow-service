package freight

func invokeDeliverySend(v DeliveryRecord, send func(string) error) error {
	if v.State != "inflight" {
		return ErrConflict
	}
	if send == nil {
		return ErrValidation
	}
	return send(v.ID)
}
