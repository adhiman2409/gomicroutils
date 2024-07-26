package gcs

import "strings"

func GetUpdatedDomain(domain string) string {
	return "unirms_" + strings.Replace(domain, ".", "_", -1)
}
