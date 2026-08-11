package hrms

type SalaryStatus int

const (
	Pending SalaryStatus = iota + 1
	WaitingApproval
	Approved
	Processed
	Paid
	TemporaryOnHold
	NeedRevision
	Cancelled
	WaitingVerification
	Verified
	PaymentFailed
)

func (r SalaryStatus) String() string {
	return [...]string{"Pending", "WaitingApproval", "Approved", "Processed", "Paid", "TemporaryOnHold", "NeedRevision", "Cancelled", "WaitingVerification", "Verified", "PaymentFailed"}[r-1]
}

func (r SalaryStatus) EnumIndex() int {
	return int(r)
}

func GetAllSalaryStatus() []string {
	return []string{"Pending", "WaitingApproval", "Approved", "Processed", "Paid", "TemporaryOnHold", "NeedRevision", "Cancelled", "WaitingVerification", "Verified", "PaymentFailed"}
}

func SalaryStatusFromString(s string) SalaryStatus {
	switch s {
	case "WaitingApproval":
		return WaitingApproval
	case "Approved":
		return Approved
	case "Processed":
		return Processed
	case "Paid":
		return Paid
	case "TemporaryOnHold":
		return TemporaryOnHold
	case "NeedRevision":
		return NeedRevision
	case "Cancelled":
		return Cancelled
	case "WaitingVerification":
		return WaitingVerification
	case "Verified":
		return Verified
	case "PaymentFailed":
		return PaymentFailed
	default:
		return Pending
	}
}
