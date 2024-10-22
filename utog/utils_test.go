package utog

import (
	"testing"

	"github.com/invopop/gobl/cbc"
	"github.com/stretchr/testify/assert"
)

// Define tests for the ParseDate function
func TestParseDate(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    string
		expectError bool
	}{
		{"Valid date", "20230515", "2023-05-15", false},
		{"Invalid date", "20231345", "", true},
		{"Empty string", "", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseDate(tt.input)
			if tt.expectError {
				assert.Error(t, err)
				assert.Empty(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result.String())
			}
		})
	}
}

// Define tests for the FindTaxKey function
func TestFindTaxKey(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Standard sales tax", "S", "standard"},
		{"Zero rated goods tax", "Z", "zero"},
		{"Tax exempt", "E", "exempt"},
		{"Unknown tax type", "X", "standard"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindTaxKey(tt.input)
			assert.Equal(t, tt.expected, string(result))
		})
	}
}

// Define tests for the TypeCodeParse function
func TestTypeCodeParse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Standard invoice", "380", "standard"},
		{"Credit note", "381", "credit-note"},
		{"Corrective invoice", "384", "corrective"},
		{"Proforma invoice", "325", "proforma"},
		{"Debit note", "383", "debit-note"},
		{"Unknown type code", "999", "other"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TypeCodeParse(tt.input)
			assert.Equal(t, tt.expected, string(result))
		})
	}
}

// Define tests for the UnitFromUNECE function
func TestUnitFromUNECE(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Known UNECE code", "HUR", "h"},
		{"Known UNECE code", "SEC", "s"},
		{"Known UNECE code", "MTR", "m"},
		{"Known UNECE code", "GRM", "g"},
		{"Unknown UNECE code", "XYZ", "XYZ"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := UnitFromUNECE(cbc.Code(tt.input))
			assert.Equal(t, tt.expected, string(result))
		})
	}
}