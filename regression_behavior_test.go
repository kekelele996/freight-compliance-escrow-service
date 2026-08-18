package freight

import "testing"

func TestPageCursorCommitRequiresParseScopeAndTransaction(t *testing.T) {
	s := &PageSession{Checkpoint: PageCheckpoint{Tenant: "north", Filter: "open", Direction: "asc", LastID: "a1", Version: 4}}
	before := s.Checkpoint
	if e := s.Resume("%%%", "north", "open", "asc"); e == nil {
		t.Fatal("malformed cursor accepted")
	}
	if s.Checkpoint != before {
		t.Fatalf("parse changed checkpoint: %+v", s.Checkpoint)
	}
	foreign := EncodeScopedCursor(ScopedCursor{Tenant: "south", Filter: "closed", Direction: "desc", LastID: "z9"})
	if e := s.Resume(foreign, "north", "open", "asc"); e == nil {
		t.Fatal("foreign cursor accepted")
	}
	if s.Checkpoint != before {
		t.Fatalf("scope changed checkpoint: %+v", s.Checkpoint)
	}
	valid := EncodeScopedCursor(ScopedCursor{Tenant: " North ", Filter: "open", Direction: "asc", LastID: "a2"})
	if e := s.Resume(valid, "north", "open", "asc"); e != nil {
		t.Fatal(e)
	}
	if s.Checkpoint.Version != 5 || s.Checkpoint.LastID != "a2" || s.Checkpoint.Tenant != "north" {
		t.Fatalf("valid cursor not committed: %+v", s.Checkpoint)
	}
}
