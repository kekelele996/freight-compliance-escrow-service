package freight

func approvalCandidate(v *ApprovalRecord) *ApprovalRecord {
	if v == nil {
		return nil
	}
	out := *v
	out.State = "approved"
	out.Version++
	return &out
}
