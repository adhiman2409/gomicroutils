package hrms

import "go.mongodb.org/mongo-driver/bson/primitive"

type InventoryManager struct {
	ID 				primitive.ObjectID 	`bson:"_id"`
	Name 			string				`bson:"name"`
	Organisation 	string 				`bson:"organisation"`
	Email 			string 				`bson:"email"`
}