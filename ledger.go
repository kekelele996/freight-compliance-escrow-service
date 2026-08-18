package freight

import (
	"fmt"
	"sort"
)

type LedgerEntry struct {
	Account, Direction string
	Amount             Money
	Reference          string
}
type Ledger struct{ entries []LedgerEntry }

func (l *Ledger) Post(entries ...LedgerEntry) error {
	totals := map[string]int64{}
	for _, e := range entries {
		if e.Account == "" || (e.Direction != "debit" && e.Direction != "credit") {
			return fmt.Errorf("%w: malformed ledger entry", ErrValidation)
		}
		d := e.Amount.Cents
		if e.Direction == "credit" {
			d = -d
		}
		totals[e.Amount.Currency] += d
	}
	for c, total := range totals {
		if total != 0 {
			return fmt.Errorf("ledger not balanced for %s", c)
		}
	}
	l.entries = append(l.entries, entries...)
	return nil
}
func (l *Ledger) Balance(account, currency string) Money {
	var cents int64
	for _, e := range l.entries {
		if e.Account != account || e.Amount.Currency != currency {
			continue
		}
		if e.Direction == "debit" {
			cents += e.Amount.Cents
		} else {
			cents -= e.Amount.Cents
		}
	}
	return Money{currency, cents}
}
func (l *Ledger) EntriesFor(ref string) []LedgerEntry {
	out := []LedgerEntry{}
	for _, e := range l.entries {
		if e.Reference == ref {
			out = append(out, e)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Account < out[j].Account })
	return out
}
func BuildRefundEntries(ref string, amount Money) []LedgerEntry {
	return []LedgerEntry{{"carrier_payable", "debit", amount, ref}, {"customer_escrow", "credit", amount, ref}}
}
