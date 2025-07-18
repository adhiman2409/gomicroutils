package grpcclient

import (
	"encoding/json"
	"net/http"
)

type AuthInfo struct {
	Authorised  bool
	Tenant      string
	Domain      string
	Department  string
	Name        string
	EmailId     string
	PhoneNumber string
	Role        string
	OrgName     string
	Country     string
	TZone       string
}

func GetAuthInfo(r *http.Request) AuthInfo {
	authInfo := AuthInfo{
		Authorised:  false,
		Tenant:      "",
		Domain:      "",
		Department:  "",
		Name:        "",
		EmailId:     "",
		PhoneNumber: "",
		Role:        "anonymous",
		OrgName:     "",
		Country:     "India",
		TZone:       "Asia/Kolkata",
	}
	claims := r.Context().Value("claims")
	claimsString, ok := claims.(string)
	if ok {
		json.Unmarshal([]byte(claimsString), &authInfo)
	}
	return authInfo
}

func GetCorrelationId(r *http.Request) string {
	cid := r.Context().Value("correlation_id")
	id, ok := cid.(string)
	if !ok {
		return ""
	}
	return id
}
