package hrms

import "time"

// DepartmentRecord tracks one department assignment in an employee's service history.
// EndDate zero value means the assignment is currently active.
type DepartmentRecord struct {
	Department string    `bson:"department"`
	StartDate  time.Time `bson:"start_date"`
	EndDate    time.Time `bson:"end_date"`
	Remarks    string    `bson:"remarks"`
	ChangedBy  string    `bson:"changed_by"`
}

// DesignationRecord tracks one designation assignment in an employee's service history.
// EndDate zero value means the assignment is currently active.
type DesignationRecord struct {
	Designation string    `bson:"designation"`
	StartDate   time.Time `bson:"start_date"`
	IsPromoted  bool      `bson:"is_promoted"`
	EndDate     time.Time `bson:"end_date"`
	Remarks     string    `bson:"remarks"`
	ChangedBy   string    `bson:"changed_by"`
}

// OfficeRecord tracks one office posting with geo-coordinates.
// EndDate zero value means the posting is currently active.
type OfficeRecord struct {
	OfficeName   string    `bson:"office_name"`
	Lat          float64   `bson:"lat"`
	Lng          float64   `bson:"lng"`
	IsTransfered bool      `bson:"is_transfered"`
	StartDate    time.Time `bson:"start_date"`
	EndDate      time.Time `bson:"end_date"`
	Remarks      string    `bson:"remarks"`
	ChangedBy    string    `bson:"changed_by"`
}

// WorkLocationRecord tracks one work location with geo-coordinates.
// EndDate zero value means the location is currently active.
type WorkLocationRecord struct {
	LocationName string    `bson:"location_name"`
	Lat          float64   `bson:"lat"`
	Lng          float64   `bson:"lng"`
	StartDate    time.Time `bson:"start_date"`
	EndDate      time.Time `bson:"end_date"`
	Remarks      string    `bson:"remarks"`
	ChangedBy    string    `bson:"changed_by"`
}

// DeputationRecord captures one deputation period to another department, office, or work location.
// EndDate zero value means the deputation is currently active.
// DurationDays is stored explicitly so queries can filter on duration without computing it.
type DeputationRecord struct {
	Department   string    `bson:"department"`
	OfficeName   string    `bson:"office_name"`
	OfficeLat    float64   `bson:"office_lat"`
	OfficeLng    float64   `bson:"office_lng"`
	WorkLocation string    `bson:"work_location"`
	WorkLat      float64   `bson:"work_lat"`
	WorkLng      float64   `bson:"work_lng"`
	StartDate    time.Time `bson:"start_date"`
	EndDate      time.Time `bson:"end_date"`
	DurationDays int       `bson:"duration_days"`
	IsActive     bool      `bson:"is_active"`
	OrderedBy    string    `bson:"ordered_by"`
	Remarks      string    `bson:"remarks"`
}
