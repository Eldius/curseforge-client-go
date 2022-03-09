package curseforge

const (
	//1=Draft
	CoreStatusDraft = 1
	//2=Test
	CoreStatusTest = 2
	//3=PendingReview
	CoreStatusPendingReview = 3
	//4=Rejected
	CoreStatusRejected = 4
	//5=Approved
	CoreStatusApproved = 5
	//6=Live
	CoreStatusLive = 6

	CoreStatusTextDraft         = "DRAFT"
	CoreStatusTextTest          = "TEST"
	CoreStatusTextPendingReview = "PENDING_REVIEW"
	CoreStatusTextRejected      = "REJECTED"
	CoreStatusTextApproved      = "APPROVED"
	CoreStatusTextLive          = "LIVE"

	CoreApiStatusPrivate = 1
	CoreApiStatusPublic  = 2

	//ModStatus
	ModStatusNew             = 1
	ModStatusChangesRequired = 2
	ModStatusUnderSoftReview = 3
	ModStatusApproved        = 4
	ModStatusRejected        = 5
	ModStatusChangesMade     = 6
	ModStatusInactive        = 7
	ModStatusAbandoned       = 8
	ModStatusDeleted         = 9
	ModStatusUnderReview     = 10

	FileStatusProcessing         = 1
	FileStatusChangesRequired    = 2
	FileStatusUnderReview        = 3
	FileStatusApproved           = 4
	FileStatusRejected           = 5
	FileStatusMalwareDetected    = 6
	FileStatusDeleted            = 7
	FileStatusArchived           = 8
	FileStatusTesting            = 9
	FileStatusReleased           = 10
	FileStatusReadyForReview     = 11
	FileStatusDeprecated         = 12
	FileStatusBaking             = 13
	FileStatusAwaitingPublishing = 14
	FileStatusFailedPublishing   = 15
)

var (
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
