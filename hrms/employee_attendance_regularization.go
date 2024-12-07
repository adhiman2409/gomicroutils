package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmpAttendanceRegularization struct {
	ID                  primitive.ObjectID `bson:"_id"`
	EID                 string             `bson:"eid"`
	Domain              string             `bson:"domain"`
	StartDate           string             `bson:"start_date"`
	EndDate             string             `bson:"end_date"`
	IsLOPRegularization bool               `bson:"is_lop_regularization"`
	CheckInTime         string             `bson:"check_in_time"`
	CheckOutTime        string             `bson:"check_out_time"`
	Reason              string             `bson:"reason"`
	Status              string             `bson:"status"`
	RegularizationCount int                `bson:"regularization_count"`
	CreatedAt           time.Time          `bson:"created_at"`
	UpdatedAt           time.Time          `bson:"updated_at"`
}
