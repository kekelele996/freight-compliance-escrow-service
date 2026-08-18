package freight

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func DecodeManifest(raw []byte) (Manifest, error) {
	var m Manifest
	d := json.NewDecoder(bytes.NewReader(raw))
	d.DisallowUnknownFields()
	if e := d.Decode(&m); e != nil {
		return Manifest{}, e
	}
	if m.ShipmentID == "" || len(m.Stops) == 0 {
		return Manifest{}, fmt.Errorf("%w: missing fields", ErrValidation)
	}
	seen := map[string]bool{}
	for _, s := range m.Stops {
		if s.Code == "" || seen[s.Code] {
			return Manifest{}, fmt.Errorf("%w: invalid stops", ErrValidation)
		}
		seen[s.Code] = true
	}
	return cloneManifest(m), nil
}
func cloneManifest(m Manifest) Manifest {
	out := Manifest{ShipmentID: m.ShipmentID, Stops: make([]ManifestStop, len(m.Stops))}
	for i, s := range m.Stops {
		out.Stops[i] = ManifestStop{Code: s.Code, Seals: append([]string(nil), s.Seals...)}
	}
	return out
}
