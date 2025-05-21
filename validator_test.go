package validator

import (
	"regexp"
	"testing"
)

func TestHasSpecialChar(t *testing.T) {
	v := NewValidator("field", "test@123")
	if v.HasSpecialChar().GetError() != nil {
		t.Errorf("HasSpecialChar() should not return an error for valid input")
	}

	v = NewValidator("field", "test123")
	if v.HasSpecialChar().GetError() == nil {
		t.Errorf("HasSpecialChar() should return an error for input without special characters")
	}
}

func TestEmail(t *testing.T) {
	v := NewValidator("field", "test@example.com")
	if v.Email().GetError() != nil {
		t.Errorf("Email() should not return an error for valid email")
	}

	v = NewValidator("field", "ctrix.com")
	if v.Email().GetError() == nil {
		t.Errorf("Email() should return an error for invalid email")
	}

	v = NewValidator("field", "ctrix@.com")
	if v.Email().GetError() == nil {
		t.Errorf("Email() should return an error for invalid email")
	}

	v = NewValidator("field", "@ctrix.com")
	if v.Email().GetError() == nil {
		t.Errorf("Email() should return an error for invalid email")
	}

	v = NewValidator("field", "invalid-email")
	if v.Email().GetError() == nil {
		t.Errorf("Email() should return an error for invalid email")
	}
}

func TestMin(t *testing.T) {
	v := NewValidator("field", "test")
	if v.Min(3).GetError() != nil {
		t.Errorf("Min() should not return an error for valid input")
	}

	v = NewValidator("field", "te")
	if v.Min(3).GetError() == nil {
		t.Errorf("Min() should return an error for input less than minimum length")
	}
}

func TestMax(t *testing.T) {
	v := NewValidator("field", "test")
	if v.Max(5).GetError() != nil {
		t.Errorf("Max() should not return an error for valid input")
	}

	v = NewValidator("field", "test123")
	if v.Max(5).GetError() == nil {
		t.Errorf("Max() should return an error for input greater than maximum length")
	}
}

func TestUrl(t *testing.T) {
	v := NewValidator("field", "http://example.com")
	if v.Url().GetError() != nil {
		t.Errorf("Url() should not return an error for valid URL")
	}

	v = NewValidator("field", "invalid-url")
	if v.Url().GetError() == nil {
		t.Errorf("Url() should return an error for invalid URL")
	}
}

func TestAlpha(t *testing.T) {
	v := NewValidator("field", "test")
	if v.Alpha().GetError() != nil {
		t.Errorf("Alpha() should not return an error for valid input")
	}

	v = NewValidator("field", "test123")
	if v.Alpha().GetError() == nil {
		t.Errorf("Alpha() should return an error for input with non-alphabetic characters")
	}
}

func TestNumeric(t *testing.T) {
	v := NewValidator("field", "123")
	if v.Numeric().GetError() != nil {
		t.Errorf("Numeric() should not return an error for valid input")
	}

	v = NewValidator("field", "test")
	if v.Numeric().GetError() == nil {
		t.Errorf("Numeric() should return an error for input with non-numeric characters")
	}
}

func TestAlphaNumeric(t *testing.T) {
	v := NewValidator("field", "test123")
	if v.AlphaNumeric().GetError() != nil {
		t.Errorf("AlphaNumeric() should not return an error for valid input")
	}

	v = NewValidator("field", "test@123")
	if v.AlphaNumeric().GetError() == nil {
		t.Errorf("AlphaNumeric() should return an error for input with non-alphanumeric characters")
	}
}

func TestDate(t *testing.T) {
	v := NewValidator("field", "2023-10-01")
	if v.Date().GetError() != nil {
		t.Errorf("Date() should not return an error for valid date")
	}

	v = NewValidator("field", "--")
	if v.Date().GetError() == nil {
		t.Errorf("Date() should return an error for invalid date")
	}

	v = NewValidator("field", "11-12")
	if v.Date().GetError() == nil {
		t.Errorf("Date() should return an error for invalid date")
	}

	v = NewValidator("field", "invalid-date")
	if v.Date().GetError() == nil {
		t.Errorf("Date() should return an error for invalid date")
	}
}

func TestMatch(t *testing.T) {
	pattern := regexp.MustCompile(`^[a-z]+$`)
	v := NewValidator("field", "test")
	if v.Match(pattern).GetError() != nil {
		t.Errorf("Match() should not return an error for valid input")
	}

	v = NewValidator("field", "test123")
	if v.Match(pattern).GetError() == nil {
		t.Errorf("Match() should return an error for input not matching the pattern")
	}
}

