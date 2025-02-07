package hrms

type ReimbursementState int

const (
	RPending ReimbursementState = iota + 1
	RPrimaryApproved
	RApproved
	RRejected
	RFinanceAccepted
	RFinanceRejected
	ROnHold
	RPaid
	RUnknown
)

func (r ReimbursementState) String() string {
	return [...]string{"RPending", "RPrimaryApproved", "RApproved",
		"RRejected", "RFinanceAccepted", "RFinanceRejected", "ROnHold", "RPaid", "RUnknown"}[r-1]
}

func (r ReimbursementState) EnumIndex() int {
	return int(r)
}

func GetAllReimbursementStates() []string {
	return []string{"RPending", "RPrimaryApproved", "RApproved",
		"RRejected", "RFinanceAccepted", "RFinanceRejected", "ROnHold", "RPaid", "RUnknown"}
}

func ReimbursementStateFromString(s string) ReimbursementState {
	if s == "RPending" {
		return RPending
	} else if s == "RPrimaryApproved" {
		return RPrimaryApproved
	} else if s == "RApproved" {
		return RApproved
	} else if s == "RRejected" {
		return RRejected
	} else if s == "RFinanceAccepted" {
		return RFinanceAccepted
	} else if s == "RFinanceRejected" {
		return RFinanceRejected
	} else if s == "ROnHold" {
		return ROnHold
	} else if s == "RPaid" {
		return RPaid
	} else {
		return RUnknown
	}
}
