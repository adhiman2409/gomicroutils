package hrms

type Role int

const (
	Superadmin Role = iota + 1
	OrgAdmin
	OrgManager
	Recruiter
	Manager
	TeamLead
	TeamMember
	Doctor
	Candidate
	Volunteer
	Anonymous
)

func (r Role) String() string {
	return [...]string{"Superadmin", "OrgAdmin", "OrgManager", "Recruiter", "Manager",
		"TeamLead", "TeamMember", "Doctor", "Candidate", "Volunteer", "Anonymous"}[r-1]
}

func (r Role) EnumIndex() int {
	return int(r)
}

func GetAllRoles() []string {
	return []string{"Superadmin", "OrgAdmin", "OrgManager", "Recruiter", "Manager",
		"TeamLead", "TeamMember", "Doctor", "Candidate", "Volunteer", "Anonymous"}
}

func RoleEnumFromString(s string) Role {
	switch s {
	case "Superadmin":
		return Superadmin
	case "OrgAdmin":
		return OrgAdmin
	case "OrgManager":
		return OrgManager
	case "Recruiter":
		return Recruiter
	case "Manager":
		return Manager
	case "TeamLead":
		return TeamLead
	case "TeamMember":
		return TeamMember
	case "Doctor":
		return Doctor
	case "Candidate":
		return Candidate
	case "Volunteer":
		return Volunteer
	default:
		return Anonymous
	}
}
