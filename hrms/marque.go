package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Marque struct {
	ID         primitive.ObjectID `bson:"_id"`
	MarqueType string             `bson:"marque_type"`
	Message    string             `bson:"message"`
	DateFrom   string             `bson:"date_from"`
	DateTo     string             `bson:"date_to"`
	IsPrivate  bool               `bson:"is_private"`
	IsRead     bool               `bson:"is_read"`
	EmployeeID string             `bson:"employee_id,omitempty"`
	CreatedAt  time.Time          `bson:"created_at,omitempty"`
}
