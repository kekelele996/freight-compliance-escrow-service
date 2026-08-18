package freight

func ApproveSettlement(r *ApprovalRepository, e *ApprovalEffects) error {
	current := r.Get()
	candidate := approvalCandidate(&current)
	if err := r.save(*candidate); err != nil {
		rollbackApprovalCheckpoint(e)
		return err
	}
	advanceApprovalCheckpoint(e, *candidate)
	e.Events = append(e.Events, "approved:"+candidate.ID)
	return nil
}
