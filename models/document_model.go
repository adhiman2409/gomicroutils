package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DocTracking struct {
	OperationType     string `bson:"operation_type"`
	PerformedByUserId string `bson:"performed_by_user_id"`
	PerformedAt       int64  `bson:"performed_at"`
}

type DocumentEntry struct {
	ID                  primitive.ObjectID `bson:"_id"`
	RequestType         string             `bson:"request_type"`
	InputFileName       string             `bson:"input_file_name"`
	InputFilePath       string             `bson:"input_file_path"`
	OutputFileName      string             `bson:"output_file_name"`
	OutputFilePath      string             `bson:"output_file_path"`
	NumberOfInputFiles  int64              `bson:"number_of_input_files"`
	InputFileList       []string           `bson:"input_file_list"`
	NumberOfOutPutFiles int64              `bson:"number_of_output_files"`
	OutputFileList      []string           `bson:"output_file_list"`
	IsProcessed         bool               `bson:"is_processed"`
	UserId              string             `bson:"user_id"`
	Year                string             `bson:"year"`
	Month               string             `bson:"month"`
	Day                 string             `bson:"day"`
	MarkedDeleted       bool               `bson:"marked_deleted"`
	OriginalSize        int64              `bson:"original_size_in_bytes"`
	OptimizedSize       int64              `bson:"optimized_size_in_bytes"`
	CompressionRatio    float64            `bson:"compression_ratio"`
	ProcessingTime      float64            `bson:"processing_time_in_ms"`
	ImageType           string             `bson:"image_type"`
	DocTrackingList     []DocTracking      `bson:"doc_tracking_list"`
	CreatedAt           int64              `bson:"created_at"`
}
