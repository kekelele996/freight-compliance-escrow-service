package freight

func normalizeDraft(v ManifestDraft) ManifestDraft {
	o := cloneDraft(v)
	if len(o.Stops) > 0 {
		o.Stops[0].Seals = append(o.Stops[0].Seals, "normalized")
	}
	o.Labels["normalized"] = "yes"
	return o
}
