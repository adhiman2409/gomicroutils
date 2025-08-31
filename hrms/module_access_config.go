package hrms

type ModuleAccessConfig struct {
	ModuleName   string `json:"module_name"`
	EmployeeId   string `json:"employee_id"`
	EmployeeName string `json:"employee_name"`
	CanView      bool   `json:"can_view"`
	CanEdit      bool   `json:"can_edit"`
	CanDelete    bool   `json:"can_delete"`
	CanCreate    bool   `json:"can_create"`
}
