package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tenant struct {
	ID                primitive.ObjectID `bson:"_id"`
	EID               string             `bson:"eid"`
	Domain            string             `bson:"domain"`
	OrgName           string             `bson:"org_name"`
	EmailId           string             `bson:"email_id"`
	PhoneNumber       string             `bson:"phone_number"`
	Name              string             `bson:"name"`
	Department        string             `bson:"department"`
	Designation       string             `bson:"designation"`
	EmailVerified     bool               `bson:"email_verified"`
	UseGoogleOAuth    bool               `json:"use_google__oauth"`
	Password          string             `bson:"password"`
	Role              string             `bson:"role"`
	ImageURL          string             `bson:"image_url"`
	FirstLoginPending bool               `bson:"first_login_pending"`
	Status            string             `bson:"status"`
	RefreshToken      string             `bson:"refresh_token"`
	CreatedBy         string             `bson:"created_by"`
	CreatedAt         time.Time          `bson:"created_at"`
	UpdatedAt         time.Time          `bson:"updated_at"`
}
