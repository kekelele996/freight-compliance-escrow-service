package freight

func ApproveSettlement(r *ApprovalRepository, e *ApprovalEffects) error {
	current := r.Get()
	candidate := approvalCandidate(&current)
	staged := stageApprovalCheckpoint(*candidate)
	if err := r.save(*candidate); err != nil {
		rollbackApprovalCheckpoint(e)
		return err
	}
	commitApprovalCheckpoint(e, staged)
	e.Events = append(e.Events, "approved:"+candidate.ID)
	return nil
}
