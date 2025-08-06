package hrms

import "time"

type PolicyBazaarEmployeeInfo struct {
	EmployeeID       string                         `bson:"employee_id"`
	Name             string                         `bson:"name"`
	Email            string                         `bson:"email"`
	DateOfJoin       string                         `bson:"date_of_joining"`
	LeavingDate      string                         `bson:"leaving_date"`
	LeftOrganization bool                           `bson:"left_organization"`
	Status           int                            `bson:"status"`
	DateOfBirth      string                         `bson:"date_of_birth"`
	Gender           string                         `bson:"gender"`
	FamilyMembers    []PolicyBazaarFamilyMemberInfo `bson:"family_members"`
	LastModified     time.Time                      `bson:"last_modified"`
}

type ThirdPartyKey struct {
	Key        string    `bson:"key"`
	ClientName string    `bson:"client_name"`
	Title      string    `bson:"title"`
	IsActive   bool      `bson:"is_active"`
	Paths      []string  `bson:"paths"`
	CreatedAt  time.Time `bson:"created_at"`
	Expiry     time.Time `bson:"expiry"`
	Remarks    string    `bson:"remarks"`
}

type PolicyBazaarFamilyMemberInfo struct {
	Name       string `bson:"name"`
	RelationID int    `bson:"relation_id"`
	DOB        string `bson:"date_of_birth"`
	Gender     string `bson:"gender"`
}
