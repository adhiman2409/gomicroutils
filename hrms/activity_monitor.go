package hrms

import "go.mongodb.org/mongo-driver/bson/primitive"

type EmployeeMMonitorConfig struct {
	Id                              primitive.ObjectID `bson:"_id"`
	StartMonitoringAfterCheckInOnly bool               `bson:"startMonitoringAfterCheckInOnly"`
	IsCheckedIn                     bool               `bson:"isCheckedIn"`
	IsCheckedOut                    bool               `bson:"isCheckedOut"`
	MacAddress                      string             `bson:"macAddress"`
	UserId                          string             `bson:"UserId"`
	Name                            string             `bson:"Name"`
	Department                      string             `bson:"Department"`
	Email                           string             `bson:"Email"`
	Domain                          string             `bson:"Domain"`
}

type ActivityLogEntry struct {
	Id            primitive.ObjectID `bson:"_id"`
	Day           int                `bson:"day"`
	Month         int                `bson:"month"`
	Year          int                `bson:"year"`
	EmployeeID    string             `bson:"employee_id"`
	Timestamp     string             `bson:"timestamp"`
	UnixTimestamp int64              `bson:"unix_timestamp"`
	MacAddress    string             `bson:"mac_address"`
	Message       string             `bson:"message"`
	Hostname      string             `bson:"hostname"`
	User          string             `bson:"user"`
	Latitude      float64            `bson:"latitude,omitempty"`
	Longitude     float64            `bson:"longitude,omitempty"`
	HasLocation   bool               `bson:"has_location,omitempty"`
}
