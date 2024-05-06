package hrms

import "go.mongodb.org/mongo-driver/bson/primitive"

type OrgHoliday struct {
	ID          primitive.ObjectID `bson:"_id"`
	OrgName     string             `bson:"org_name"`
	Day         string             `bson:"day"`
	Month       string             `bson:"month"`
	Year        string             `bson:"year"`
	Type        string             `bson:"type"`
	WeekDay     string             `bson:"week_day"`
	Description string             `bson:"description"`
}
