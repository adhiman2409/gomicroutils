package hrms

type ExpenseState int

const (
	ECreated ExpenseState = iota + 1
	ESubmitted
	EPrimaryApproved
	ESecondaryApproved
	EApproved
	EPrimaryRejected
	ESecondaryRejected
	ERejected
	EFinanceAccepted
	EFinanceRejected
	EOnHold
	EPaid
	EUnknown
)

func (r ExpenseState) String() string {
	return [...]string{"ECreated", "ESubmitted", "EPrimaryApproved", "ESecondaryApproved", "EApproved",
		"EPrimaryRejected", "ESecondaryRejected", "ERejected", "EFinanceAccepted", "EFinanceRejected", "EOnHold", "EPaid", "EUnknown"}[r-1]
}

func (r ExpenseState) EnumIndex() int {
	return int(r)
}

func GetAllExpenseStates() []string {
	return []string{"ECreated", "ESubmitted", "EPrimaryApproved", "ESecondaryApproved", "EApproved",
		"EPrimaryRejected", "ESecondaryRejected", "ERejected", "EFinanceAccepted", "EFinanceRejected", "EOnHold", "EPaid", "EUnknown"}
}

func ExpenseStateFromString(s string) ExpenseState {
	if s == "ECreated" {
		return ECreated
	} else if s == "ESubmitted" {
		return ESubmitted
	} else if s == "EPrimaryApproved" {
		return EPrimaryApproved
	} else if s == "ESecondaryApproved" {
		return ESecondaryApproved
	} else if s == "EApproved" {
		return EApproved
	} else if s == "EPrimaryRejected" {
		return EPrimaryRejected
	} else if s == "ESecondaryRejected" {
		return ESecondaryRejected
	} else if s == "ERejected" {
		return ERejected
	} else if s == "EFinanceAccepted" {
		return EFinanceAccepted
	} else if s == "EFinanceRejected" {
		return EFinanceRejected
	} else if s == "EOnHold" {
		return EOnHold
	} else if s == "EPaid" {
		return EPaid
	} else {
		return EUnknown
	}
}
