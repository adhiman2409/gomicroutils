package hrms

import "go.mongodb.org/mongo-driver/bson/primitive"

type CongratulationList struct {
	EmployeeId string `bson:"employee_id"`
	Name       string `bson:"name"`
}

type CongratulationInfo struct {
	ID                  primitive.ObjectID   `bson:"_id"`
	EmployeeId          string               `bson:"employee_id"`
	CongratulationCount int64                `bson:"congratulation_count"`
	Congratulations     []CongratulationList `bson:"congratulations"`
}
