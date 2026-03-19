package hrms

type ShiftType int

const (
	Regular ShiftType = iota + 1
	Morning
	Evening
	Night
	WeeklyOff
	Holiday
	WorkFromHome
)

func (r ShiftType) String() string {
	if r == 0 {
		return "Regular"
	}
	return [...]string{"Regular", "Morning", "Evening", "Night", "WeeklyOff", "Holiday", "WorkFromHome"}[r-1]
}

func (r ShiftType) EnumIndex() int {
	return int(r)
}

func GetAllShiftType() []string {
	return []string{"Regular", "Morning", "Evening", "Night", "WeeklyOff", "Holiday", "WorkFromHome"}
}

func ShiftTypeFromString(s string) ShiftType {
	switch s {
	case "Morning":
		return Morning
	case "Evening":
		return Evening
	case "Night":
		return Night
	case "WeeklyOff":
		return WeeklyOff
	case "Holiday":
		return Holiday
	case "WorkFromHome":
		return WorkFromHome
	default:
		return Regular
	}
}
