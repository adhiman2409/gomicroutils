package hrms

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Experience struct {
	StartYearAndMonth  string   `bson:"start_year_and_month"`
	EndYearAndMonth    string   `bson:"end_year_and_month"`
	ExperienceInMonths int      `bson:"experience_in_months"`
	CompanyName        string   `bson:"company_name"`
	Designation        string   `bson:"designation"`
	Skills             []string `bson:"skills"`
}

type Project struct {
	ProjectName      string   `bson:"project_name"`
	ProjectType      string   `bson:"project_type"`
	IsCurrentProject bool     `bson:"is_current_project"`
	IsBillable       bool     `bson:"is_billable"`
	StartDate        string   `bson:"start_date"`
	EndDate          string   `bson:"end_date"`
	Description      string   `bson:"description"`
	CompanyName      string   `bson:"company_name"`
	ClientName       string   `bson:"client_name"`
	Skills           []string `bson:"skills"`
	Domain           string   `bson:"domain"`
	URL              string   `bson:"url"`
}

type Achivement struct {
	Year        string `bson:"year"`
	Month       string `bson:"month"`
	Title       string `bson:"title"`
	Description string `bson:"description"`
	CompanyName string `bson:"company_name"`
	ClientName  string `bson:"client_name"`
	URL         string `bson:"url"`
}

type Award struct {
	Year        string `bson:"year"`
	Month       string `bson:"month"`
	Title       string `bson:"title"`
	Description string `bson:"description"`
	CompanyName string `bson:"company_name"`
	ClientName  string `bson:"client_name"`
	URL         string `bson:"url"`
}

type EmployeeTechInfo struct {
	ID                     primitive.ObjectID `bson:"_id"`
	EmployeeId             string             `bson:"employee_id"`
	About                  string             `bson:"about"`
	TotalExperinceInMonths int                `bson:"total_experince_in_months"`
	Skills                 []string           `bson:"skills"`
	Experiences            []Experience       `bson:"experinces"`
	Projects               []Project          `bson:"projects"`
	Achivements            []Achivement       `bson:"achivements"`
	Awards                 []Award            `bson:"awards"`
	SkillsRegex            string             `bson:"skills_regex"`
	DomainRegex            string             `bson:"domain_regex"`
	IsProfileEditingLocked bool               `bson:"is_profile_editing_locked"`
}