func TestPhoneNumber(t *testing.T) {
	v := NewValidator("field", "+1234567890")
	if v.PhoneNumber().GetError() != nil {
		t.Errorf("PhoneNumber() should not return an error for valid phone number")
	}

	v = NewValidator("field", "@+97")
	if v.PhoneNumber().GetError() == nil {
		t.Errorf("PhoneNumber() should return an error for invalid phone number")
	}

	v = NewValidator("field", "invalid-phone")
	if v.PhoneNumber().GetError() == nil {
		t.Errorf("PhoneNumber() should return an error for invalid phone number")
	}
}

func TestCreditCard(t *testing.T) {
	v := NewValidator("field", "4111111111111111")
	if v.CreditCard().GetError() != nil {
		t.Errorf("CreditCard() should not return an error for valid credit card number")
	}

	v = NewValidator("field", "0000000000000000")
	if v.CreditCard().GetError() == nil {
		t.Errorf("CreditCard() should return an error for invalid credit card number")
	}

	v = NewValidator("field", "invalid-card")
	if v.CreditCard().GetError() == nil {
		t.Errorf("CreditCard() should return an error for invalid credit card number")
	}
}

func TestIPAddress(t *testing.T) {
	v := NewValidator("field", "192.168.1.1")
	if v.IPAddress().GetError() != nil {
		t.Errorf("IPAddress() should not return an error for valid IP address")
	}

	v = NewValidator("field", "300.300.0.0")
	if v.IPAddress().GetError() == nil {
		t.Errorf("IPAddress() should return an error for invalid IP address")
	}
	v = NewValidator("field", "invalid-ip")
	if v.IPAddress().GetError() == nil {
		t.Errorf("IPAddress() should return an error for invalid IP address")
	}
}

func TestNotRequired(t *testing.T) {
	v := NewValidator("field", "")
	if v.NotRequired().GetError() != nil {
		t.Errorf("NotRequired() should not return an error for empty input")
	}
}

func TestGetError(t *testing.T) {
	v := NewValidator("field", "test")
	if v.GetError() != nil {
		t.Errorf("GetError() should return nil for valid input")
	}

	v = NewValidator("field", "")
	v.Email()
	if v.GetError() == nil {
		t.Errorf("GetError() should return an error for empty required field")
	}
}

func TestChangeErrorMessage(t *testing.T) {
	v := NewValidator("field", "")
	// First validate to generate an error
	v.Email()
	// Then change the error message
	v.ChangeErrorMessage("field", "Custom error message")
	errors := v.GetError()
	if len(errors) == 0 {
		t.Errorf("ChangeErrorMessage() should preserve errors")
	}
	if errors[0].Message != "Custom error message" {
		t.Errorf("ChangeErrorMessage() should change the error message")
	}
}

func TestTransform(t *testing.T) {
	v := NewValidator("field", "test")
	transformed := v.Transform(func(data interface{}) interface{} {
		return data.(string) + " transformed"
	})
	if transformed.GetError() != nil {
		t.Errorf("Transform() should not return an error for valid input")
	}
}

func TestNextField(t *testing.T) {
	v := NewValidator("field", "test")
	next := v.NextField("nextField", "nextValue")
	if next.GetError() != nil {
		t.Errorf("NextField() should not return an error for valid input")
	}
}

func TestMultipleValidations(t *testing.T) {
	// Test multiple validations on the same field
	v := NewValidator("field", "test@example.com")
	v.Email().Min(5).HasSpecialChar()
	if v.GetError() != nil {
		t.Errorf("Multiple validations should pass for valid input")
	}

	v = NewValidator("field", "test")
	v.Email().Min(5).HasSpecialChar()
	if v.GetError() == nil {
		t.Errorf("Multiple validations should fail for invalid input")
	}
}

