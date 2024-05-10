package hrms

import "go.mongodb.org/mongo-driver/bson/primitive"

type Marque struct {
	ID         primitive.ObjectID `bson:"_id"`
	MarqueType string             `bson:"marque_type"`
	Message    string             `bson:"message"`
	DateFrom   string             `bson:"date_from"`
	DateTo     string             `bson:"date_to"`
}
