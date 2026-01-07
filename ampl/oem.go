package ampl

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OEM struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	OEMId string             `bson:"oem_id"`

	OEMCode      string `bson:"oem_code"`
	ShowroomName string `bson:"showroom_name"`
	Brand        Brand  `bson:"brand"`
	Domain       string `bson:"domain"`

	Address     Address  `bson:"address"`
	GeoLocation GeoPoint `bson:"geo_location"`

	Contact  ContactInfo `bson:"contact"`
	Email    string      `bson:"email"`
	MobileNo string      `bson:"phone"`

	Status string `bson:"status"`

	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

type Brand struct {
	Name string `bson:"name"`
	Logo string `bson:"logo"`
}

type Address struct {
	Line1    string `bson:"line1"`
	Line2    string `bson:"line2,omitempty"`
	Landmark string `bson:"landmark,omitempty"`
	City     string `bson:"city"`
	State    string `bson:"state"`
	Country  string `bson:"country"`
	Pincode  string `bson:"pincode"`
}

type GeoPoint struct {
	Latitude  float64 `bson:"lat"`
	Longitude float64 `bson:"lon"`
}

type ContactInfo struct {
	PrimaryAdminName  string `bson:"primary_admin_name"`
	PrimaryAdminPhone string `bson:"primary_admin_phone"`
	PrimaryAdminEmail string `bson:"primary_admin_email"`
}
