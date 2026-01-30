package hrms

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrgEmployeeShifts struct {
	ID                primitive.ObjectID `bson:"_id"`
	ShiftType         ShiftType          `bson:"shift_type"`
	ShiftName         string             `bson:"shift_name"`
	Abbrivation       string             `bson:"abbrivation"`
	IsActive          bool               `bson:"is_active"`
	IsOverlappedShift bool               `bson:"is_overlapped_shift"`
	OrgCheckInTime    string             `bson:"org_check_in_time"`
	OrgCheckOutTime   string             `bson:"org_check_out_time"`
	Color             string             `bson:"color"`
	TextColor         string             `bson:"text_color"`
	Country           string             `bson:"country"`
	IsDefault         bool               `bson:"is_default"`
}
