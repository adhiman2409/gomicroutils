package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Organization struct {
	ID               primitive.ObjectID `bson:"_id"`
	Domain           string             `bson:"domain"`
	Name             string             `bson:"name"`
	AdminEmailId     string             `bson:"admin_email_id"`
	AdminPhoneNumber string             `bson:"admin_phone_number"`
	AdminName        string             `bson:"admin_name"`
	OrgEmailId       string             `bson:"org_email_id"`
	OrgPhoneNumber   string             `bson:"org_phone_number"`
	Status           string             `bson:"status"`
	Logo             string             `bson:"logo"`
	Address          string             `bson:"address"`
	ZipCode          string             `bson:"zip_code"`
	CountryCode      string             `bson:"country_code"`
	Country          string             `bson:"country"`
	TimeZone         string             `bson:"time_zone"`
	Description      string             `bson:"description"`
	SocialAccounts   []string           `bson:"social_accounts"`
	Documents        []OrgDocument      `bson:"documents"`
	UseGoogleAuth    bool               `bson:"use_google_auth"`
	Website          string             `bson:"website"`
	CustomFields     []OrgCustomField   `bson:"custom_fields"`
	CreatedBy        string             `bson:"created_by"`
	CreatedAt        time.Time          `bson:"created_at"`
	UpdatedAt        time.Time          `bson:"updated_at"`
}

// OrgCustomField lets a user store arbitrary org attributes (e.g. GST number, PAN).
// Key should be stored normalized (e.g. lowercase, slug form) by the service layer so
// lookups via a "custom_fields.key" index stay consistent regardless of input casing.
type OrgCustomField struct {
	Key   string `bson:"key"`
	Value string `bson:"value"`
}

type OrgDocument struct {
	DocumentName string    `bson:"doc_name"`
	DocumentType string    `bson:"doc_type"`
	DocumentURL  string    `bson:"doc_url"`
	MetaTags     []string  `bson:"meta_tags"`
	UploadedBy   string    `bson:"uploaded_by"`
	CreatedAt    time.Time `bson:"created_at"`
}

type CompanyHeadInfo struct {
	EmployeeId   string    `bson:"employee_id"`
	Name         string    `bson:"name"`
	Email        string    `bson:"email"`
	Designation  string    `bson:"designation"`
	ProfileImage string    `bson:"profile_image"`
	CreatedAt    time.Time `bson:"created_at"`
}

type CompanyHeadsDoc struct {
	HRHead      CompanyHeadInfo `bson:"hr_head"`
	FinanceHead CompanyHeadInfo `bson:"finance_head"`
	AddedBy     string          `bson:"added_by"`
	UpdatedAt   time.Time       `bson:"updated_at"`
	CreatedAt   time.Time       `bson:"created_at"`
}

type OrgAddress struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Label        string             `bson:"label"`
	Country      string             `bson:"country"`
	State        string             `bson:"state"`
	City         string             `bson:"city"`
	AddressLine1 string             `bson:"address_line_1"`
	AddressLine2 string             `bson:"address_line_2"`
	Zipcode      string             `bson:"zipcode"`
	Landmark     string             `bson:"landmark"`
	Lattitude    float64            `bson:"latitude"`
	Longitude    float64            `bson:"longitude"`
	IsActive     bool               `bson:"is_active"`
	CreatedBy    string             `bson:"created_by"`
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
}

type OrganizationConfigLevel string

const (
	OrganizationConfigLevelGlobal  OrganizationConfigLevel = "global"
	OrganizationConfigLevelCountry OrganizationConfigLevel = "country"
	OrganizationConfigLevelState   OrganizationConfigLevel = "state"
	OrganizationConfigLevelCity    OrganizationConfigLevel = "city"
	OrganizationConfigLevelOffice  OrganizationConfigLevel = "office"
)

// GetAllConfigTypes returns all possible config levels in order of specificity, from most specific to least specific.
func GetAllConfigTypes() []OrganizationConfigLevel {
	return []OrganizationConfigLevel{
		OrganizationConfigLevelOffice,
		OrganizationConfigLevelCity,
		OrganizationConfigLevelState,
		OrganizationConfigLevelCountry,
		OrganizationConfigLevelGlobal,
	}
}

type OrganizationConfig struct {
	ID                              primitive.ObjectID      `bson:"_id"`
	TenantResignationConfigLevel    OrganizationConfigLevel `bson:"tenant_resignation_config"`
	TenantPayrollConfigLevel        OrganizationConfigLevel `bson:"tenant_payroll_config"`
	TenantLeaveConfigLevel          OrganizationConfigLevel `bson:"tenant_leave_config"`
	TenantAttendanceConfigLevel     OrganizationConfigLevel `bson:"tenant_attendance_config"`
	TenantWorkFromHomeConfigLevel   OrganizationConfigLevel `bson:"tenant_work_from_home_config"`
	TenantRegularizationConfigLevel OrganizationConfigLevel `bson:"tenant_regularization_config"`
	TenantExpenseConfigLevel        OrganizationConfigLevel `bson:"tenant_expense_config"`
	TenantCRMConfigLevel            OrganizationConfigLevel `bson:"tenant_crm_config"`
	TenantAppraisalConfigLevel      OrganizationConfigLevel `bson:"tenant_appraisal_config"`
	TenantTrainingConfigLevel       OrganizationConfigLevel `bson:"tenant_training_config"`
}
