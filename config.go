package freight

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseConfig(v map[string]string) (Config, error) {
	c := Config{DefaultCurrency: strings.ToUpper(strings.TrimSpace(v["currency"])), WorkerParallelism: 4, MaxWeightGrams: 250000}
	if c.DefaultCurrency == "" {
		return Config{}, fmt.Errorf("%w: currency", ErrValidation)
	}
	if raw := v["workers"]; raw != "" {
		n, e := strconv.Atoi(raw)
		if e != nil || n < 1 || n > 128 {
			return Config{}, fmt.Errorf("%w: workers", ErrValidation)
		}
		c.WorkerParallelism = n
	}
	if raw := v["max_weight_grams"]; raw != "" {
		n, e := strconv.ParseInt(raw, 10, 64)
		if e != nil || n < 1 {
			return Config{}, fmt.Errorf("%w: max weight", ErrValidation)
		}
		c.MaxWeightGrams = n
	}
	return c, nil
}
