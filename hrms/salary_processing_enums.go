package hrms

type SalaryStatus int

const (
	Pending SalaryStatus = iota + 1
	WaitingApproval
	Approved
	Paid
	TemporaryOnHold
	NeedRevision
	Cancelled
)

func (r SalaryStatus) String() string {
	return [...]string{"Pending", "WaitingApproval", "Approved", "Paid", "TemporaryOnHold", "NeedRevision", "Cancelled"}[r-1]
}

func (r SalaryStatus) EnumIndex() int {
	return int(r)
}

func GetAllSalaryStatus() []string {
	return []string{"Pending", "WaitingApproval", "Approved", "Paid", "TemporaryOnHold", "NeedRevision", "Cancelled"}
}

func SalaryStatusFromString(s string) SalaryStatus {
	switch s {
	case "WaitingApproval":
		return WaitingApproval
	case "Approved":
		return Approved
	case "Paid":
		return Paid
	case "TemporaryOnHold":
		return TemporaryOnHold
	case "NeedRevision":
		return NeedRevision
	case "Cancelled":
		return Cancelled
	default:
		return Pending
	}
}
