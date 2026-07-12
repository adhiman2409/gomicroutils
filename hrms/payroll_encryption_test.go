package hrms

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const testKey = "unit-test-encryption-key-0123456789"

func sampleSalaryStructure() SalaryStructureDetails {
	return SalaryStructureDetails{
		ID:               primitive.NewObjectID(),
		EmployeeId:       "PT24003",
		EmployeeName:     "Jane Doe",
		EmploymentStatus: "active",
		EmployeeType:     "permanent",
		ActiveSalaryStructure: CountrySalaryStructure{
			SSID:          "SS-1",
			EmployeeId:    "PT24003",
			FinancialYear: "2025-2026",
			PAN:           "ABCDE1234F",
			UAN:           "100200300400",
			UID:           "999988887777",
			AccountNumber: "1234567890",
			IFSC:          "HDFC0001234",
			AnnualCTC:     1200000,
			MonthlyNetPay: 85000,
			AnnualBasicSalary: 600000,
		},
	}
}

// storedForm marshals via the codec and decodes to a raw map to inspect what actually
// lands in MongoDB.
func storedForm(t *testing.T, v interface{}) bson.M {
	t.Helper()
	raw, err := bson.Marshal(v)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var m bson.M
	if err := bson.Unmarshal(raw, &m); err != nil {
		t.Fatalf("unmarshal to map: %v", err)
	}
	return m
}

func TestSalaryStructureEncryptsSensitiveKeepsQueryKeys(t *testing.T) {
	t.Setenv("ENCRYPTION_KEY", testKey)
	ss := sampleSalaryStructure()

	stored := storedForm(t, ss)

	// Query keys stay plaintext.
	if stored["employee_id"] != "PT24003" {
		t.Errorf("employee_id should be plaintext, got %v", stored["employee_id"])
	}
	if stored["employee_name"] != "Jane Doe" {
		t.Errorf("employee_name should be plaintext, got %v", stored["employee_name"])
	}
	// Sensitive nested sub-document is not present in plaintext at the top level.
	if _, ok := stored["active_salary_structure"]; ok {
		t.Error("active_salary_structure must not be stored in plaintext")
	}
	// The encrypted blob is present.
	blob, ok := stored[encKey].(string)
	if !ok || blob == "" {
		t.Fatalf("expected encrypted blob under %q, got %v", encKey, stored[encKey])
	}
	// The raw stored bytes must not leak any sensitive value.
	raw, _ := bson.Marshal(ss)
	for _, secret := range []string{"ABCDE1234F", "1234567890", "HDFC0001234", "999988887777"} {
		if bytesContains(raw, secret) {
			t.Errorf("sensitive value %q leaked into stored document", secret)
		}
	}
}

