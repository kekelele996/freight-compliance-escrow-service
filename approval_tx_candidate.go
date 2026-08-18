package freight

func approvalCandidate(v *ApprovalRecord) *ApprovalRecord {
	v.State = "approved"
	v.Version++
	return v
}
