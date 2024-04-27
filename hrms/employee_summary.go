package hrms

type EmployeeSummary struct {
	EmpId           string `bson:"emp_id"`
	Name            string `bson:"name"`
	Department      string `bson:"department"`
	Designation     string `bson:"designation"`
	Status          string `bson:"status"`
	ProfileImageURL string `bson:"profile_image_url"`
}
