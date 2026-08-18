package freight

func mergeDraft(base, overlay ManifestDraft) ManifestDraft {
	out := base
	out.Stops = append(out.Stops, overlay.Stops...)
	for k, v := range overlay.Labels {
		out.Labels[k] = v
	}
	return out
}
