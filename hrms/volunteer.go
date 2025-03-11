package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type VolunteerRequest struct {
	ID                           primitive.ObjectID `bson:"_id"`
	RequestId                    string             `bson:"request_id"`
	FullName                     string             `bson:"full_name"`
	FatherName                   string             `bson:"father_name"`
	Gender                       string             `bson:"gender"`
	DOB                          string             `bson:"dob"`
	PhoneNumber                  string             `bson:"phone_number"`
	PhoneNumberVerified          bool               `bson:"phone_number_verified"`
	EmailId                      string             `bson:"email_id"`
	EmailIdVerified              bool               `bson:"email_id_verified"`
	Profession                   string             `bson:"profession"`
	Address                      Address            `bson:"address"`
	PersonalIdNumber             string             `bson:"personal_id_number"`
	PersonalIdType               string             `bson:"personal_id_type"`
	PersonalIdURL                string             `bson:"personal_id_url"`
	IsPreviouslyVolunteered      bool               `bson:"is_previously_volunteered"`
	PreviousVolunteerExperiences []Experience       `bson:"previous_volunteer_experiences"`
	VolunteerIntrestAreas        []string           `bson:"volunteer_intrest_areas"`
	VolunteerAvailability        [7]Availability    `bson:"volunteer_availability"`
	RequestDate                  string             `bson:"request_date"`
	RequestStatus                string             `bson:"request_status"`
	ReviewerId                   string             `bson:"reviewer_id"`
	ReviewerName                 string             `bson:"reviewer_name"`
	ReviewerRemarks              string             `bson:"reviewer_remarks"`
	ImgURL                       string             `bson:"img_url"`
	RegexText                    string             `bson:"regex_text"`
	CreatedAt                    time.Time          `bson:"created_at"`
	UpdatedAt                    time.Time          `bson:"updated_at"`
}

type VolunteerInfo struct {
	ID                           primitive.ObjectID `bson:"_id"`
	VolunteerId                  string             `bson:"volunteer_id"`
	FullName                     string             `bson:"full_name"`
	FatherName                   string             `bson:"father_name"`
	About                        string             `bson:"about"`
	Gender                       string             `bson:"gender"`
	DOB                          string             `bson:"dob"`
	PhoneNumber                  string             `bson:"phone_number"`
	PhoneNumberVerified          bool               `bson:"phone_number_verified"`
	EmailId                      string             `bson:"email_id"`
	EmailIdVerified              bool               `bson:"email_id_verified"`
	Profession                   string             `bson:"profession"`
	Address                      Address            `bson:"address"`
	PersonalIdNumber             string             `bson:"personal_id_number"`
	PersonalIdType               string             `bson:"personal_id_type"`
	PersonalIdURL                string             `bson:"personal_id_url"`
	IsPreviouslyVolunteered      bool               `bson:"is_previously_volunteered"`
	PreviousVolunteerExperiences []Experience       `bson:"previous_volunteer_experiences"`
	VolunteerIntrestAreas        []string           `bson:"volunteer_intrest_areas"`
	VolunteerStatus              string             `bson:"volunteer_status"`
	Remarks                      string             `bson:"remarks"`
	VolunteerAvailability        [7]Availability    `bson:"volunteer_availability"`
	JoiningDate                  string             `bson:"joining_date"`
	ImgURL                       string             `bson:"img_url"`
	IsProfileEditingLocked       bool               `bson:"is_profile_editing_locked"`
	RegexText                    string             `bson:"regex_text"`
	CreatedBy                    string             `bson:"created_by"`
	CreatedAt                    time.Time          `bson:"created_at"`
	UpdatedAt                    time.Time          `bson:"updated_at"`
}

type Availability struct {
	Day       string `bson:"day"`
	Morning   bool   `bson:"morning"`
	Afternoon bool   `bson:"afternoon"`
	Evening   bool   `bson:"evening"`
	FullDay   bool   `bson:"full_day"`
}

type VolunteerIdCounter struct {
	ID                 primitive.ObjectID `bson:"_id"`
	IsTemporaryCounter bool               `bson:"is_temporary_counter"`
	Prefix             string             `bson:"prefix"`
	Counter            int64              `bson:"counter"`
}
