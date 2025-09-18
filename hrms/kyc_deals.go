package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Deal struct {
	Id                 primitive.ObjectID `bson:"_id,omitempty"` // Unique Deal ID
	DealId             string             `bson:"deal_id"`
	LeadId             string             `bson:"lead_id"`
	DealName           string             `bson:"deal_name"`             // Deal Name / Reference
	DealOwner          ContactInfo        `bson:"deal_owner"`            // Internal team member
	ExpectedCloseDate  time.Time          `bson:"expected_closing_date"` // Expected closing date
	Stage              DealStage          `bson:"stage"`                 // Discovery, Demo/POC, Proposal Sent, etc.
	DealValue          float64            `bson:"deal_value"`
	Currency           string             `bson:"currency"`                 // INR / USD
	DocumentLinks      []Attachment       `bson:"document_links,omitempty"` // Proposals, NDAs, MOMs
	Timeline           []LeadRemark       `bson:"timeline"`
	WinLossReason      string             `bson:"win_loss_reason,omitempty"` // Reason for closed-lost
	ProbabilityToClose int                `bson:"probability_to_close,omitempty"`
	DealSource         string             `bson:"deal_source,omitempty"` // Referral, Inbound, Outbound, etc.
	CreatedAt          time.Time          `bson:"created_at"`
}

type DealStage int

const (
	Discovery DealStage = iota + 1
	DemoPOC
	ProposalSentStage
	PriceNegotiation
	ContractSignOff
	Won
	Lost
)

func (s DealStage) String() string {
	stages := []string{
		"Discovery",
		"Demo/POC",
		"Proposal Sent",
		"Price Negotiation",
		"Contract Sign-off",
		"Won",
		"Lost",
	}
	i := int(s) - 1
	if i < 0 || i >= len(stages) {
		return "Unknown"
	}
	return stages[i]
}

func GetDealStages() []string {
	return []string{
		"Discovery",
		"Demo/POC",
		"Proposal Sent",
		"Price Negotiation",
		"Contract Sign-off",
		"Won",
		"Lost",
	}
}

func DealStageFromString(s string) DealStage {
	switch s {
	case "Discovery":
		return Discovery
	case "Demo/POC":
		return DemoPOC
	case "Proposal Sent":
		return ProposalSentStage
	case "Price Negotiation":
		return PriceNegotiation
	case "Contract Sign-off":
		return ContractSignOff
	case "Won":
		return Won
	case "Lost":
		return Lost
	default:
		return Discovery // fallback
	}
}
