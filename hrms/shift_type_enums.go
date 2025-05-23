package hrms

type ShiftType int

const (
	Regular ShiftType = iota + 1
	Morning
	Evening
	Night
	WeeklyOff
	Holiday
)

func (r ShiftType) String() string {
	if r == 0 {
		return "Regular"
	}
	return [...]string{"Regular", "Morning", "Evening", "Night", "WeeklyOff", "Holiday"}[r-1]
}

func (r ShiftType) EnumIndex() int {
	return int(r)
}

func GetAllShiftType() []string {
	return []string{"Regular", "Morning", "Evening", "Night", "WeeklyOff", "Holiday"}
}

func ShiftTypeFromString(s string) ShiftType {
	if s == "Morning" {
		return Morning
	} else if s == "Evening" {
		return Evening
	} else if s == "Night" {
		return Night
	} else if s == "WeeklyOff" {
		return WeeklyOff
	} else if s == "Holiday" {
		return Holiday
	} else {
		return Regular
	}
}
