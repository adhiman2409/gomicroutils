package hrms

import (
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
)

const ORG_COLLECTION_ORGANIZATION = "organization"
const ORG_COLLECTION_DEPARTMENTS = "org-departments"
const ORG_COLLECTION_DOCUMENTS = "org-documents"
const ORG_COLLECTION_EMPLOYEE_TYPE = "org-employee-types"
const ORG_COLLECTION_PERMISSIONS = "org-permissions"
const ORG_COLLECTION_AUTH = "auth"

const DEPARTMENT_COLLECTION_ROLES = "department-roles"

const EMPLOYEE_COLLECTION_PERSONAL = "employee-personal"
const EMPLOYEE_COLLECTION_PROFESSIONAL = "employee-professional"
const EMPLOYEE_COLLECTION_DOCUMENTS = "employee-documents"
const EMPLOYEE_COLLECTION_SUMMARY = "employee-summary"

func GetOrgCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_ORGANIZATION)
}

func GetOrgDepartmentCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_DEPARTMENTS)
}

func GetOrgDocumentCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_DOCUMENTS)
}

func GetDepRolesCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(DEPARTMENT_COLLECTION_ROLES)
}

func GetOrgPermissionsCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_PERMISSIONS)
}

func GetOrgEmployeeTypeCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_EMPLOYEE_TYPE)
}

func GetEmpPersonalCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(EMPLOYEE_COLLECTION_PERSONAL)
}

func GetEmpProfessionalCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(EMPLOYEE_COLLECTION_PROFESSIONAL)
}

func GetEmpDocumentCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(EMPLOYEE_COLLECTION_DOCUMENTS)
}

func GetEmpSummaryCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(EMPLOYEE_COLLECTION_SUMMARY)
}

func GetDocumentCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(EMPLOYEE_COLLECTION_DOCUMENTS)
}
