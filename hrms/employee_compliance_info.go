package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// EmployeeComplianceInfo holds Indian statutory compliance attributes for an
// employee (Aadhaar linkage, ESIC insurance number, disability status) that are
// tracked separately from the salary/payroll structures.
//
// See collections.go for PAYROLL_COLLECTION_EMPLOYEE_COMPLIANCE_INFO and
// GetEmployeeComplianceInfoCollection, which back this struct's Mongo collection.
type EmployeeComplianceInfo struct {
	ID                     primitive.ObjectID `bson:"_id"`
	EmployeeID             string             `bson:"employee_id"`
	AadhaarLinked          bool               `bson:"aadhaar_linked"`
	ESICInsuranceNumber    string             `bson:"esic_insurance_number"`
	IsPersonWithDisability bool               `bson:"is_person_with_disability"`
	UpdatedAt              time.Time          `bson:"updated_at"`
}

// plaintextEmployeeComplianceInfo lists the top-level keys kept in the clear; every
// other field is packed into the encrypted __enc blob by
// marshalEncrypted/unmarshalEncrypted (payroll_encryption.go). ESICInsuranceNumber and
// IsPersonWithDisability are sensitive personal data, so both are deliberately left out of
// this allowlist, mirroring how PAN/UAN/bank details are excluded from
// plaintextSalaryStructure/plaintextSalaryConfig for the other payroll structs.
// (is_person_with_disability was plaintext up to v1.19.9-compliance-phase1.2; documents
// written by those versions still decode, since unmarshalEncrypted merges top-level keys with
// the blob, and are rewritten without the plaintext key on their next upsert.)
var plaintextEmployeeComplianceInfo = set(
	"_id", "employee_id", "aadhaar_linked", "updated_at",
)

func (c EmployeeComplianceInfo) MarshalBSON() ([]byte, error) {
	type alias EmployeeComplianceInfo
	return marshalEncrypted(alias(c), plaintextEmployeeComplianceInfo)
}

func (c *EmployeeComplianceInfo) UnmarshalBSON(data []byte) error {
	type alias EmployeeComplianceInfo
	return unmarshalEncrypted(data, (*alias)(c))
}
