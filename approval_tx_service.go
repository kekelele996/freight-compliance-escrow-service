package freight

func ApproveSettlement(r *ApprovalRepository, e *ApprovalEffects) error {
	current := r.Get()
	candidate := approvalCandidate(&current)
	advanceApprovalCheckpoint(e, *candidate)
	e.Events = append(e.Events, "approved:"+candidate.ID)
	return r.save(*candidate)
}
