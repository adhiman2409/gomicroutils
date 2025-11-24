package grpcclient

type OrgInitInfo struct {
	Id                string
	EID               string
	Domain            string
	OrgName           string
	AdminEmailId      string
	AdminPhoneNumber  string
	AdminName         string
	Department        string
	Designation       string
	EmailVerified     bool
	Password          string
	Role              string
	FirstLoginPending bool
	Status            string
	UseGoogleOAuth    bool
	Country           string
	TimeZone          string
}

type DeductExpenseInfo struct {
	Domain       string
	DepartmentId string
	FiscalYear   string
	Amount       float64
	Currency     string
	ExpenseId    string
}

type DeductExpenseInfoResponse struct {
	IsError bool
	Message string
}
