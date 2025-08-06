package hrms

type AccessConfigType string

const (
	DepartmentLevel      AccessConfigType = "department_level"
	DepartmentRole       AccessConfigType = "department_role"
	DepartmentEmployee   AccessConfigType = "department_employee"
	OrganizationRole     AccessConfigType = "organization_role"
	OrganizationEmployee AccessConfigType = "organization_employee"
)

type SidebarItem struct {
	Title  string `json:"title"`
	Icon   string `json:"icon"`
	Read   bool   `json:"read"`
	Insert bool   `json:"insert"`
	Update bool   `json:"update"`
	Delete bool   `json:"delete"`
}

type UnirmsConfig struct {
	SidebarItems []SidebarItem `json:"sidebar_items"`
}

type AccessConfig struct {
	Type   AccessConfigType `bson:"type"`
	Name   string           `bson:"name"`
	Id     string           `bson:"id"`
	Read   bool             `bson:"read"`
	Insert bool             `bson:"insert"`
	Update bool             `bson:"update"`
	Delete bool             `bson:"delete"`
}

type SidebarConfig struct {
	Id                       string         `bson:"_id"`
	Title                    string         `bson:"title"`
	Icon                     string         `bson:"icon"`
	DepartmentLevelAccess    []AccessConfig `bson:"department_level_access"`
	DepartmentRoleAccess     []AccessConfig `bson:"department_role_access"`
	DepartmentEmployeeAccess []AccessConfig `bson:"department_employee_access"`
	OrgRoleAccess            []AccessConfig `bson:"org_role_access"`
	OrgEmployeeAccess        []AccessConfig `bson:"org_employee_access"`
}
