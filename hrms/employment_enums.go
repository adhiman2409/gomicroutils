package hrms

type EmploymentStatus int

const (
	Active EmploymentStatus = iota + 1
	ServingNotice
	OnPIP
	Suspended
	Separated
)

func (r EmploymentStatus) String() string {
	return [...]string{"Active", "ServingNotice", "OnPIP", "Suspended", "Separated"}[r-1]
}

func (r EmploymentStatus) EnumIndex() int {
	return int(r)
}

func GetAllEmploymentStatus() []string {
	return []string{"Active", "ServingNotice", "OnPIP", "Suspended", "Separated"}
}

func EmploymentStatusEnumFromString(s string) EmploymentStatus {
	if s == "Active" {
		return Active
	} else if s == "ServingNotice" {
		return ServingNotice
	} else if s == "OnPIP" {
		return OnPIP
	} else if s == "Suspended" {
		return Suspended
	} else {
		return Separated
	}
}
