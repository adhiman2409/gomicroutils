package hrms

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestGetEmployeeTaxSavingDeclarationCollection_ReturnsNamedCollection(t *testing.T) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatalf("mongo.Connect: %v", err)
	}
	t.Cleanup(func() {
		_ = client.Disconnect(context.Background())
	})

	got := GetEmployeeTaxSavingDeclarationCollection(client, "test.domain")

	if got.Name() != PAYROLL_COLLECTION_EMPLOYEE_TAX_SAVING_DECLARATION {
		t.Errorf("collection name: got %q want %q", got.Name(), PAYROLL_COLLECTION_EMPLOYEE_TAX_SAVING_DECLARATION)
	}
}

func sampleEmployeeTaxSavingDeclaration() EmployeeTaxSavingDeclaration {
	return EmployeeTaxSavingDeclaration{
		EmployeeID:                           "PT24003",
		EmployeeName:                         "Jane Doe",
		FinancialYear:                        "2025-2026",
		OptedTaxRegime:                       "old",
		IncomeOtherThanSalary:                50000,
		HomeLoanInterest:                     185000,
		HRAExemption:                         120000,
		Section80CDeduction:                  150000,
		Section80CCD2EmployerNPSContribution: 60000,
		Section80DHealthInsuranceForSelf:     25000,
		Section80DHealthInsuranceForFamily:   25000,
		OtherDeductions:                      10000,
		ProofVerified:                        true,
	}
}

func TestEmployeeTaxSavingDeclarationEncryptsSensitiveKeepsQueryKeys(t *testing.T) {
	t.Setenv("ENCRYPTION_KEY", testKey)
	decl := sampleEmployeeTaxSavingDeclaration()

	stored := storedForm(t, decl)

	// Query keys stay plaintext.
	if stored["employee_id"] != "PT24003" {
		t.Errorf("employee_id should be plaintext, got %v", stored["employee_id"])
	}
	if stored["financial_year"] != "2025-2026" {
		t.Errorf("financial_year should be plaintext, got %v", stored["financial_year"])
	}
	if stored["opted_tax_regime"] != "old" {
		t.Errorf("opted_tax_regime should be plaintext, got %v", stored["opted_tax_regime"])
	}
	if stored["proof_verified"] != true {
		t.Errorf("proof_verified should be plaintext, got %v", stored["proof_verified"])
	}
	// Sensitive financial figures are not present in plaintext at the top level.
	for _, key := range []string{
		"home_loan_interest", "hra_exemption", "section_80c_deduction",
		"section_80ccd2_employer_nps_contribution", "section_80d_health_insurance_for_self",
		"section_80d_health_insurance_for_family", "other_deductions", "income_other_than_salary",
	} {
		if _, ok := stored[key]; ok {
			t.Errorf("%s must not be stored in plaintext", key)
		}
	}
	// The encrypted blob is present.
	blob, ok := stored[encKey].(string)
	if !ok || blob == "" {
		t.Fatalf("expected encrypted blob under %q, got %v", encKey, stored[encKey])
	}
	// The raw stored bytes must not leak any sensitive value.
	raw, _ := bson.Marshal(decl)
	for _, secret := range []string{"185000", "150000", "60000"} {
		if bytesContains(raw, secret) {
			t.Errorf("sensitive value %q leaked into stored document", secret)
		}
	}
}

func TestEmployeeTaxSavingDeclarationRoundTrip(t *testing.T) {
	t.Setenv("ENCRYPTION_KEY", testKey)
	decl := sampleEmployeeTaxSavingDeclaration()

	raw, err := bson.Marshal(decl)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var got EmployeeTaxSavingDeclaration
	if err := bson.Unmarshal(raw, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if got.EmployeeID != decl.EmployeeID {
		t.Errorf("employee_id: got %q want %q", got.EmployeeID, decl.EmployeeID)
	}
	if got.HomeLoanInterest != decl.HomeLoanInterest {
		t.Errorf("home_loan_interest: got %v want %v", got.HomeLoanInterest, decl.HomeLoanInterest)
	}
	if got.Section80CDeduction != decl.Section80CDeduction {
		t.Errorf("section_80c_deduction: got %v want %v", got.Section80CDeduction, decl.Section80CDeduction)
	}
	if got.Section80DHealthInsuranceForFamily != decl.Section80DHealthInsuranceForFamily {
		t.Errorf("section_80d_health_insurance_for_family: got %v want %v", got.Section80DHealthInsuranceForFamily, decl.Section80DHealthInsuranceForFamily)
	}
	if got.ProofVerified != decl.ProofVerified {
		t.Errorf("proof_verified: got %v want %v", got.ProofVerified, decl.ProofVerified)
	}
}

func TestEmployeeTaxSavingDeclarationFailsFastWithoutKey(t *testing.T) {
	t.Setenv("ENCRYPTION_KEY", "")
	decl := sampleEmployeeTaxSavingDeclaration()
	if _, err := bson.Marshal(decl); err == nil {
		t.Error("expected marshal to fail when ENCRYPTION_KEY is unset")
	}
}
