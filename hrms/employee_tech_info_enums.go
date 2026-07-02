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
	ResignationSubmitted ResignationState = iota + 1
	ResignationAccepted
	ResignationRejected
	ResignationRevoked
)

func (r ResignationState) String() string {
	return [...]string{"ResignationSubmitted", "ResignationAccepted", "ResignationRejected", "ResignationRevoked"}[r-1]
}

func (r ResignationState) EnumIndex() int {
	return int(r)
}

func GetAllResignationState() []string {
	return []string{"ResignationSubmitted", "ResignationAccepted", "ResignationRejected", "ResignationRevoked"}
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
	default:
		return ResignationSubmitted
	}
}
