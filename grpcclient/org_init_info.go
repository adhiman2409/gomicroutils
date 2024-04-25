package grpcclient

type OrgInitInfo struct {
	Id                string
	Domain            string
	OrgName           string
	AdminEmailId      string
	AdminPhoneNumber  string
	AdminName         string
	Department        string
	EmailVerified     bool
	Password          string
	Role              string
	FirstLoginPending bool
	Status            string
}
