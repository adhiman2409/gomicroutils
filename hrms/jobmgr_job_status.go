package hrms

type JobPublishStatus int

const (
	Created JobPublishStatus = iota + 1
	Published
	Unpublished
	Deleted
)

func (a JobPublishStatus) String() string {
	return [...]string{"Created", "Published", "Unpublished", "Deleted"}[a-1]
}

func (a JobPublishStatus) EnumIndex() int {
	return int(a)
}

func GetAllJobPublishStatus() []string {
	return []string{"Created", "Published", "Unpublished", "Deleted"}
}

func JobPublishStatusEnumFromString(s string) JobPublishStatus {
	switch s {
	case "Created":
		return Created
	case "Published":
		return Published
	case "Unpublished":
		return Unpublished
	default:
		return Deleted
	}
}

type JobStatus int

const (
	JobOpened JobStatus = iota + 1
	JobClosed
	JobPartiallyClosed
	JobOnHold
)

func (a JobStatus) String() string {
	return [...]string{"JobOpened", "JobOpened", "JobClosed", "JobPartiallyClosed", "JobOnHold"}[a-1]
}

func (a JobStatus) EnumIndex() int {
	return int(a)
}

func GetAllJobStatus() []string {
	return []string{"JobOpened", "JobClosed", "JobPartiallyClosed", "JobOnHold"}
}

func JobStatusEnumFromString(s string) JobStatus {
	switch s {
	case "JobOpened":
		return JobOpened
	case "JobClosed":
		return JobClosed
	case "JobPartiallyClosed":
		return JobPartiallyClosed
	default:
		return JobOnHold
	}
}
