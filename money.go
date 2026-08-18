package freight

import (
	"fmt"
	"math"
)

func NewMoney(currency string, cents int64) (Money, error) {
	if cents < 0 {
		return Money{}, ErrInvalidAmount
	}
	if currency == "" {
		return Money{}, fmt.Errorf("%w: currency", ErrValidation)
	}
	return Money{currency, cents}, nil
}
func MustMoney(currency string, cents int64) Money {
	m, e := NewMoney(currency, cents)
	if e != nil {
		panic(e)
	}
	return m
}
func (m Money) Add(n Money) (Money, error) {
	if m.Currency != n.Currency {
		return Money{}, ErrCurrencyMismatch
	}
	if n.Cents > 0 && m.Cents > math.MaxInt64-n.Cents {
		return Money{}, fmt.Errorf("amount overflow")
	}
	return Money{m.Currency, m.Cents + n.Cents}, nil
}
func (m Money) Sub(n Money) (Money, error) {
	if m.Currency != n.Currency {
		return Money{}, ErrCurrencyMismatch
	}
	if m.Cents < n.Cents {
		return Money{}, ErrInvalidAmount
	}
	return Money{m.Currency, m.Cents - n.Cents}, nil
}
func (m Money) MultiplyBasisPoints(bp int64) (Money, error) {
	if bp < 0 {
		return Money{}, ErrInvalidAmount
	}
	if m.Cents > math.MaxInt64/10000 {
		return Money{}, fmt.Errorf("amount overflow")
	}
	return Money{m.Currency, (m.Cents*bp + 5000) / 10000}, nil
}
func (m Money) String() string { return fmt.Sprintf("%s %.2f", m.Currency, float64(m.Cents)/100) }
