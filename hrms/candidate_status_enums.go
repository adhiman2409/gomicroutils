package hrms

type CandidateStatus int

const (
	AwaitingDocuments CandidateStatus = iota + 1
	DocumentsSubmitted
	DocumentsVerified
	OfferFinalized
	AwaitingOfferApproval
	OfferApproved
	AwaitingOfferAcceptance
	OfferAccepted
	OfferRejected
	OfferExpired
	OfferWithdrawn
	AwaitingJoining
	CandidateDropped
	CandidateNotResponding
	CandidateJoined
	CandidateNotJoined
)

func (r CandidateStatus) String() string {
	return [...]string{"AwaitingDocuments", "DocumentsSubmitted", "DocumentsVerified",
		"OfferFinalized", "AwaitingOfferApproval", "OfferApproved", "AwaitingOfferAcceptance",
		"OfferAccepted", "OfferRejected", "OfferExpired", "OfferWithdrawn", "AwaitingJoining",
		"CandidateDropped", "CandidateNotResponding", "CandidateJoined", "CandidateNotJoined"}[r-1]
}

func (r CandidateStatus) EnumIndex() int {
	return int(r)
}

func GetAllCandidateStatus() []string {
	return []string{"AwaitingDocuments", "DocumentsSubmitted", "DocumentsVerified",
		"OfferFinalized", "AwaitingOfferApproval", "OfferApproved", "AwaitingOfferAcceptance",
		"OfferAccepted", "OfferRejected", "OfferExpired", "OfferWithdrawn", "AwaitingJoining",
		"CandidateDropped", "CandidateNotResponding", "CandidateJoined", "CandidateNotJoined"}
}

func CandidateStatusFromString(s string) CandidateStatus {
	switch s {
	case "AwaitingDocuments":
		return AwaitingDocuments
	case "DocumentsSubmitted":
		return DocumentsSubmitted
	case "DocumentsVerified":
		return DocumentsVerified
	case "OfferFinalized":
		return OfferFinalized
	case "AwaitingOfferApproval":
		return AwaitingOfferApproval
	case "OfferApproved":
		return OfferApproved
	case "AwaitingOfferAcceptance":
		return AwaitingOfferAcceptance
	case "OfferAccepted":
		return OfferAccepted
	case "OfferRejected":
		return OfferRejected
	case "OfferExpired":
		return OfferExpired
	case "OfferWithdrawn":
		return OfferWithdrawn
	case "AwaitingJoining":
		return AwaitingJoining
	case "CandidateDropped":
		return CandidateDropped
	case "CandidateNotResponding":
		return CandidateNotResponding
	case "CandidateJoined":
		return CandidateJoined
	case "CandidateNotJoined":
		return CandidateNotJoined
	default:
		return 0
	}
}
