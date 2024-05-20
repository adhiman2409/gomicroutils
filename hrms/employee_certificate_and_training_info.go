package hrms

import "go.mongodb.org/mongo-driver/bson/primitive"

type EmpCertificateAndTrainingInfo struct {
	ID          primitive.ObjectID `bson:"_id"`
	EID         string             `bson:"eid"`
	Certificate []Certificate      `bson:"certificate"`
	Training    []Training         `bson:"training"`
}

type Certificate struct {
	CertificateNumber   string `bson:"certificate_number"`
	CertificateName     string `bson:"certificate_name"`
	IssuingOrganization string `bson:"issuing_organization"`
	IssueDate           string `bson:"issue_date"`
	Expiry              string `bson:"expiry"`
	Description         string `bson:"description"`
	AttachmentURL       string `bson:"attachment_url"`
}

type Training struct {
	TrainingProgramName string `bson:"training_program_name"`
	TrainingProvider    string `bson:"training_provider"`
	StartDate           string `bson:"start_date"`
	EndDate             string `bson:"end_date"`
	TrainingDuration    string `bson:"training_duration"`
	CertificateAwarded  bool   `bson:"certificate_awarded"`
	CertificateNumber   string `bson:"certificate_number"`
	Description         string `bson:"description"`
	AttachmentURL       string `bson:"attachment_url"`
}
