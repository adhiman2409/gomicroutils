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
const ORG_COLLECTION_ATTENDANCE_CONF = "org-attendance-conf"
const ORG_COLLECTION_HOLIDAYS = "org-holidays"
const ORG_COLLECTION_LEAVES_CONF = "org-leaves-conf"
const ORG_COLLECTION_MARQUE = "org-marque"
const ORG_COLLECTION_POST = "org-post"
const ORG_COLLECTION_ROLES = "org-roles"
const ORG_COLLECTION_INVENTORY = "org-inventory"

const DEPARTMENT_COLLECTION_DESIGNATIONS = "department-designations"
const DEPARTMENT_COLLECTION_ATTENDANCE_CONF = "department-attendance-conf"

const EMPLOYEE_COLLECTION_PERSONAL = "employee-personal"
const EMPLOYEE_COLLECTION_PROFESSIONAL = "employee-professional"
const EMPLOYEE_COLLECTION_DOCUMENTS = "employee-documents"
const EMPLOYEE_COLLECTION_SUMMARY = "employee-summary"
const EMPLOYEE_COLLECTION_ATTENDANCE_CONF = "employee-attendance-conf"
const EMPLOYEE_COLLECTION_ATTENDANCE = "employee-attendance"
const EMPLOYEE_COLLECTION_DAILY_ATN_OBJECT = "employee-daily-atn-object"
const EMPLOYEE_COLLECTION_LEAVES = "employee-leaves"
const EMPLOYEE_COLLECTION_LEAVE_STATS = "employee-leave-stats"
const EMPLOYEE_COLLECTION_QUALIFICATION = "employee-qualification"
const EMPLOYEE_COLLECTION_CERTIFICATE_AND_TRAINING = "employee-cert-and-training"

func GetOrgCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_ORGANIZATION)
}

func GetDailyAttendanceObjectCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(EMPLOYEE_COLLECTION_DAILY_ATN_OBJECT)
}

func GetOrgAuthCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_AUTH)
}

func GetOrgLeaveConfCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_LEAVES_CONF)
}

func GetOrgDepartmentCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_DEPARTMENTS)
}

func GetOrgDocumentCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_DOCUMENTS)
}

func GetOrgHolidayCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_HOLIDAYS)
}

func GetOrgRolesCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_ROLES)
}

func GetOrgPermissionsCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_PERMISSIONS)
}

func GetOrgEmployeeTypeCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_EMPLOYEE_TYPE)
}

func GetOrgAttendanceConfCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_ATTENDANCE_CONF)
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

func GetEmpAttendanceConfCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(EMPLOYEE_COLLECTION_ATTENDANCE_CONF)
}

func GetEmpAttendanceCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(EMPLOYEE_COLLECTION_ATTENDANCE)
}

func GetEmpLeavesCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(EMPLOYEE_COLLECTION_LEAVES)
}

func GetDocumentCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(EMPLOYEE_COLLECTION_DOCUMENTS)
}

func GetDepDesignationCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(DEPARTMENT_COLLECTION_DESIGNATIONS)
}

func GetDepAttendanceConfCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(DEPARTMENT_COLLECTION_ATTENDANCE_CONF)
}

func GetOrgHolidaysCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_HOLIDAYS)
}

func GetEmpLeaveStatsCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(EMPLOYEE_COLLECTION_LEAVE_STATS)
}

func GetOrgMarqueCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_MARQUE)
}

func GetOrgPostCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_POST)
}

func GetEmpQualificationCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(EMPLOYEE_COLLECTION_QUALIFICATION)
}

func GetOrgInventoryCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_INVENTORY)
}

func GetEmpCertAndTrainingCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(EMPLOYEE_COLLECTION_CERTIFICATE_AND_TRAINING)
}
