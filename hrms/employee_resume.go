package hrms

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Experience struct {
	StartYearAndMonth string   `bson:"start_year_and_month"`
	EndYearAndMonth   string   `bson:"end_year_and_month"`
	DurationInMonths  int      `bson:"duration_in_months"`
	CompanyName       string   `bson:"company_name"`
	Designation       string   `bson:"designation"`
	Skills            []string `bson:"skills"`
}

type Project struct {
	ProjectName string   `bson:"project_name"`
	Description string   `bson:"description"`
	CompanyName string   `bson:"company_name"`
	ClientName  string   `bson:"client_name"`
	TechStack   []string `bson:"tech_stack"`
	URL         string   `bson:"url"`
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

type Qualification struct {
	Type      string `bson:"type"`
	Institute string `bson:"institute"`
	StartYear string `bson:"start_year"`
	EndYear   string `bson:"end_year"`
	Degree    string `bson:"degree"`
	Major     string `bson:"major"`
}

type EmployeeResume struct {
	ID                     primitive.ObjectID `bson:"_id"`
	EmployeeId             string             `bson:"employee_id"`
	FullName               string             `bson:"full_name"`
	DOB                    string             `bson:"dob"`
	PhoneNumber            string             `bson:"phone_number"`
	EmailId                string             `bson:"email_id"`
	TotalExperinceInMonths int                `bson:"total_experince_in_months"`
	Designation            string             `bson:"designation"`
	JoiningDate            string             `bson:"joining_date"`
	JoiningDay             int64              `bson:"joining_day"`
	JoiningMonth           int64              `bson:"joining_month"`
	JoiningYear            int64              `bson:"joining_year"`
	ImgURL                 string             `bson:"img_url"`
	TenureInDays           int                `bson:"tenure_in_days"`
	About                  string             `bson:"about"`
	Skills                 []string           `bson:"skills"`
	Experiences            []Experience       `bson:"experinces"`
	Projects               []Project          `bson:"projects"`
	Hobbies                []string           `bson:"hobbies"`
	Achivements            []Achivement       `bson:"achivements"`
	Awards                 []Award            `bson:"awards"`
	Qualifications         []Qualification    `bson:"qualifications"`
	SkillsRegex            string             `bson:"skills_regex"`
	TechStackRegex         string             `bson:"tech_stack_regex"`
}
