package grpcclient

type AuthInfo struct {
	Authorised  bool
	Tenant      string
	Domain      string
	Department  string
	Name        string
	EmailId     string
	PhoneNumber string
	Roles       []string
}
