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
	Candidate
	Anonymous
)

func (r Role) String() string {
	return [...]string{"Superadmin", "OrgAdmin", "OrgManager", "Recruiter", "Manager",
		"TeamLead", "TeamMember", "Candidate", "Anonymous"}[r-1]
}

func (r Role) EnumIndex() int {
	return int(r)
}

func GetAllRoles() []string {
	return []string{"Superadmin", "OrgAdmin", "OrgManager", "Recruiter", "Manager",
		"TeamLead", "TeamMember", "Candidate", "Anonymous"}
}

func RoleEnumFromString(s string) Role {
	if s == "Superadmin" {
		return Superadmin
	} else if s == "OrgAdmin" {
		return OrgAdmin
	} else if s == "OrgManager" {
		return OrgManager
	} else if s == "Recruiter" {
		return Recruiter
	} else if s == "Manager" {
		return Manager
	} else if s == "TeamLead" {
		return TeamLead
	} else if s == "TeamMember" {
		return TeamMember
	} else if s == "Candidate" {
		return Candidate
	} else {
		return Anonymous
	}
}
