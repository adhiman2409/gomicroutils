package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AssignAsset struct {
	ID               primitive.ObjectID `bson:"_id"`
	BucketID         string             `bson:"bucket_id"`
	StoreID          string             `bson:"store_id"`
	BucketName       string             `bson:"bucket_name"`
	StoreName        string             `bson:"store_name"`
	EID              string             `bson:"eid"`
	AssetImages      string             `bson:"asset_images"`
	AssetDocs        string             `bson:"asset_docs"`
	AssetID          string             `bson:"asset_id"`
	AssetName        string             `bson:"asset_name"`
	AssetDescription string             `bson:"asset_description"`
	AssetPrice       string             `bson:"asset_price"`
	TimeAllocated    string             `bson:"time_allocated"`
	DateOfExpiry     string             `bson:"date_of_expiry"`
	AssetLocation    string             `bson:"asset_location"`
	AssetCategory    string             `bson:"asset_category"`
	AssetType        string             `bson:"asset_type"`
	AssetCondition   string             `bson:"asset_condition"`
	AssetStatus      string             `bson:"asset_status"`
	StockConsumed    int32              `bson:"stock_consumed"`
	Country          string             `bson:"country"`
	State            string             `bson:"state"`
	CreatedAt        time.Time          `bson:"created_at"`
	UpdatedAt        time.Time          `bson:"updated_at"`
}
