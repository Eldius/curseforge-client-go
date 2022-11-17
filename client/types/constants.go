package types

var (
	// ModStatusMap is the mod status map
	ModStatusMap = map[int]string{
		1:  "New",
		2:  "ChangesRequired",
		3:  "UnderSoftReview",
		4:  "Approved",
		5:  "Rejected",
		6:  "ChangesMade",
		7:  "Inactive",
		8:  "Abandoned",
		9:  "Deleted",
		10: "UnderReview",
	}

	// FileStatusMap is the file status map
	FileStatusMap = map[int]string{
		1:  "Processing",
		2:  "ChangesRequired",
		3:  "UnderReview",
		4:  "Approved",
		5:  "Rejected",
		6:  "MalwareDetected",
		7:  "Deleted",
		8:  "Archived",
		9:  "Testing",
		10: "Released",
		11: "ReadyForReview",
		12: "Deprecated",
		13: "Baking",
		14: "AwaitingPublishing",
		15: "FailedPublishing",
	}
)
