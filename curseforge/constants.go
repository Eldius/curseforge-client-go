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
)
