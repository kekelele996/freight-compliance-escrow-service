package freight

import "testing"

func TestManifestPipelineKeepsStageOwnership(t *testing.T) {
	input := ManifestDraft{Stops: []ManifestStop{{Code: "A", Seals: []string{"s1"}}}, Labels: map[string]string{"source": "input"}}
	overlay := ManifestDraft{Stops: []ManifestStop{{Code: "B", Seals: []string{"s2"}}}, Labels: map[string]string{"carrier": "x"}}
	var p ManifestPipeline
	p.Run(input, overlay)
	input.Stops[0].Seals[0] = "input-change"
	input.Labels["source"] = "input-change"
	if p.Decoded.Stops[0].Seals[0] != "s1" || p.Decoded.Labels["source"] != "input" {
		t.Fatalf("decoded=%+v", p.Decoded)
	}
	p.Decoded.Stops[0].Seals[0] = "decoded-change"
	p.Decoded.Labels["source"] = "decoded-change"
	if p.Normalized.Stops[0].Seals[0] != "s1" || p.Normalized.Labels["source"] != "input" {
		t.Fatalf("normalized=%+v", p.Normalized)
	}
	p.Normalized.Stops[0].Seals[0] = "normalized-change"
	p.Normalized.Labels["source"] = "normalized-change"
	overlay.Stops[0].Seals[0] = "overlay-change"
	overlay.Labels["carrier"] = "overlay-change"
	if p.Merged.Stops[0].Seals[0] != "s1" || p.Merged.Stops[1].Seals[0] != "s2" || p.Merged.Labels["carrier"] != "x" {
		t.Fatalf("merged=%+v", p.Merged)
	}
	p.Merged.Stops[0].Seals[0] = "merged-change"
	p.Merged.Labels["carrier"] = "merged-change"
	if p.Persisted.Stops[0].Seals[0] != "s1" || p.Persisted.Labels["carrier"] != "x" {
		t.Fatalf("persisted=%+v", p.Persisted)
	}
}
