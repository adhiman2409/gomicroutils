package hrms

type AnalyticsType int

const (
	Hirings AnalyticsType = iota + 1
	Exits
	Employees
	Employments
	Events
	ImageURL
	About
	PersonalDetails
	BankDetails
	Skills
	EducationDetails
	ExperienceDetails
	ProjectDetails
	DocumentDetails
)

func (r AnalyticsType) String() string {
	return [...]string{"Hirings", "Exits", "Employees", "Employments", "Events", "ImageURL", "About",
		"PersonalDetails", "BankDetails", "PrimarySkill", "Skills", "EducationDetails",
		"ExperienceDetails", "ProjectDetails", "DocumentDetails"}[r-1]
}

func (r AnalyticsType) EnumIndex() int {
	return int(r)
}

func GetAllAnalyticsType() []string {
	return []string{"Hirings", "Exits", "Employees", "Employments", "Events", "ImageURL", "About",
		"PersonalDetails", "BankDetails", "PrimarySkill", "Skills", "EducationDetails",
		"ExperienceDetails", "ProjectDetails", "DocumentDetails"}
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
	} else if s == "ImageURL" {
		return ImageURL
	} else if s == "About" {
		return About
	} else if s == "PersonalDetails" {
		return PersonalDetails
	} else if s == "BankDetails" {
		return BankDetails
	} else if s == "Skills" {
		return Skills
	} else if s == "EducationDetails" {
		return EducationDetails
	} else if s == "ExperienceDetails" {
		return ExperienceDetails
	} else if s == "ProjectDetails" {
		return ProjectDetails
	} else if s == "DocumentDetails" {
		return DocumentDetails
	} else {
		return Events
	}
}
