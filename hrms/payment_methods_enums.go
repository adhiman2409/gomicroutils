package hrms

type PaymentMethod int

const (
	Razorpay PaymentMethod = iota + 1
	Bank
	Cash
	Cheque
	Other
)

func (r PaymentMethod) String() string {
	return [...]string{"Razorpay", "Bank", "Cash", "Cheque", "Other"}[r-1]
}

func (r PaymentMethod) EnumIndex() int {
	return int(r)
}

func GetAllPaymentMethods() []string {
	return []string{"Razorpay", "Bank", "Cash", "Cheque", "Other"}
}

func GetOfflinePaymentMethods() []string {
	return []string{"Bank", "Cash", "Cheque"}
}

func PaymentMethodEnumFromString(s string) PaymentMethod {
	switch s {
	case "Razorpay":
		return Razorpay
	case "Bank":
		return Bank
	case "Cash":
		return Cash
	case "Cheque":
		return Cheque
	default:
		return Other
	}
}
