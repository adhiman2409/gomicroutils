package hrms

import (
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// PAYROLL_COLLECTION_EMPLOYEE_COMPLIANCE_INFO is the Mongo collection name backing
// EmployeeComplianceInfo documents, following the "payroll-employee-*" naming
// convention used by the other payroll collections in this package (see
// PAYROLL_COLLECTION_EMPLOYEE_PAYROLL_MASTER in collections.go).
const PAYROLL_COLLECTION_EMPLOYEE_COMPLIANCE_INFO = "payroll-employee-compliance-info"

// EmployeeComplianceInfo holds Indian statutory compliance attributes for an
// employee (Aadhaar linkage, ESIC insurance number, disability status) that are
// tracked separately from the salary/payroll structures.
type EmployeeComplianceInfo struct {
	ID                     primitive.ObjectID `bson:"_id"`
	EmployeeID             string             `bson:"employee_id"`
	AadhaarLinked          bool               `bson:"aadhaar_linked"`
	ESICInsuranceNumber    string             `bson:"esic_insurance_number"`
	IsPersonWithDisability bool               `bson:"is_person_with_disability"`
	UpdatedAt              time.Time          `bson:"updated_at"`
}

// GetEmployeeComplianceInfoCollection returns the per-tenant collection handle for
// EmployeeComplianceInfo documents, mirroring the shape of
// GetPayrollEmployeePayrollMasterCollection (collections.go).
func GetEmployeeComplianceInfoCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(PAYROLL_COLLECTION_EMPLOYEE_COMPLIANCE_INFO)
}

// plaintextEmployeeComplianceInfo lists the top-level keys kept in the clear; every
// other field is packed into the encrypted __enc blob by
// marshalEncrypted/unmarshalEncrypted (payroll_encryption.go). ESICInsuranceNumber is
// the only sensitive field on this struct, so it is deliberately left out of this
// allowlist, mirroring how PAN/UAN/bank details are excluded from
// plaintextSalaryStructure/plaintextSalaryConfig for the other payroll structs.
var plaintextEmployeeComplianceInfo = set(
	"_id", "employee_id", "aadhaar_linked", "is_person_with_disability", "updated_at",
)

func (c EmployeeComplianceInfo) MarshalBSON() ([]byte, error) {
	type alias EmployeeComplianceInfo
	return marshalEncrypted(alias(c), plaintextEmployeeComplianceInfo)
}

func (c *EmployeeComplianceInfo) UnmarshalBSON(data []byte) error {
	type alias EmployeeComplianceInfo
	return unmarshalEncrypted(data, (*alias)(c))
}
