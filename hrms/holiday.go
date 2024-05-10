package hrms

type Holiday struct {
	SNo         string `json:"sno"`
	Date        string `json:"date"`
	Day         string `json:"day"`
	Holiday     string `json:"holiday"`
	HolidayType string `json:"holiday_type"`
}