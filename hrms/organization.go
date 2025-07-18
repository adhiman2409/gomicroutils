package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Organization struct {
	ID               primitive.ObjectID `bson:"_id"`
	Domain           string             `bson:"domain"`
	Name             string             `bson:"name"`
	AdminEmailId     string             `bson:"admin_email_id"`
	AdminPhoneNumber string             `bson:"admin_phone_number"`
	AdminName        string             `bson:"admin_name"`
	OrgEmailId       string             `bson:"org_email_id"`
	OrgPhoneNumber   string             `bson:"org_phone_number"`
	Status           string             `bson:"status"`
	Logo             string             `bson:"logo"`
	Address          string             `bson:"address"`
	ZipCode          string             `bson:"zip_code"`
	CountryCode      string             `bson:"country_code"`
	Country          string             `bson:"country"`
	TimeZone         string             `bson:"time_zone"`
	Description      string             `bson:"description"`
	SocialAccounts   []string           `bson:"social_accounts"`
	UseGoogleAuth    bool               `bson:"use_google_auth"`
	Website          string             `bson:"website"`
	CreatedBy        string             `bson:"created_by"`
	CreatedAt        time.Time          `bson:"created_at"`
	UpdatedAt        time.Time          `bson:"updated_at"`
}
