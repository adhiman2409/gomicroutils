package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Organization struct {
	ID             primitive.ObjectID `bson:"_id"`
	Domain         string             `bson:"domain"`
	Name           string             `bson:"name"`
	AdminId        string             `bson:"admin_id"`
	EmailId        string             `bson:"email_id"`
	PhoneNumber    string             `bson:"phone_number"`
	Status         string             `bson:"status"`
	Logo           string             `bson:"logo"`
	Address        string             `bson:"address"`
	ZipCode        string             `bson:"zip_code"`
	CountryCode    string             `bson:"country_code"`
	Description    string             `bson:"description"`
	SocialAccounts []string           `bson:"social_accounts"`
	Website        string             `bson:"website"`
	CreatedBy      string             `bson:"created_by"`
	CreatedAt      time.Time          `bson:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at"`
}
