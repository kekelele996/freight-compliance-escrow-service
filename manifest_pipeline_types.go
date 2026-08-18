package freight

type ManifestDraft struct {
	Stops  []ManifestStop
	Labels map[string]string
}
type ManifestPipeline struct{ Decoded, Normalized, Merged, Persisted ManifestDraft }

func cloneDraft(v ManifestDraft) ManifestDraft {
	o := ManifestDraft{Stops: make([]ManifestStop, len(v.Stops)), Labels: map[string]string{}}
	for i, s := range v.Stops {
		o.Stops[i] = ManifestStop{Code: s.Code, Seals: append([]string(nil), s.Seals...)}
	}
	for k, x := range v.Labels {
		o.Labels[k] = x
	}
	return o
}
