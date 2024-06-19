package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InventoryStore struct {
	ID					primitive.ObjectID 			`bson:"_id"`
	DepartmentID		string					`bson:"department_id"`
	Department			string						`bson:"department"`
	Team				string						`bson:"team"`
	Project				string						`bson:"project"`
	Name				string						`bson:"name"`
	Type				string						`bson:"type"`
	Location 			string						`bson:"location"`
	CreatedAt			time.Time					`bson:"created_at"`
	Status				string						`bson:"status"`
	IsActive			bool						`bson:"is_active"`
	OpenTime			string						`bson:"open_time"`
	ClosingTime			string						`bson:"closing_time"`
	Buckets				int32						`bson:"buckets"`
	UpdatedAt			time.Time					`bson:"updated_at"`		
}