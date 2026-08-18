package freight

func (p *ManifestPipeline) Run(input, overlay ManifestDraft) {
	p.Decoded = decodeDraft(input)
	p.Normalized = normalizeDraft(p.Decoded)
	p.Merged = mergeDraft(p.Normalized, overlay)
	p.Persisted = persistDraft(p.Merged)
}
