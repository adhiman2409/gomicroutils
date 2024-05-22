package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InventoryStore struct {
	ID					primitive.ObjectID 			`bson:"_id"`
	ManagerID			string						`bson:"manager_id"`
	ManagerName			string						`bson:"manager_name"`
	ManagerEmail		string						`bson:"manager_email"`
	Department			string						`bson:"department"`
	Team				string						`bson:"team"`
	Project				string						`bson:"project"`
	Name				string						`bson:"name"`
	Type				string						`bson:"type"`
	Location 			string						`bson:"location"`
	CreatedAt			time.Time					`bson:"created_at"`
	Status				string						`bson:"status"`
	IsActive			bool						`bson:"is_active"`
	OpenTime			time.Time					`bson:"open_time"`
	ClosingTime			time.Time					`bson:"closing_time"`
	Buckets				int32						`bson:"buckets"`
	UpdatedAt			time.Time					`bson:"updated_at"`		
}