package hrms

import "go.mongodb.org/mongo-driver/bson/primitive"

type Congratulation struct {
	EmployeeId string `bson:"employee_id"`
	Name       string `bson:"name"`
	ImageURL   string `bson:"image_url"`
	Remarks    string `bson:"remarks"`
	CreatedAt  string `bson:"created_at"`
}

type CongratulationInfo struct {
	ID                  primitive.ObjectID `bson:"_id"`
	Type                string             `bson:"type"`
	Status              string             `bson:"status"`
	Day                 int64              `bson:"day"`
	Month               int64              `bson:"month"`
	Year                int64              `bson:"year"`
	EmployeeId          string             `bson:"employee_id"`
	CongratulationCount int64              `bson:"congratulation_count"`
	Congratulations     []Congratulation   `bson:"congratulations"`
	CreatedAt           string             `bson:"created_at"`
}
