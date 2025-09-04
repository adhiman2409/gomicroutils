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
	switch s {
	case "EApplied":
		return EApplied
	case "EPrimaryApproved":
		return EPrimaryApproved
	case "EPrimaryRejected":
		return EPrimaryRejected
	case "ESecondaryApproved":
		return ESecondaryApproved
	case "ESecondaryRejected":
		return ESecondaryRejected
	case "EFinanceAccepted":
		return EFinanceAccepted
	case "EFinanceRejected":
		return EFinanceRejected
	case "EPaid":
		return EPaid
	default:
		return EUnknown
	}
}
