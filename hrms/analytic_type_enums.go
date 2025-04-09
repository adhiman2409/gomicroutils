package hrms

type AnalyticsType int

const (
	Hirings AnalyticsType = iota + 1
	Exits
	Employees
	Employments
	Events
)

func (r AnalyticsType) String() string {
	return [...]string{"Hirings", "Exits", "Employees", "Employments", "Events"}[r-1]
}

func (r AnalyticsType) EnumIndex() int {
	return int(r)
}

func GetAllAnalyticsType() []string {
	return []string{"Hirings", "Exits", "Employees", "Employments", "Events"}
}

func AnalyticsTypeFromString(s string) AnalyticsType {
	if s == "Hirings" {
		return Hirings
	} else if s == "Exits" {
		return Exits
	} else if s == "Employees" {
		return Employees
	} else if s == "Employments" {
		return Employments
	} else {
		return Events
	}
}
