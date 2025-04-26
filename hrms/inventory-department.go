package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InventoryDepartment struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         string             `bson:"name"`
	Organisation string             `bson:"organisation"`
	OrgId        string             `bson:"org_id"`
	IsActive     bool               `bson:"is_active"`
	AdminId      string             `bson:"admin_id"`
	AdminName    string             `bson:"admin_name"`
	Description  string             `bson:"description"`
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
}
