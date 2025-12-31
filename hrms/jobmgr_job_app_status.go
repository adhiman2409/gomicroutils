package hrms

type JobAppStatus int

const (
	Pipeline JobAppStatus = iota + 1
	InternalSelect
	InternalReject
	ClientR1Scheduled
	ClientR2Scheduled
	ClientShared
	InternalScheduled
	Dropout
	Dropped
	DropoutSelect
	ClientR1Reject
	ClientR1Select
	ClientR2Select
	ClientR2Reject
	PositionClosed
	Joined
	NotJoined
	Duplicate
	OnHold
	Sourced
	InterviewSelect
	Viewed
	Offered
	Naukri
	Callback
	ScreenSelect
	ScreenReject
)

func (a JobAppStatus) String() string {
	return [...]string{"Pipeline", "Internal Select", "Internal Reject",
		"Client R1 Scheduled", "Client R2 Scheduled", "Client Shared", "Internal Scheduled",
		"Dropout", "Dropped", "Dropout Select", "Client R1 Reject", "Client R1 Select",
		"Client R2 Select", "Client R2 Reject", "Position Closed", "Joined", "Not Joined", "Duplicate", "On Hold", "Sourced",
		"Interview Select", "Viewed", "Offered", "Naukri", "Callback", "Screen Select", "Screen Reject"}[a-1]
}

func (a JobAppStatus) EnumIndex() int {
	return int(a)
}

func JobAppStatusEnumFromInt(a int) string {
	return [...]string{"Pipeline", "Internal Select", "Internal Reject",
		"Client R1 Scheduled", "Client R2 Scheduled", "Client Shared", "Internal Scheduled",
		"Dropout", "Dropped", "Dropout Select", "Client R1 Reject", "Client R1 Select",
		"Client R2 Select", "Client R2 Reject", "Position Closed", "Joined", "Not Joined", "Duplicate", "On Hold", "Sourced",
		"Interview Select", "Viewed", "Offered", "Naukri", "Callback", "Screen Select", "Screen Reject"}[a-1]
}

func GetAllJobAppStatus() []string {
	return []string{"Pipeline", "Internal Select", "Internal Reject",
		"Client R1 Scheduled", "Client R2 Scheduled", "Client Shared", "Internal Scheduled",
		"Dropout", "Dropped", "Dropout Select", "Client R1 Reject", "Client R1 Select",
		"Client R2 Select", "Client R2 Reject", "Position Closed", "Joined", "Not Joined", "Duplicate", "On Hold", "Sourced",
		"Interview Select", "Viewed", "Offered", "Naukri", "Callback", "Screen Select", "Screen Reject"}
}

type JobAppStatusIdx struct {
	State string
	Index int
}

func GetJobAppStatusNewIndex() []JobAppStatusIdx {
	return []JobAppStatusIdx{
		{State: "Naukri", Index: int(Naukri)},
		{State: "Pipeline", Index: int(Pipeline)},
		{State: "Viewed", Index: int(Viewed)},
		{State: "Callback", Index: int(Callback)},
		{State: "Sourced", Index: int(Sourced)},
		{State: "Internal Scheduled", Index: int(InternalScheduled)},
		{State: "Internal Select", Index: int(InternalSelect)},
		{State: "Interview Select", Index: int(InterviewSelect)},
		{State: "Client Shared", Index: int(ClientShared)},
		{State: "Client R1 Scheduled", Index: int(ClientR1Scheduled)},
		{State: "Client R1 Select", Index: int(ClientR1Select)},
		{State: "Client R2 Scheduled", Index: int(ClientR2Scheduled)},
		{State: "Client R2 Select", Index: int(ClientR2Select)},
		{State: "Internal Reject", Index: int(InternalReject)},
		{State: "Client R1 Reject", Index: int(ClientR1Reject)},
		{State: "Client R2 Reject", Index: int(ClientR2Reject)},
		{State: "Offered", Index: int(Offered)},
		{State: "Joined", Index: int(Joined)},
		{State: "Not Joined", Index: int(NotJoined)},
		{State: "Duplicate", Index: int(Duplicate)},
		{State: "On Hold", Index: int(OnHold)},
		{State: "Dropout", Index: int(Dropout)},
		{State: "Dropout Select", Index: int(DropoutSelect)},
		{State: "Dropped", Index: int(Dropped)},
		{State: "Position Closed", Index: int(PositionClosed)},
		{State: "Screen Select", Index: int(ScreenSelect)},
		{State: "Screen Reject", Index: int(ScreenReject)},
	}
}

func JobAppStatusEnumFromString(s string) JobAppStatus {
	if s == "Pipeline" {
		return Pipeline
	} else if s == "Viewed" {
		return Viewed
	} else if s == "Internal Select" {
		return InternalSelect
	} else if s == "Internal Reject" {
		return InternalReject
	} else if s == "Client R1 Scheduled" {
		return ClientR1Scheduled
	} else if s == "Client R2 Scheduled" {
		return ClientR2Scheduled
	} else if s == "Client Shared" {
		return ClientShared
	} else if s == "Internal Scheduled" {
		return InternalScheduled
	} else if s == "Dropout" {
		return Dropout
	} else if s == "Dropped" {
		return Dropped
	} else if s == "Dropout Select" {
		return DropoutSelect
	} else if s == "Client R1 Reject" {
		return ClientR1Reject
	} else if s == "Client R1 Select" {
		return ClientR1Select
	} else if s == "Client R2 Select" {
		return ClientR2Select
	} else if s == "Client R2 Reject" {
		return ClientR2Reject
	} else if s == "Position Closed" {
		return PositionClosed
	} else if s == "Joined" {
		return Joined
	} else if s == "Not Joined" {
		return NotJoined
	} else if s == "Duplicate" {
		return Duplicate
	} else if s == "On Hold" {
		return OnHold
	} else if s == "Sourced" {
		return Sourced
	} else if s == "Interview Select" {
		return InterviewSelect
	} else if s == "Offered" {
		return Offered
	} else if s == "Callback" {
		return Callback
	} else if s == "Screen Select" {
		return ScreenSelect
	} else if s == "Screen Reject" {
		return ScreenReject
	} else {
		return Naukri
	}
}
