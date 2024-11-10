package hrms

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StatusCount struct {
	Status string `bson:"status"`
	Count  int    `bson:"count"`
}

type CalculatedStats struct {
	Day             int           `bson:"day"`
	Month           int           `bson:"month"`
	Year            int           `bson:"year"`
	Weekday         string        `bson:"weekday"`
	TotalCount      int           `bson:"total_count"`
	StatusWithCount []StatusCount `bson:"status_with_count"`
	IsToday         bool          `bson:"is_today"`
	IsHoliday       bool          `bson:"is_holiday"`
}

type RecruiterStatsResponse struct {
	ID                   primitive.ObjectID `bson:"_id"`
	RecruiterId          string             `bson:"recruiter_id"`
	DailyStats           []CalculatedStats  `bson:"daily_stats"`
	MonthlyHourlyCount   [24]int            `bson:"monthly_hourly_count"`
	DailyHourlyCount     [24]int            `bson:"daily_hourly_count"`
	UpdatedAt            string             `bson:"updated_at"`
	StatusWiseTotalCount []StatusCount      `bson:"status_wise_total_count"`
}
