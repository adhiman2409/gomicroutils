package grpcclient

type EmailTokenInfo struct {
	TenantId     string
	Domain       string
	SubDomain    string
	EmailId      string
	SecretId     string
	Name         string
	IsError      bool
	ErrorMessage string
}
