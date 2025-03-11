package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Volunteer struct {
	ID                            primitive.ObjectID `bson:"_id"`
	VolunteerId                   string             `bson:"volunteer_id"`
	FullName                      string             `bson:"full_name"`
	FatherName                    string             `bson:"father_name"`
	Gender                        string             `bson:"gender"`
	DOB                           string             `bson:"dob"`
	PhoneNumber                   string             `bson:"phone_number"`
	PhoneNumberVerified           bool               `bson:"phone_number_verified"`
	EmailId                       string             `bson:"email_id"`
	EmailIdVerified               bool               `bson:"email_id_verified"`
	Profession                    string             `bson:"profession"`
	Address                       Address            `bson:"address"`
	PersonalIdType                string             `bson:"personal_id_type"`
	PersonalIdURL                 string             `bson:"personal_id_url"`
	IsPreviouslyVolunteered       bool               `bson:"is_previously_volunteered"`
	PreviousVolunteerExperiences  []Experience       `bson:"previous_volunteer_experiences"`
	VolunteerIntrestAreas         []string           `bson:"volunteer_intrest_areas"`
	IsVolunteerActive             bool               `bson:"is_volunteer_active"`
	IsVolunteerProfileApproved    bool               `bson:"is_volunteer_profile_approved"`
	IsVolunteerProfileRejected    bool               `bson:"is_volunteer_profile_rejected"`
	IsVolunteerProfileUnderReview bool               `bson:"is_volunteer_profile_under_review"`
	IsVolunteerProfileSuspended   bool               `bson:"is_volunteer_profile_suspended"`
	VolunteerAvailability         [7]Availability    `bson:"volunteer_availability"`
	JoiningDate                   string             `bson:"joining_date"`
	ReportingManagerName          string             `bson:"reporting_manager"`
	ReportingManagerId            string             `bson:"reporting_manager_id"`
	ReportingManagerEmail         string             `bson:"reporting_manager_email"`
	NewJoineeRewardId             string             `bson:"new_joinee_reward_id"`
	ImgURL                        string             `bson:"img_url"`
	IsProfileEditingLocked        bool               `bson:"is_profile_editing_locked"`
	RegexText                     string             `bson:"regex_text"`
	CreatedBy                     string             `bson:"created_by"`
	CreatedAt                     time.Time          `bson:"created_at"`
	UpdatedAt                     time.Time          `bson:"updated_at"`
}

type Availability struct {
	Day       string `bson:"day"`
	Morning   bool   `bson:"morning"`
	Afternoon bool   `bson:"afternoon"`
	Evening   bool   `bson:"evening"`
	FullDay   bool   `bson:"full_day"`
}

type VolunteerIdCounter struct {
	ID      primitive.ObjectID `bson:"_id"`
	Prefix  string             `bson:"prefix"`
	Counter int64              `bson:"counter"`
}
