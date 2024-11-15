package gcs

import "strings"

func GetUpdatedDomain(domain string) string {
	if domain == "" {
		return ""
	}
	return "unirms_" + strings.Replace(domain, ".", "_", -1)
}

func GetUpdatedFinanceDomain(domain string) string {
	if domain == "" {
		return ""
	}
	return "unirms_finance_" + strings.Replace(domain, ".", "_", -1)
}
