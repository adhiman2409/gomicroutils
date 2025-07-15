package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type VerificationStatus string

const (
	VerificationStatusPending  VerificationStatus = "pending"
	VerificationStatusApproved VerificationStatus = "approved"
	VerificationStatusRejected VerificationStatus = "rejected"
)

type VerificationInfo struct {
	Status     VerificationStatus `bson:"status"`
	VerifiedBy string             `bson:"verified_by"`
	VerifiedAt time.Time          `bson:"verified_at"`
	Remarks    string             `bson:"remarks"`
}

type DocumentLabel struct {
	Id              primitive.ObjectID `bson:"_id"`
	Label           string             `bson:"label"`
	Description     string             `bson:"description"`
	AllowedExtTypes []string           `bson:"allowed_ext_types"`
	MaxDocSizeInMB  int64              `bson:"max_doc_size_in_mb"`
	CreatedBy       string             `bson:"created_by"`
	CreatedAt       time.Time          `bson:"created_at"`
}

type DocumentCategoryLabel struct {
	Id                   primitive.ObjectID `bson:"_id"`
	Label                string             `bson:"label"`
	Description          string             `bson:"description"`
	DocumentLabels       []DocumentLabel    `bson:"document_label"`
	TotalCount           int64              `bson:"total_count"`
	MandatoryCount       int64              `bson:"mandatory_count"`
	RequiredVerification bool               `bson:"required_verification"`
	VerificationInfo     VerificationInfo   `bson:"verification_info"`
	CreatedBy            string             `bson:"created_by"`
	CreatedAt            time.Time          `bson:"created_at"`
}

type DocumentListLabel struct {
	Id             primitive.ObjectID      `bson:"_id"`
	Label          string                  `bson:"label"`
	Description    string                  `bson:"description"`
	CategoryLabels []DocumentCategoryLabel `bson:"category_labels"`
	AccessLevel    string                  `bson:"access_level"`
	AccessList     []DocumentAccessList    `bson:"access_list"`
	CreatedBy      string                  `bson:"created_by"`
	CreatedAt      time.Time               `bson:"created_at"`
}

type Document struct {
	Label                string           `bson:"label"`
	Description          string           `bson:"description"`
	URL                  string           `bson:"url"`
	Status               string           `bson:"status"`
	AllowedExtTypes      []string         `bson:"allowed_ext_types"`
	MaxDocSizeInMB       int64            `bson:"max_doc_size_in_mb"`
	RequiredVerification bool             `bson:"required_verification"`
	VerificationInfo     VerificationInfo `bson:"verification_info"`
}

type DocumentCategory struct {
	Label                string           `bson:"label"`
	Description          string           `bson:"description"`
	Documents            []Document       `bson:"documents"`
	TotalCount           int64            `bson:"total_count"`
	MandatoryCount       int64            `bson:"mandatory_count"`
	RequiredVerification bool             `bson:"required_verification"`
	VerificationInfo     VerificationInfo `bson:"verification_info"`
}

type DocumentAccessList struct {
	Viewer   []string `bson:"viewer"`
	Editor   []string `bson:"editor"`
	Approver []string `bson:"approver"`
}

type DocumentList struct {
	Id           primitive.ObjectID   `bson:"_id"`
	Label        string               `bson:"label"`
	CategoryList []DocumentCategory   `bson:"category_list"`
	AccessLevel  string               `bson:"access_level"`
	AccessList   []DocumentAccessList `bson:"access_list"`
	CreatedBy    string               `bson:"created_by"`
	CreatedAt    time.Time            `bson:"created_at"`
}
