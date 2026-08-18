package main

import (
	"fmt"
	freight "github.com/kekelele996/freight-compliance-escrow-service"
	"time"
)

func main() {
	cfg, e := freight.ParseConfig(map[string]string{"currency": "USD"})
	if e != nil {
		panic(e)
	}
	fmt.Printf("freight compliance escrow service ready (%s) at %s\n", cfg.DefaultCurrency, time.Now().UTC().Format(time.RFC3339))
}
