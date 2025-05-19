package hrms

import "go.mongodb.org/mongo-driver/bson/primitive"

type HourlyUserData struct {
	ActiveUserCount   int `bson:"active_user_count"`
	InactiveUserCount int `bson:"inactive_user_count"`
	TotalUserCount    int `bson:"total_user_count"`
}

type DailyUserData struct {
	Day            int              `bson:"day"`
	HourlyUserData []HourlyUserData `bson:"hourly_user_data"`
}

type MonthlyUserData struct {
	ID             primitive.ObjectID `bson:"_id"`
	TenantDomain   string             `bson:"tenant_domain"`
	Month          int                `bson:"month"`
	Year           int                `bson:"year"`
	DailyUserCount []DailyUserData    `bson:"daily_user_data"`
}
