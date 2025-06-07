package hrms

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserModel struct {
	ID        primitive.ObjectID `bson:"_id"`
	ObjectId  string             `bson:"_id"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	Phone     string             `bson:"phone,omitempty"`
	Status    string             `bson:"status"`
	CreatedBy string             `bson:"created_by"`
	CreatedAt string             `bson:"created_at"`
	UpdatedAt string             `bson:"updated_at"`
}
