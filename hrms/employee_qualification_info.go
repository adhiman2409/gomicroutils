package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmpQualificationInfo struct {
	ID                           primitive.ObjectID `bson:"_id"`
	EID                          string             `bson:"eid"`
	SecondarySchoolName          string             `bson:"secondary_school_name"`
	SecondarySchoolStartYear     string             `bson:"secondary_school_start_year"`
	SecondarySchoolEndYear       string             `bson:"secondary_school_end_year"`
	SecondarySchoolBoard         string             `bson:"secondary_school_board"`
	SecondarySchoolRollNo        string             `bson:"secondary_school_rollno"`
	SecondarySchoolAddress       string             `bson:"secondary_school_address"`
	SecondarySchoolAttachmentURL string             `bson:"secondary_school_attachment_url"`
	HighSchoolName               string             `bson:"high_school_name"`
	HighSchoolStartYear          string             `bson:"high_school_start_year"`
	HighSchoolEndYear            string             `bson:"high_school_end_year"`
	HighSchoolStream             string             `bson:"high_school_stream"`
	HighSchoolBoard              string             `bson:"high_school_board"`
	HighSchoolRollNo             string             `bson:"high_school_rollno"`
	HighSchoolAddress            string             `bson:"high_school_address"`
	HighSchoolAttachmentURL      string             `bson:"high_school_attachment_url"`
	GraduationInstitute          string             `bson:"graduation_institute"`
	GraduationStartYear          string             `bson:"graduation_start_year"`
	GraduationEndYear            string             `bson:"graduation_end_year"`
	GraduationDegree             string             `bson:"graduation_degree"`
	GraduationMajor              string             `bson:"graduation_major"`
	GraduationRollNo             string             `bson:"graduation_rollno"`
	GraduationInstituteAddress   string             `bson:"graduation_institute_address"`
	GraduationAttachmentURL      string             `bson:"graduation_attachment_url"`
	PostGradInstitute            string             `bson:"postgrad_institute"`
	PostGradStartYear            string             `bson:"postgrad_start_year"`
	PostGradEndYear              string             `bson:"postgrad_end_year"`
	PostGradDegree               string             `bson:"postgrad_degree"`
	PostGradSpecialization       string             `bson:"postgrad_specialization"`
	PostGradRollNo               string             `bson:"postgrad_rollno"`
	PostGradInstituteAddress     string             `bson:"postgrad_institute_address"`
	PostGradAttachmentURL        string             `bson:"postgrad_attachment_url"`
	DoctorateUniversity          string             `bson:"doctorate_university"`
	DoctorateStartYear           string             `bson:"doctorate_start_year"`
	DoctorateEndYear             string             `bson:"doctorate_end_year"`
	DoctorateDegree              string             `bson:"doctorate_degree"`
	DoctorateSpecialization      string             `bson:"doctorate_specialization"`
	DoctorateRollNo              string             `bson:"doctorate_rollno"`
	DoctorateUniversityAddress   string             `bson:"doctorate_university_Address"`
	DoctorateAttachmentURL       string             `bson:"doctorate_attachment_url"`
	CreatedAt                    time.Time          `bson:"created_at"`
	UpdatedAt                    time.Time          `bson:"updated_at"`
}
