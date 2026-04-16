package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmpDocumentInfo struct {
	ID                   primitive.ObjectID `bson:"_id"`
	EID                  string             `bson:"eid"`
	Domain               string             `bson:"domain"`
	Department           string             `bson:"department"`
	DocName              string             `bson:"doc_name"`
	DocCategory          string             `bson:"doc_category"`
	DocType              string             `bson:"doc_type"`
	DocPath              string             `bson:"doc_path"`
	IsVerified           string             `bson:"is_verified"`
	IsDocumentInfoLocked bool               `bson:"is_document_info_locked"`
	RegexText            string             `bson:"regex_text"`
	CreatedBy            string             `bson:"created_by"`
	CreatedAt            time.Time          `bson:"created_at"`
}

type DocumentOTP struct {
	ID           string    `bson:"id"`
	EID          string    `bson:"eid"`
	EmpName      string    `bson:"emp_name"`
	EmpEmail     string    `bson:"emp_email"`
	DocType      string    `bson:"doc_type"`
	DocCategory  string    `bson:"doc_category"`
	RequesterEID string    `bson:"requester_eid"`
	OTP          string    `bson:"otp"`
	ExpiresAt    time.Time `bson:"expires_at"`
	IsUsed       bool      `bson:"is_used"`
	CreatedAt    time.Time `bson:"created_at"`
}
