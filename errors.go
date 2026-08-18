package freight

import "errors"

var (
	ErrInvalidAmount    = errors.New("amount must be non-negative")
	ErrCurrencyMismatch = errors.New("currency mismatch")
	ErrDuplicateRequest = errors.New("duplicate request")
	ErrNotFound         = errors.New("not found")
	ErrWindowClosed     = errors.New("service window is closed")
	ErrValidation       = errors.New("validation failed")
	ErrConflict         = errors.New("optimistic concurrency conflict")
)
