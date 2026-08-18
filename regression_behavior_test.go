package freight

import (
	"errors"
	"testing"
)

func TestDeliveryFailureIsFinalizedAsRetryable(t *testing.T) {
	q := NewDeliveryQueue()
	q.Add("e1")
	sent := 0
	err := q.PublishOne("e1", func(string) error {
		sent++
		active := q.Get("e1")
		if active.State != "inflight" || active.Attempts != 1 {
			t.Fatalf("claim=%+v", active)
		}
		return errors.New("broker unavailable")
	})
	if err == nil || sent != 1 {
		t.Fatalf("send err=%v count=%d", err, sent)
	}
	v := q.Get("e1")
	if v.State != "pending" || v.Attempts != 1 || v.LastError == "" {
		t.Fatalf("failure=%+v", v)
	}
	if ids := q.PendingIDs(); len(ids) != 1 || ids[0] != "e1" {
		t.Fatalf("pending=%v", ids)
	}

	q.Add("e2")
	if err = q.PublishOne("e2", nil); !errors.Is(err, ErrValidation) {
		t.Fatalf("nil sender err=%v", err)
	}
	v = q.Get("e2")
	if v.State != "pending" || v.Attempts != 1 || v.LastError == "" {
		t.Fatalf("nil sender state=%+v", v)
	}
}
