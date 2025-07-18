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
	Countries         []string
}
