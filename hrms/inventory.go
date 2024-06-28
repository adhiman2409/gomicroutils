package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Inventory struct {
	ID 							primitive.ObjectID	`bson:"_id"`
	StoreID						string				`bson:"store_id"`
	BucketID					string				`bson:"bucket_id"`
	DepartmentID				string				`bson:"department_id"`
	ProductCode 				string				`bson:"product_code"`
	ProductName					string				`bson:"product_name"`
	ProductDescription  		string 				`bson:"product_description"`
	ProductOwnership		 	string				`bson:"product_ownership"`
	ProductCategory				string				`bson:"product_category"`
	AddedDate					time.Time			`bson:"added_date"`
	Price						int32				`bson:"price"`
	CurrentStock				int32				`bson:"current_stock"`
	MinimumStock				int32				`bson:"minimum_stock"`
	MaximumStock				int32				`bson:"maximum_stock"`
	ProductImages				[]string			`bson:"product_images"`
	Invoices					[]string			`bson:"invoices"`
	ProviderID					string				`bson:"provider_id"`
	ProviderName				string				`bson:"provider_name"`
	ProviderEmail				string				`bson:"provider_email"`
	CreatedAt					time.Time			`bson:"created_at"`
	UpdatedAt					time.Time			`bson:"updated_at"`
	CreatedBy					string				`bson:"created_by"`
	DepartmentName 				string				`bson:"department_name"`
	IsActive					bool				`bson:"is_active"`
	IsConsumeable				bool				`bson:"is_consumeable"`
	ShelfLifeInDays				int32				`bson:"shelf_life_in_days"`
	ExpiryDate					string				`bson:"expiry_date"`
	Type						string				`bson:"type"`
	Remarks						string				`bson:"remarks"`
}