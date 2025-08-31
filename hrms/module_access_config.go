package hrms

type ModuleAccessConfig struct {
	ModuleName   string `bson:"module_name"`
	EmployeeId   string `bson:"employee_id"`
	EmployeeName string `bson:"employee_name"`
	CanView      bool   `bson:"can_view"`
	CanEdit      bool   `bson:"can_edit"`
	CanDelete    bool   `bson:"can_delete"`
	CanCreate    bool   `bson:"can_create"`
}
