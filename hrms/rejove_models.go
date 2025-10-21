package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Practice struct {
	UserId string `bson:"user_id"`
	Cases  []Case `bson:"cases"`
}

type AdditionalPlans struct {
	ArchToBeTreated string    `bson:"arch_to_be_treated"` //Upper, Lower, Both
	ImpressionType  string    `bson:"impression_type"`    //Digital,Physical Impressions
	Remarks         string    `bson:"remarks"`
	Attachments     []string  `bson:"attachments"`
	VoiceNoteUrl    string    `bson:"voice_note_url"`
	CreatedAt       time.Time `bson:"created_at"`
	UpdatedAt       time.Time `bson:"updated_at"`
}

type Patient struct {
	PatientId      string    `bson:"patient_id"`
	DoctorId       string    `bson:"doctor_id"`
	EmailId        string    `bson:"email_id"`
	PhoneNumber    string    `bson:"phone_number"`
	FirstName      string    `bson:"first_name"`
	LastName       string    `bson:"last_name"`
	Gender         string    `bson:"gender"`
	DOB            time.Time `bson:"dob"`
	Age            int64     `bson:"age"`
	PatientType    string    `bson:"patient_type"` //Local or OutStation
	OutStationDate time.Time `bson:"out_station_date"`
	CreatedAt      time.Time `bson:"created_at"`
	UpdatedAt      time.Time `bson:"updated_at"`
}

type UserIdCounter struct {
	ID      primitive.ObjectID `bson:"_id"`
	Prefix  string             `bson:"prefix"`
	Counter int64              `bson:"counter"`
}

type CaseIdCounter struct {
	ID       primitive.ObjectID `bson:"_id"`
	DoctorId string             `bson:"doctor_id"`
	Prefix   string             `bson:"prefix"`
	Counter  int64              `bson:"counter"`
}

type PatientIdCounter struct {
	ID      primitive.ObjectID `bson:"_id"`
	Prefix  string             `bson:"prefix"`
	Counter int64              `bson:"counter"`
}

type DoctorAuthObject struct {
	ID                primitive.ObjectID `bson:"_id"`
	EID               string             `bson:"eid"`
	Domain            string             `bson:"domain"`
	OrgName           string             `bson:"org_name"`
	EmailId           string             `bson:"email_id"`
	PhoneNumber       string             `bson:"phone_number"`
	Name              string             `bson:"name"`
	Department        string             `bson:"department"`
	Designation       string             `bson:"designation"`
	EmailVerified     bool               `bson:"email_verified"`
	UseGoogleOAuth    bool               `bson:"use_google_oauth"`
	Password          string             `bson:"password"`
	Role              string             `bson:"role"`
	ImageURL          string             `bson:"image_url"`
	FirstLoginPending bool               `bson:"first_login_pending"`
	Status            string             `bson:"status"`
	RefreshToken      string             `bson:"refresh_token"`
	CreatedBy         string             `bson:"created_by"`
	CreatedAt         time.Time          `bson:"created_at"`
	UpdatedAt         time.Time          `bson:"updated_at"`
}

type DoctorInfo struct {
	ID                primitive.ObjectID `bson:"_id"`
	DoctorId          string             `bson:"doctor_id"`
	Password          string             `bson:"password"`
	Designation       string             `bson:"designation"`
	FirstLoginPending bool               `bson:"first_login_pending"`
	OrgName           string             `bson:"org_name"`
	EmailId           string             `bson:"email_id"`
	PhoneNumber       string             `bson:"phone_number"`
	Department        string             `bson:"department"`
	OfficeName        string             `bson:"office_name"`
	FirstName         string             `bson:"first_name"`
	LastName          string             `bson:"last_name"`
	Specialization    string             `bson:"specialization"`
	EmailVerified     bool               `bson:"email_verified"`
	Status            string             `bson:"status"` //Active and Inactive
	BillingAddress    Address            `bson:"billing_address"`
	ShippingAddress   Address            `bson:"shipping_address"`
	RefreshToken      string             `bson:"refresh_token"`
	ImageURL          string             `bson:"image_url"`
	GSTNumber         string             `bson:"gst_number"`
	Role              string             `bson:"role"`
	CreatedBy         string             `bson:"created_by"`
	CreatedAt         time.Time          `bson:"created_at"`
	UpdatedAt         time.Time          `bson:"updated_at"`
}