func TestNumericTypes(t *testing.T) {
	// Test different numeric types
	tests := []struct {
		name     string
		value    interface{}
		min      int
		expected bool
	}{
		{"int", 42, 40, true},
		{"int8", int8(42), 40, true},
		{"int16", int16(42), 40, true},
		{"int32", int32(42), 40, true},
		{"int64", int64(42), 40, true},
		{"uint8", uint8(42), 40, true},
		{"uint16", uint16(42), 40, true},
		{"uint32", uint32(42), 40, true},
		{"uint64", uint64(42), 40, true},
		{"float32", float32(42.0), 40, true},
		{"float64", float64(42.0), 40, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NewValidator("field", tt.value)
			v.Min(tt.min)
			if (v.GetError() == nil) != tt.expected {
				t.Errorf("Min() for %s = %v; want %v", tt.name, v.GetError() == nil, tt.expected)
			}
		})
	}
}

func TestEdgeCases(t *testing.T) {
	// Test edge cases
	tests := []struct {
		name     string
		value    interface{}
		expected bool
	}{
		{"empty string", "", false},
		{"zero int", 0, false},
		{"zero float", 0.0, false},
		{"special chars only", "@#$%", true},
		{"special chars with number", "123@#$%", true},
		{"whitespace only", "   ", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NewValidator("field", tt.value)
			v.HasSpecialChar()
			if (v.GetError() == nil) != tt.expected {
				t.Errorf("HasSpecialChar() for %s = %v; want %v", tt.name, v.GetError() == nil, tt.expected)
			}
		})
	}
}

func TestCombinedValidations(t *testing.T) {
	// Test combinations of different validations
	tests := []struct {
		name        string
		value       string
		validations func(*validatorApp)
		expected    bool
	}{
		{
			"valid email with min length",
			"test@example.com",
			func(v *validatorApp) { v.Email().Min(10) },
			true,
		},
		{
			"valid phone with special chars",
			"+12345678900",
			func(v *validatorApp) { v.PhoneNumber().HasSpecialChar() },
			true,
		},
		{
			"valid URL with alpha numeric",
			"example123com",
			func(v *validatorApp) { v.Url().AlphaNumeric() },
			false,
		},
		{
			"invalid combination",
			"test",
			func(v *validatorApp) { v.Email().PhoneNumber() },
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NewValidator("field", tt.value)
			tt.validations(v)
			if (v.GetError() == nil) != tt.expected {
				t.Errorf("Combined validations for %s = %v; want %v", tt.name, v.GetError() == nil, tt.expected)
			}
		})
	}
}

func TestTransformAndValidate(t *testing.T) {
	// Test transformation followed by validation
	v := NewValidator("field", "test@example.com")
	transformed := v.Transform(func(data interface{}) interface{} {
		return "invalid-email" // Transform to clearly invalid email
	})
	transformed.Email()
	if transformed.GetError() == nil {
		t.Errorf("Email validation should fail after transformation")
	}

	v = NewValidator("field", "test")
	transformed = v.Transform(func(data interface{}) interface{} {
		return data.(string) + "@example.com"
	})
	transformed.Email()
	if transformed.GetError() != nil {
		t.Errorf("Email validation should pass after transformation")
	}
}

func TestNextFieldValidation(t *testing.T) {
	// Test validation across multiple fields
	v := NewValidator("email", "test@example.com")
	v.Email()

	v = v.NextField("phone", "+1234567890")
	v.PhoneNumber()

	v = v.NextField("url", "https://example.com")
	v.Url()

	if v.GetError() != nil {
		t.Errorf("All fields should be valid")
	}

	// Test with invalid fields
	v = NewValidator("email", "invalid-email")
	v.Email()

	v = v.NextField("phone", "invalid-phone")
	v.PhoneNumber()

	v = v.NextField("url", "invalid-url")
	v.Url()

	if v.GetError() == nil && len(v.GetError()) == 3 {
		t.Errorf("Should have errors for invalid fields")
	}
}

func TestNotRequiredWithValidations(t *testing.T) {
	// Test NotRequired with various validations
	tests := []struct {
		name        string
		value       interface{}
		validations func(*validatorApp)
		expected    bool
	}{
		{
			"empty not required",
			"",
			func(v *validatorApp) { v.NotRequired().Email() },
			true,
		},
		{
			"empty not required with multiple validations",
			"",
			func(v *validatorApp) { v.NotRequired().Email().Min(5) },
			true,
		},
		{
			"invalid not required",
			"invalid",
			func(v *validatorApp) { v.NotRequired().Email() },
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NewValidator("field", tt.value)
			tt.validations(v)
			if (v.GetError() == nil) != tt.expected {
				t.Errorf("NotRequired with validations for %s = %v; want %v", tt.name, v.GetError() == nil, tt.expected)
			}
		})
	}
}
