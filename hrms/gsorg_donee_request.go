package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StudentDoneeRequest struct {
	ID                         primitive.ObjectID `bson:"_id"`
	RequestId                  string             `bson:"request_id"`
	RequestStatus              string             `bson:"request_status"`
	StudentName                string             `bson:"student_name"`
	Gender                     string             `bson:"gender"`
	DOB                        string             `bson:"dob"`
	IsOrphan                   bool               `bson:"is_orphan"`
	IsSemiOrphan               bool               `bson:"is_semi_orphan"`
	IsParentDisabled           bool               `bson:"is_parent_disabled"`
	ParentDisability           string             `bson:"parent_disability"`
	GuardianName               string             `bson:"guardian_name"`
	GuardianPhoneNumber        string             `bson:"guardian_phone_number"`
	Address                    Address            `bson:"address"`
	GuardianRelation           string             `bson:"guardian_relation"`
	CaseReferredBy             string             `bson:"case_referred_by"`
	ReferrerPhoneNumber        string             `bson:"referrer_phone_number"`
	ReferrerEmail              string             `bson:"referrer_email"`
	ReferrerRemarks            string             `bson:"referrer_remarks"`
	CaseOfficerId              string             `bson:"case_officer_id"`
	CaseOfficerName            string             `bson:"case_officer_name"`
	CaseOfficerRemarks         string             `bson:"case_officer_remarks"`
	IsSocialMediaConsentGiven  bool               `bson:"is_social_media_consent_given"`
	IsRequestForCurrentSession bool               `bson:"is_request_for_current_session"`
	RequestedAmount            float64            `bson:"requested_amount"`
	RequestedItems             []string           `bson:"requested_items"`
	RequestFullfillmentDate    time.Time          `bson:"request_fullfillment_date"`
	CreatedAt                  time.Time          `bson:"created_at"`
	UpdatedAt                  time.Time          `bson:"updated_at"`
}
