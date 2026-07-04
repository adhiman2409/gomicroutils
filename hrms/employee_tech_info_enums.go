package hrms

type InputType int

const (
	Text InputType = iota + 1
	Binary
	StarRating
	SingleSelect
	MultiSelect
	Number
	None
)

func (r InputType) String() string {
	return [...]string{"Text", "Binary", "StarRating", "SingleSelect", "MultiSelect", "Number", "None"}[r-1]
}

func (r InputType) EnumIndex() int {
	return int(r)
}

func GetAllInputTypes() []string {
	return []string{"Text", "Binary", "StarRating", "SingleSelect", "MultiSelect", "Number", "None"}
}

func InputTypeFromString(s string) InputType {
	switch s {
	case "Binary":
		return Binary
	case "StarRating":
		return StarRating
	case "SingleSelect":
		return SingleSelect
	case "MultiSelect":
		return MultiSelect
	case "Number":
		return Number
	case "None":
		return None
	default:
		return Text
	}
}

type ResignationState int

const (
	RegularEmployment ResignationState = iota
	ResignationSubmitted
	ResignationAccepted
	ResignationRejected
	ResignationRevoked
	ServingNoticePeriod
	ServingPIP
	Exited
)

func (r ResignationState) String() string {
	return [...]string{"RegularEmployment", "ResignationSubmitted", "ResignationAccepted", "ResignationRejected", "ResignationRevoked", "ServingNoticePeriod", "ServingPIP", "Exited"}[r]
}

func (r ResignationState) EnumIndex() int {
	return int(r)
}

func GetAllResignationState() []string {
	return []string{"RegularEmployment", "ResignationSubmitted", "ResignationAccepted", "ResignationRejected", "ResignationRevoked", "ServingNoticePeriod", "ServingPIP", "Exited"}
}

func ResignationStateFromString(s string) ResignationState {
	switch s {
	case "ResignationSubmitted":
		return ResignationSubmitted
	case "ResignationAccepted":
		return ResignationAccepted
	case "ResignationRejected":
		return ResignationRejected
	case "ResignationRevoked":
		return ResignationRevoked
	case "ServingNoticePeriod":
		return ServingNoticePeriod
	case "ServingPIP":
		return ServingPIP
	case "Exited":
		return Exited
	default:
		return RegularEmployment
	}
}
