package freight

func invokeDeliverySend(v DeliveryRecord, send func(string) error) error { return send(v.ID) }
