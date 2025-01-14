package hrms

type ShiftType int

const (
	Regular ShiftType = iota + 1
	Morning
	Afternoon
	Evening
	Night
)

func (r ShiftType) String() string {
	return [...]string{"Regular", "Morning", "Afternoon", "Evening", "Night"}[r-1]
}

func (r ShiftType) EnumIndex() int {
	return int(r)
}

func GetAllShiftType() []string {
	return []string{"Regular", "Morning", "Afternoon", "Evening", "Night"}
}

func ShiftTypeFromString(s string) ShiftType {
	if s == "Morning" {
		return Morning
	} else if s == "Afternoon" {
		return Afternoon
	} else if s == "Evening" {
		return Evening
	} else if s == "Night" {
		return Night
	} else {
		return Regular
	}
}
