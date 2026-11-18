package hrms

import (
	"context"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestGetEmployeeComplianceInfoCollection_ReturnsNamedCollection(t *testing.T) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatalf("mongo.Connect: %v", err)
	}
	t.Cleanup(func() {
		_ = client.Disconnect(context.Background())
	})

	got := GetEmployeeComplianceInfoCollection(client, "test.domain")

	if got.Name() != PAYROLL_COLLECTION_EMPLOYEE_COMPLIANCE_INFO {
		t.Errorf("collection name: got %q want %q", got.Name(), PAYROLL_COLLECTION_EMPLOYEE_COMPLIANCE_INFO)
	}
}

func sampleEmployeeComplianceInfo() EmployeeComplianceInfo {
	return EmployeeComplianceInfo{
		ID:                     primitive.NewObjectID(),
		EmployeeID:             "PT24003",
		AadhaarLinked:          true,
		ESICInsuranceNumber:    "3412345678901234",
		IsPersonWithDisability: false,
		UpdatedAt:              time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC),
	}
}

func TestEmployeeComplianceInfoEncryptsSensitiveKeepsQueryKeys(t *testing.T) {
	t.Setenv("ENCRYPTION_KEY", testKey)
	info := sampleEmployeeComplianceInfo()

	stored := storedForm(t, info)

	// Query keys stay plaintext.
	if stored["employee_id"] != "PT24003" {
		t.Errorf("employee_id should be plaintext, got %v", stored["employee_id"])
	}
	if stored["aadhaar_linked"] != true {
		t.Errorf("aadhaar_linked should be plaintext, got %v", stored["aadhaar_linked"])
	}
	if stored["is_person_with_disability"] != false {
		t.Errorf("is_person_with_disability should be plaintext, got %v", stored["is_person_with_disability"])
	}
	// The sensitive field is not present in plaintext at the top level.
	if _, ok := stored["esic_insurance_number"]; ok {
		t.Error("esic_insurance_number must not be stored in plaintext")
	}
	// The encrypted blob is present.
	blob, ok := stored[encKey].(string)
	if !ok || blob == "" {
		t.Fatalf("expected encrypted blob under %q, got %v", encKey, stored[encKey])
	}
	// The raw stored bytes must not leak the sensitive value.
	raw, _ := bson.Marshal(info)
	if bytesContains(raw, info.ESICInsuranceNumber) {
		t.Errorf("sensitive value %q leaked into stored document", info.ESICInsuranceNumber)
	}
}

func TestEmployeeComplianceInfoRoundTrip(t *testing.T) {
	t.Setenv("ENCRYPTION_KEY", testKey)
	info := sampleEmployeeComplianceInfo()

	raw, err := bson.Marshal(info)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var got EmployeeComplianceInfo
	if err := bson.Unmarshal(raw, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if got.EmployeeID != info.EmployeeID {
		t.Errorf("employee_id: got %q want %q", got.EmployeeID, info.EmployeeID)
	}
	if got.ESICInsuranceNumber != info.ESICInsuranceNumber {
		t.Errorf("esic_insurance_number: got %q want %q", got.ESICInsuranceNumber, info.ESICInsuranceNumber)
	}
	if got.AadhaarLinked != info.AadhaarLinked {
		t.Errorf("aadhaar_linked: got %v want %v", got.AadhaarLinked, info.AadhaarLinked)
	}
	if !got.UpdatedAt.Equal(info.UpdatedAt) {
		t.Errorf("updated_at: got %v want %v", got.UpdatedAt, info.UpdatedAt)
	}
}

func TestEmployeeComplianceInfoFailsFastWithoutKey(t *testing.T) {
	t.Setenv("ENCRYPTION_KEY", "")
	info := sampleEmployeeComplianceInfo()
	if _, err := bson.Marshal(info); err == nil {
		t.Error("expected marshal to fail when ENCRYPTION_KEY is unset")
	}
}
