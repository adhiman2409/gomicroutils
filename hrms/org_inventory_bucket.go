package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)
	
type StoreBucket struct {			
	ID			 		primitive.ObjectID 		`bson:"_id"`
	StoreID				string					`bson:"store_id"`
	DepartmentID		string					`bson:"department_id"`
	Department			string					`bson:"department"`
	StoreName			string					`bson:"store_name"`
	Category			string					`bson:"category"`
	Description			string					`bson:"description"`
	Capacity			int32					`bson:"capacity"`
	NoOfInventory		int32					`bson:"no_of_inventory":`
	Status				string					`bson:"status"`
	IsActive			bool					`bson:"is_active"`
	CreatedAt			time.Time				`bson:"created_at"`
	UpdatedAt			time.Time				`bson:"updated_at"`		
}