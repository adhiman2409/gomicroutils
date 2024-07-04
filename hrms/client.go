package hrms

type Client struct {
	ClientID    string `bson:"client_id"`
	Name        string `bson:"name"`
	EmailID     string `bson:"email_id"`
	PhoneNumber string `bson:"phone_number"`
}
