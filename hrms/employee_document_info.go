package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmpDocumentInfo struct {
	ID          primitive.ObjectID `bson:"_id"`
	EID         string             `bson:"eid"`
	Domain      string             `bson:"domain"`
	Department  string             `bson:"department"`
	DocName     string             `bson:"doc_name"`
	DocCategory string             `bson:"doc_category"`
	DocType     string             `bson:"doc_type"`
	DocPath     string             `bson:"doc_path"`
	RegexText   string             `bson:"regex_text"`
	CreatedBy   string             `bson:"created_by"`
	CreatedAt   time.Time          `bson:"created_at"`
}