type Case struct {
	CaseId                    string                    `bson:"case_id"`
	DoctorId                  string                    `bson:"doctor_id"`
	PatientId                 string                    `bson:"patient_id"`
	CaseType                  string                    `bson:"case_type"` //Aligner, RetainerOnly
	PlanType                  string                    `bson:"plan_type"` //NewCase, MidcourseCorrection, Refinements
	IsSubmitted               bool                      `bson:"is_submitted"`
	SubmissionVoiceNoteUrl    string                    `bson:"submission_voice_note_url"`
	CasePreferences           CasePreference            `bson:"case_preferences"`
	ToothMovementRestrictions ToothMovementRestrictions `bson:"tooth_movement_restrictions"`
	CaseAttachments           CaseAttachments           `bson:"case_attachments"`
	CaseStatus                CaseStatus                `bson:"case_status"`
	MidCourseCorrection       []AdditionalPlans         `bson:"mid_course_correction"`
	Refinements               []AdditionalPlans         `bson:"refinements"`
	AllignerTacking           []AdditionalPlans         `bson:"alligner_tacking"`
	Remarks                   string                    `bson:"remarks"`
	CreatedAt                 time.Time                 `bson:"created_at"`
	UpdatedAt                 time.Time                 `bson:"updated_at"`
}

type CasePreference struct {
	ArchToBeTreated        string `bson:"arch_to_be_treated"`       //Upper, Lower, Both
	ArchCorrection         string `bson:"arch_correction"`          //Anterior(Social Six),Full
	OverJet                string `bson:"over_jet"`                 //Maintain,Increase,Decrease
	OverBite               string `bson:"over_bite"`                //Maintain,Increase,Decrease
	MidLine                string `bson:"mid_line"`                 //Maintain,Improve
	CanineRelation         string `bson:"canine_relation"`          //Maintain,Improve,None
	MolarRelation          string `bson:"molar_relation"`           //Maintain,Improve,None
	SpaceAlterations       string `bson:"space_alterations"`        //CloseAllSpaces,CreateSpaceFor{number},LeaveSpaceDistalTo{number},None
	SpaceGainingPreference string `bson:"space_gaining_preference"` //IPR,Extraction,Expansion,None
	Upper                  string `bson:"upper"`                    //Anterior, Posterior, Both
	Lower                  string `bson:"lower"`                    //Anterior, Posterior, Both
}

type ToothMovementRestrictions struct {
	MovementType string    `bson:"movement_type"`
	Teeth        [32]Tooth `bson:"teeth"`
}

type Tooth struct {
	HasCrown      bool `bson:"has_crown"`
	HasImplant    bool `bson:"has_implant"`
	HasAttachment bool `bson:"has_attachment"`
}

type CaseAttachments struct {
	//IntraOral
	LeftOcclusion      string `bson:"left_occlusion"`
	FrontOcclusion     string `bson:"front_occlusion"`
	RightOcclusion     string `bson:"right_occlusion"`
	MaxillaryOcclusal  string `bson:"maxillary_occlusal"`
	MandibularOcclusal string `bson:"mandibular_occlusal"`
	//ExtraOral
	Profile        string `bson:"profile"`
	FrontOral      string `bson:"front_oral"`
	FrontalDynamic string `bson:"frontal_dynamic"` //Smile
	//RadioGraphs
	OPG                string   `bson:"opg"`
	LateralCephalogram string   `bson:"lateral_cephalogram"`
	OtherRecords       []string `bson:"other_records"`
	//STL Files
	STLFiles []string `bson:"stl_files"`
}

type AllProducts struct {
	Submitted                     int64 `bson:"submitted"`
	UnderInspection               int64 `bson:"under_inspection"`
	PendingCaseRecords            int64 `bson:"pending_case_records"`
	VirtualSetupInProgress        int64 `bson:"virtual_setup_in_progress"`
	ApprovedVirtualSetup          int64 `bson:"approved_virtual_setup"`
	AligneerFabricationInProgress int64 `bson:"aligneer_fabrication_in_progress"`
	OngoinCase                    int64 `bson:"ongoing_case"`
	Delivered                     int64 `bson:"delivered"`
	Retainers                     int64 `bson:"retainers"`
	Cancelled                     int64 `bson:"cancelled"`
}
