package hrms

type ExpenseState int

const (
	EApplied ExpenseState = iota + 1
	EPrimaryApproved
	EPrimaryRejected
	ESecondaryApproved
	ESecondaryRejected
	EFinanceAccepted
	EFinanceRejected
	EPaid
	EUnknown
)

func (r ExpenseState) String() string {
	return [...]string{"EApplied", "EPrimaryApproved", "EPrimaryRejected", "ESecondaryApproved",
		"ESecondaryRejected", "EFinanceAccepted", "EFinanceRejected", "EPaid", "EUnknown"}[r-1]
}

func (r ExpenseState) EnumIndex() int {
	return int(r)
}

func GetAllExpenseStates() []string {
	return []string{"EApplied", "EPrimaryApproved", "EPrimaryRejected", "ESecondaryApproved",
		"ESecondaryRejected", "EFinanceAccepted", "EFinanceRejected", "EPaid", "EUnknown"}
}

func ExpenseStateFromString(s string) ExpenseState {
	if s == "EApplied" {
		return EApplied
	} else if s == "EPrimaryApproved" {
		return EPrimaryApproved
	} else if s == "EPrimaryRejected" {
		return EPrimaryRejected
	} else if s == "ESecondaryApproved" {
		return ESecondaryApproved
	} else if s == "ESecondaryRejected" {
		return ESecondaryRejected
	} else if s == "EFinanceAccepted" {
		return EFinanceAccepted
	} else if s == "EFinanceRejected" {
		return EFinanceRejected
	} else if s == "EPaid" {
		return EPaid
	} else {
		return EUnknown
	}
}
