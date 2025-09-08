package hrms

type CaseStatus int

const (
	Submitted = iota + 1
	UnderInspection
	PendingCaseRecords
	VirtualSetupInProgress
	ApprovedVirtualSetup
	AligneerFabricationInProgress
	OngoinCase
	Delivered
	Retainers
	CCancelled
)

func (r CaseStatus) String() string {
	return [...]string{"Submitted", "UnderInspection", "PendingCaseRecords", "VirtualSetupInProgress", "ApprovedVirtualSetup", "AligneerFabricationInProgress", "OngoingCase", "Delivered", "Retainers", "CCancelled"}[r-1]
}

func (r CaseStatus) CaseEnumIndex() int {
	return int(r)
}

func GetCaseStatusTypes() []string {
	return []string{"Submitted", "UnderInspection", "PendingCaseRecords", "VirtualSetupInProgress", "ApprovedVirtualSetup", "AligneerFabricationInProgress", "OngoingCase", "Delivered", "Retainers", "CCancelled"}
}

func CaseStatusFromString(s string) CaseStatus {
	switch s {
	case "Submitted":
		return Submitted
	case "UnderInspection":
		return UnderInspection
	case "PendingCaseRecords":
		return PendingCaseRecords
	case "VirtualSetupInProgress":
		return VirtualSetupInProgress
	case "ApprovedVirtualSetup":
		return ApprovedVirtualSetup
	case "AligneerFabricationInProgress":
		return AligneerFabricationInProgress
	case "OngoingCase":
		return OngoinCase
	case "Delivered":
		return Delivered
	case "Retainers":
		return Retainers
	default:
		return CCancelled
	}
}
