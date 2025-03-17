package hrms

type ScholarshipState int

const (
	DataCollection ScholarshipState = iota + 1
	Verification
	PrimaryApproval
	SecondaryApproval
	ReadyToPublish
	PublishedForDonation
	DonationReceived
	Disbursement
	Closed
)

func (r ScholarshipState) String() string {
	return [...]string{"DataCollection", "Verification", "PrimaryApproval", "SecondaryApproval", "ReadyToPublish",
		"PublishedForDonation", "DonationReceived", "Disbursement", "Closed"}[r-1]
}

func (r ScholarshipState) EnumIndex() int {
	return int(r)
}

func GetAllScholarshipStates() []string {
	return []string{"DataCollection", "Verification", "PrimaryApproval", "SecondaryApproval", "ReadyToPublish",
		"PublishedForDonation", "DonationReceived", "Disbursement", "Closed"}
}

func ScholarshipStateEnumFromString(s string) ScholarshipState {
	if s == "DataCollection" {
		return DataCollection
	} else if s == "Verification" {
		return Verification
	} else if s == "PrimaryApproval" {
		return PrimaryApproval
	} else if s == "SecondaryApproval" {
		return SecondaryApproval
	} else if s == "ReadyToPublish" {
		return ReadyToPublish
	} else if s == "PublishedForDonation" {
		return PublishedForDonation
	} else if s == "DonationReceived" {
		return DonationReceived
	} else if s == "Disbursement" {
		return Disbursement
	} else {
		return Closed
	}
}
