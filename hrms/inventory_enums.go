package hrms

type Frequency int

const (
	OneTime Frequency = iota + 1
	Monthly
	Quarterly
	HalfYearly
	Yearly
	TwoYearly
	ThreeYearly
	FourYearly
	FiveYearly
)

func (r Frequency) String() string {
	return [...]string{"OneTime", "Monthly", "Quarterly", "HalfYearly", "Yearly", "TwoYearly", "ThreeYearly", "FourYearly", "FiveYearly"}[r-1]
}

func (r Frequency) AssetEnumIndex() int {
	return int(r)
}

func GetFrequencyTypes() []string {
	return []string{"OneTime", "Monthly", "Quarterly", "HalfYearly", "Yearly", "TwoYearly", "ThreeYearly", "FourYearly", "FiveYearly"}
}

func AssetFrequencyFromString(s string) Frequency {
	switch s {
	case "FiveYearly":
		return FiveYearly
	case "FourYearly":
		return FourYearly
	case "ThreeYearly":
		return ThreeYearly
	case "TwoYearly":
		return TwoYearly
	case "HalfYearly":
		return HalfYearly
	case "Yearly":
		return Yearly
	case "Quarterly":
		return Quarterly
	case "Monthly":
		return Monthly
	default:
		return OneTime
	}
}

type AssetStatus int

const (
	Available = iota + 1
	Allocated
	InTransit
	Stolen
	Damaged
	UnderMaintenance
	Retired
)

func (r AssetStatus) String() string {
	return [...]string{"Available", "Allocated", "InTransit", "Stolen", "Damaged", "UnderMaintenance", "Retired"}[r-1]
}

func (r AssetStatus) AssetEnumIndex() int {
	return int(r)
}

func GetStatusTypes() []string {
	return []string{"Available", "Allocated", "InTransit", "Stolen", "Damaged", "UnderMaintenance", "Retired"}
}

func AssetStatusFromString(s string) AssetStatus {
	switch s {
	case "Available":
		return Available
	case "Allocated":
		return Allocated
	case "InTransit":
		return InTransit
	case "UnderMaintenance":
		return UnderMaintenance
	case "Stolen":
		return Stolen
	case "Damaged":
		return Damaged
	default:
		return Retired
	}
}
