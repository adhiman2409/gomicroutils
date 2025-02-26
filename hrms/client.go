package hrms

type Client struct {
	ClientName                  string `bson:"client_name"`
	ProjectName                 string `bson:"project_name"`
	IsInternalProject           bool   `bson:"is_internal_project"`
	ReportingManagerName        string `bson:"reporting_manager_name"`
	ReportingManagerEmailID     string `bson:"reporting_manager_email_id"`
	ReportingManagerPhoneNumber string `bson:"reporting_manager_phone_number"`
	IsClientBillable            bool   `bson:"is_client_billable"`
	ReleaseDate                 string `bson:"release_date"`
}
