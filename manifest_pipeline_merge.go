package freight

func mergeDraft(base, overlay ManifestDraft) ManifestDraft {
	out := cloneDraft(base)
	extra := cloneDraft(overlay)
	out.Stops = append(out.Stops, extra.Stops...)
	for k, v := range extra.Labels {
		out.Labels[k] = v
	}
	return out
}
