package hrms

type EmploymentStatus int

const (
	Active EmploymentStatus = iota + 1
	NoticePeriod
	OnPIP
	Suspended
	Inactive
)

func (r EmploymentStatus) String() string {
	return [...]string{"Active", "NoticePeriod", "OnPIP", "Suspended", "Inactive"}[r-1]
}

func (r EmploymentStatus) EnumIndex() int {
	return int(r)
}

func GetAllEmploymentStatus() []string {
	return []string{"Active", "NoticePeriod", "OnPIP", "Suspended", "Inactive"}
}

func EmploymentStatusEnumFromString(s string) EmploymentStatus {
	if s == "Active" {
		return Active
	} else if s == "NoticePeriod" {
		return NoticePeriod
	} else if s == "OnPIP" {
		return OnPIP
	} else if s == "Suspended" {
		return Suspended
	} else {
		return Inactive
	}
}
