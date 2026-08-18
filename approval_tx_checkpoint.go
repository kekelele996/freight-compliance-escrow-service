package freight

func advanceApprovalCheckpoint(e *ApprovalEffects, v ApprovalRecord) {
	e.Checkpoint = v.Version
	e.Observed = v
}
func rollbackApprovalCheckpoint(e *ApprovalEffects) {}