func TestSalaryStructureRoundTrip(t *testing.T) {
	t.Setenv("ENCRYPTION_KEY", testKey)
	ss := sampleSalaryStructure()

	raw, err := bson.Marshal(ss)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var got SalaryStructureDetails
	if err := bson.Unmarshal(raw, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if got.EmployeeId != ss.EmployeeId {
		t.Errorf("employee_id: got %q want %q", got.EmployeeId, ss.EmployeeId)
	}
	if got.ActiveSalaryStructure.AnnualCTC != ss.ActiveSalaryStructure.AnnualCTC {
		t.Errorf("annual_ctc: got %v want %v", got.ActiveSalaryStructure.AnnualCTC, ss.ActiveSalaryStructure.AnnualCTC)
	}
	if got.ActiveSalaryStructure.PAN != ss.ActiveSalaryStructure.PAN {
		t.Errorf("pan: got %q want %q", got.ActiveSalaryStructure.PAN, ss.ActiveSalaryStructure.PAN)
	}
	if got.ActiveSalaryStructure.AccountNumber != ss.ActiveSalaryStructure.AccountNumber {
		t.Errorf("account_number: got %q want %q", got.ActiveSalaryStructure.AccountNumber, ss.ActiveSalaryStructure.AccountNumber)
	}
}

func TestProcessingRequestNestedArrayRoundTrip(t *testing.T) {
	t.Setenv("ENCRYPTION_KEY", testKey)
	req := SalaryProcessingRequest{
		Id:            primitive.NewObjectID(),
		RequestId:     "REQ-1",
		Month:         "April",
		FinancialYear: "2025-2026",
		RequestStatus: "WaitingApproval",
		ApproverId:    "PT24002",
		TotalSalaryAmount: 500000,
		EmployeeMonthlySalaryDetails: []EmployeeMonthlySalaryDetail{
			{EmployeeID: "PT24003", EmployeeName: "Jane", PAN: "ABCDE1234F", NetSalary: 85000, AccountNumber: "1234567890"},
			{EmployeeID: "PT24004", EmployeeName: "John", PAN: "ZZZZZ9999Z", NetSalary: 92000, AccountNumber: "9999999999"},
		},
	}

	stored := storedForm(t, req)
	if stored["approver_id"] != "PT24002" || stored["request_status"] != "WaitingApproval" {
		t.Error("approver_id/request_status must stay plaintext for leaves.unirms.com counts")
	}
	if _, ok := stored["employee_monthly_salary_details"]; ok {
		t.Error("employee_monthly_salary_details must be encrypted, not plaintext")
	}

	raw, _ := bson.Marshal(req)
	var got SalaryProcessingRequest
	if err := bson.Unmarshal(raw, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if len(got.EmployeeMonthlySalaryDetails) != 2 {
		t.Fatalf("expected 2 employee details, got %d", len(got.EmployeeMonthlySalaryDetails))
	}
	if got.EmployeeMonthlySalaryDetails[1].NetSalary != 92000 {
		t.Errorf("net_salary: got %v want 92000", got.EmployeeMonthlySalaryDetails[1].NetSalary)
	}
	if got.TotalSalaryAmount != 500000 {
		t.Errorf("total_salary_amount: got %v want 500000", got.TotalSalaryAmount)
	}
}

func TestDualModeReadsLegacyPlaintext(t *testing.T) {
	t.Setenv("ENCRYPTION_KEY", testKey)
	ss := sampleSalaryStructure()

	// Simulate a legacy document written before encryption: default codec, no blob.
	type alias SalaryStructureDetails
	legacy, err := bson.Marshal(alias(ss))
	if err != nil {
		t.Fatalf("marshal legacy: %v", err)
	}
	if _, ok := storedForm(t, alias(ss))[encKey]; ok {
		t.Fatal("legacy fixture should not contain a blob")
	}

	var got SalaryStructureDetails
	if err := bson.Unmarshal(legacy, &got); err != nil {
		t.Fatalf("unmarshal legacy: %v", err)
	}
	if got.ActiveSalaryStructure.AnnualCTC != ss.ActiveSalaryStructure.AnnualCTC {
		t.Errorf("legacy read annual_ctc: got %v want %v", got.ActiveSalaryStructure.AnnualCTC, ss.ActiveSalaryStructure.AnnualCTC)
	}
	if got.ActiveSalaryStructure.PAN != ss.ActiveSalaryStructure.PAN {
		t.Errorf("legacy read pan: got %q want %q", got.ActiveSalaryStructure.PAN, ss.ActiveSalaryStructure.PAN)
	}
}

func TestFailsFastWithoutKey(t *testing.T) {
	t.Setenv("ENCRYPTION_KEY", "")
	ss := sampleSalaryStructure()
	if _, err := bson.Marshal(ss); err == nil {
		t.Error("expected marshal to fail when ENCRYPTION_KEY is unset")
	}
}

func bytesContains(haystack []byte, needle string) bool {
	n := []byte(needle)
	if len(n) == 0 {
		return false
	}
	for i := 0; i+len(n) <= len(haystack); i++ {
		if string(haystack[i:i+len(n)]) == string(n) {
			return true
		}
	}
	return false
}
