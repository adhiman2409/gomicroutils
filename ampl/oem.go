package ampl

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OEM struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	OEMId     string             `bson:"oem_id"`
	OEMName   string             `bson:"oem_name"`
	Domain    string             `bson:"domain"`
	EmailId   string             `bson:"email_id"`
	MobileNo  string             `bson:"phone"`
	Status    string             `bson:"status"`
	EID       string             `bson:"eid"`
	Password  string             `bson:"password"`
	Role      string             `bson:"role"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
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
