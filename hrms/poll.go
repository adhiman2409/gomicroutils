package hrms

import "time"

type PollOption struct {
	Option    string `bson:"option"`
	OptionId  string `bson:"option_id"`
	VoteCount string `bson:"vote_count"`
}

type Vote struct {
	EmployeeId string    `bson:"employee_id"`
	Option     string    `bson:"option"`
	OptionId   string    `bson:"option_id"`
	Date       time.Time `bson:"date"`
}

type Poll struct {
	ID          string       `bson:"id"`
	Title       string       `bson:"title"`
	Content     string       `bson:"content"`
	Options     []PollOption `bson:"options"`
	Votes       []Vote       `bson:"vote"`
	VoteCount   int          `bson:"vote_count"`
	StartDate   string       `bson:"start_date"`
	StartTime   string       `bson:"start_time"`
	EndDate     string       `bson:"end_date"`
	EndTime     string       `bson:"end_time"`
	PostedBy    string       `bson:"posted_by"`
	PostedById  string       `bson:"posted_by_id"`
	IsPublished bool         `bson:"is_published"`
	IsOpen      bool         `bson:"is_open"`
	IsClosed    bool         `bson:"is_closed"`
	ListDate    string       `bson:"list_date"`
	UnlistDate  string       `bson:"unlist_date"`
	CreatedAt   time.Time    `bson:"created_at"`
	UpdatedAt   time.Time    `bson:"updated_at"`
	Remarks     string       `bson:"remarks"`
}
