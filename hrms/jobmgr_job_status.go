package hrms

type JobStatus int

const (
	Created JobStatus = iota + 1
	Published
	Unpublished
	Deleted
)

func (a JobStatus) String() string {
	return [...]string{"Created", "Published", "Unpublished", "Deleted"}[a-1]
}

func (a JobStatus) EnumIndex() int {
	return int(a)
}

func GetAllJobStatus() []string {
	return []string{"Created", "Published", "Unpublished", "Deleted"}
}

func JobStatusEnumFromString(s string) JobStatus {
	if s == "Created" {
		return Created
	} else if s == "Published" {
		return Published
	} else if s == "Unpublished" {
		return Unpublished
	} else {
		return Deleted
	}
}
