package freight

func normalizeDraft(v ManifestDraft) ManifestDraft {
	if len(v.Stops) > 0 {
		v.Stops[0].Seals = append(v.Stops[0].Seals, "normalized")
	}
	v.Labels["normalized"] = "yes"
	return v
}
