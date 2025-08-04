package hrms

import "go.mongodb.org/mongo-driver/bson/primitive"

type TaskState struct {
	ID          primitive.ObjectID `bson:"_id"`
	ProjectId   string             `bson:"project_id"`
	Name        string             `bson:"name"`
	Color       string             `bson:"color"`
	Description string             `bson:"description"`
	Icon        string             `bson:"icon"`
	Order       string             `bson:"order"`
	Default     bool               `bson:"default"`
}
