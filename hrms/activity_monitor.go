package hrms

import "go.mongodb.org/mongo-driver/bson/primitive"

type EmployeeMMonitorConfig struct {
	Id                              primitive.ObjectID `bson:"_id"`
	StartMonitoringAfterCheckInOnly bool               `bson:"start_monitoring_after_check_in_only"`
	IsCheckedIn                     bool               `bson:"is_checked_in"`
	IsCheckedOut                    bool               `bson:"is_checked_out"`
	MacAddress                      string             `bson:"mac_address"`
	UserId                          string             `bson:"user_id"`
	Name                            string             `bson:"name"`
	Department                      string             `bson:"department"`
	Email                           string             `bson:"email"`
	Domain                          string             `bson:"domain"`
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
