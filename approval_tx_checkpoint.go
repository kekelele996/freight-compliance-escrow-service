package freight

func stageApprovalCheckpoint(v ApprovalRecord) ApprovalEffects {
	return ApprovalEffects{Checkpoint: v.Version, Observed: v}
}
func commitApprovalCheckpoint(dst *ApprovalEffects, staged ApprovalEffects) {
	dst.Checkpoint = staged.Checkpoint
	dst.Observed = staged.Observed
}
func rollbackApprovalCheckpoint(e *ApprovalEffects) { e.Checkpoint = 0; e.Observed = ApprovalRecord{} }
