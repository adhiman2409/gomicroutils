package hrms

type SalaryStatus int

const (
	Pending SalaryStatus = iota + 1
	Approved
	Paid
	TemporaryOnHold
	Cancelled
)

func (r SalaryStatus) String() string {
	return [...]string{"Pending", "Approved", "Paid", "TemporaryOnHold", "Cancelled"}[r-1]
}

func (r SalaryStatus) EnumIndex() int {
	return int(r)
}

func GetAllSalaryStatus() []string {
	return []string{"Pending", "Approved", "Paid", "TemporaryOnHold", "Cancelled"}
}

func SalaryStatusFromString(s string) SalaryStatus {
	switch s {
	case "Approved":
		return Approved
	case "Paid":
		return Paid
	case "TemporaryOnHold":
		return TemporaryOnHold
	case "Cancelled":
		return Cancelled
	default:
		return Pending
	}
}
