package ampl

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Showroom struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	ShowroomId     string             `bson:"showroom_id"`
	ShowroomName   string             `bson:"showroom_name"`
	OemId          string             `bson:"oem_id"`
	Domain         string             `bson:"domain"`
	Address        Address            `bson:"address"`
	GeoLocation    GeoPoint           `bson:"geo_location"`
	Contact        ContactInfo        `bson:"contact"`
	EmailId        string             `bson:"email_id"`
	MobileNo       string             `bson:"mobile_no"`
	Status         string             `bson:"status"`
	CreatedAt      time.Time          `bson:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at"`
	BankPreference BankPreference     `bson:"bank_preference"`
}

type Brand struct {
	Name string `bson:"name"`
	Logo string `bson:"logo"`
}

type Address struct {
	Line1    string `bson:"line1"`
	Line2    string `bson:"line2,omitempty"`
	Landmark string `bson:"landmark,omitempty"`
	City     string `bson:"city"`
	State    string `bson:"state"`
	Country  string `bson:"country"`
	Pincode  string `bson:"pincode"`
}

type GeoPoint struct {
	Latitude  float64 `bson:"lat"`
	Longitude float64 `bson:"lon"`
}

type ContactInfo struct {
	PrimaryAdminName  string `bson:"primary_admin_name"`
	PrimaryAdminPhone string `bson:"primary_admin_phone"`
	PrimaryAdminEmail string `bson:"primary_admin_email"`
}

type Employee struct {
	Id          primitive.ObjectID `bson:"_id"`
	EmployeeId  string             `bson:"employee_id"`
	ShowroomId  string             `bson:"showroom_id"`
	EmailId     string             `bson:"email_id"`
	FirstName   string             `bson:"first_name"`
	LastName    string             `bson:"last_name"`
	Phone       string             `bson:"phone"`
	Role        string             `bson:"role"`
	Designation string             `bson:"designation"`
	Department  string             `bson:"department"`
	ManagerName string             `bson:"manager_name"`
	ManagerEid  string             `bson:"manager_eid"`
	DOB         string             `bson:"dob"`
	IsActive    bool               `bson:"is_active"`
	Gender      string             `bson:"gender"`
	JoiningDate string             `bson:"joining_date"`
	CreatedBy   string             `bson:"created_by"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

type IdCounter struct {
	ID      primitive.ObjectID `bson:"_id"`
	Prefix  string             `bson:"prefix"`
	Counter int64              `bson:"counter"`
}

type CustomerLead struct {
	ID                     primitive.ObjectID `bson:"_id,omitempty"`
	ShowroomID             string             `bson:"showroom_id"`
	CustomerID             string             `bson:"customer_id"`
	FirstName              string             `bson:"first_name"`
	LastName               string             `bson:"last_name"`
	Phone                  string             `bson:"phone"`
	Email                  string             `bson:"email"`
	City                   string             `bson:"city"`
	FuelType               string             `bson:"fuel_type"`
	VehicleType            string             `bson:"vehicle_type"`
	MaxBudget              int64              `bson:"max_budget"`
	LeadSource             LeadSource         `bson:"lead_source"`
	Timeline               []TimelineEvent    `bson:"timeline"`
	AssociatedShowrooms    []string           `bson:"associated_showrooms"`
	AssignedToEid          string             `bson:"assigned_to_eid"`
	AssignedToName         string             `bson:"assigned_to_name"`
	TimeToBuyVehicleInDays string             `bson:"time_to_buy_vehicle_in_days"`
	Status                 string             `bson:"status"`
	CreatedAt              time.Time          `bson:"created_at"`
	UpdatedAt              time.Time          `bson:"updated_at"`
}

type LeadSource struct {
	Source              string `bson:"source"`
	EmployeeID          string `bson:"employee_id"`
	EmployeeName        string `bson:"employee_name"`
	CustomerID          string `bson:"customer_id"`
	CustomerName        string `bson:"customer_name"`
	SocialMediaPlatform string `bson:"social_media_platform"`
}

type TimelineEvent struct {
	Showroom       ShowroomBrief `bson:"showroom"`
	EventType      string        `bson:"event_type"`
	AttendedBy     string        `bson:"attended_by"`
	AttendedByName string        `bson:"attended_by_name"`
	AttendedByDept string        `bson:"attended_by_dept"`
	Remarks        string        `bson:"remarks"`
	CreatedAt      time.Time     `bson:"created_at"`
}

type BankPreference struct {
	Public  []string `bson:"public"`
	Private []string `bson:"private"`
	NBFC    []string `bson:"nbfc"`
}

type ShowroomBrief struct {
	ShowroomID string `bson:"showroom_id"`
	Name       string `bson:"name"`
	City       string `bson:"city"`
	State      string `bson:"state"`
}

type AssignmentRecord struct {
	ShowroomID   string    `bson:"showroom_id"`
	EID          string    `bson:"eid"`
	EmployeeName string    `bson:"employee_name"`
	CustomerId   string    `bson:"customer_id"`
	Time         time.Time `bson:"created_at"`
}

type SalesAssignmentDoc struct {
	ID         primitive.ObjectID                                  `bson:"_id,omitempty"`
	ShowroomID string                                              `bson:"showroom_id"`
	CreatedAt  time.Time                                           `bson:"created_at"`
	UpdatedAt  time.Time                                           `bson:"updated_at"`
	Years      map[string]map[string]map[string][]AssignmentRecord `bson:"years"`
}

type CustomerFinancialProfile struct {
	Id                     primitive.ObjectID `bson:"_id"`
	CustomerID             string             `bson:"customer_id"`
	CustomerAge            int                `bson:"customer_age"`
	PanCard                string             `bson:"pan_card"`
	AadharCard             string             `bson:"aadhar_card"`
	IncomeProof            string             `bson:"income_proof"`
	EmploymentType         string             `bson:"employment_type"`
	GrossMonthlyIncome     float64            `bson:"gross_monthly_income"`
	InHandMonthlyIncome    float64            `bson:"in_hand_monthly_income"`
	GrossAnnualIncome      float64            `bson:"gross_annual_income"`
	CIBILScore             int                `bson:"cibil_score"`
	CreditScore            int                `bson:"credit_score"`
	MaxDelayDaysInCIBIL    int                `bson:"max_delay_days_in_cibil"`
	HasOverdue             bool               `bson:"has_overdue"`
	ResidenceType          string             `bson:"residence_type"`
	ResidenceDocs          []Document         `bson:"residence_docs"`
	OngoingLoanCount       int                `bson:"ongoing_loan_count"`
	AvgMonthlyBalance      float64            `bson:"avg_monthly_balance"`
	FOIR                   float64            `bson:"foir"`
	CustomerState          string             `bson:"customer_state"`
	TotalOngoingMonthlyEMI float64            `bson:"total_ongoing_monthly_emi"`
	ConsentForIncomeProof  bool               `bson:"consent_for_income_proof"`
	YearsInCurrentJob      float64            `bson:"years_in_current_job"`
	IsCashCustomer         bool               `bson:"is_cash_customer"`
	VehicleType            string             `bson:"vehicle_type"`
	VehicleOnRoadPrice     float64            `bson:"vehicle_on_road_price"`
	VehicleExShowroomPrice float64            `bson:"vehicle_ex_showroom_price"`
	PreferredBankType      string             `bson:"preferred_bank_type"`
	PreferredBankName      string             `bson:"preferred_bank_name"`
	CreatedAt              time.Time          `bson:"created_at"`
	UpdatedAt              time.Time          `bson:"updated_at"`
}

type Document struct {
	Name string `bson:"name"`
	Path string `bson:"path"`
}
