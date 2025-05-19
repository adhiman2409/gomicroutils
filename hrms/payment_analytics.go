package hrms

type DailyUserData struct {
	Day                    int   `bson:"day"`
	HourlyActiveUserCount  []int `bson:"user_count"`
	MaxActiveUserCount     int   `bson:"max_user_count"`
	MinActiveUserCount     int   `bson:"min_user_count"`
	TotalUserCount         int   `bson:"total_user_count"`
	TotalInactiveUserCount int   `bson:"total_inactive_user_count"`
}

type MonthlyUserData struct {
	TenantDomain   string          `bson:"tenant_domain"`
	Month          int             `bson:"month"`
	Year           int             `bson:"year"`
	DailyUserCount []DailyUserData `bson:"daily_user_count"`
}
