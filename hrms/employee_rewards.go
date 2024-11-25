package hrms

import "go.mongodb.org/mongo-driver/bson/primitive"

type Greeting struct {
	EmployeeId string `bson:"employee_id"`
	Name       string `bson:"name"`
	ImageURL   string `bson:"image_url"`
	Remarks    string `bson:"remarks"`
	CreatedAt  string `bson:"created_at"`
}

type RewardInfo struct {
	ID             primitive.ObjectID `bson:"_id"`
	EmployeeId     string             `bson:"employee_id"`
	Type           string             `bson:"type"`
	IsActive       bool               `bson:"is_active"`
	Status         string             `bson:"status"`
	AutoPublish    bool               `bson:"auto_publish"`
	AutoUnpublish  bool               `bson:"auto_unpublish"`
	PublishDay     int64              `bson:"publish_day"`
	PublishMonth   int64              `bson:"publish_month"`
	PublishYear    int64              `bson:"publish_year"`
	UnpublishDay   int64              `bson:"unpublish_day"`
	UnpublishMonth int64              `bson:"unpublish_month"`
	UnpublishYear  int64              `bson:"unpublish_year"`
	BadgeURL       string             `bson:"badge_url"`
	BadgeInfo      string             `bson:"badge_info"`
	ImageURL       string             `bson:"image_url"`
	Remarks        string             `bson:"remarks"`
	GreetingsCount int64              `bson:"Greetings_count"`
	Greetings      []Greeting         `bson:"Greetings"`
	CreatedAt      string             `bson:"created_at"`
	CreatedBy      string             `bson:"created_by"`
}
