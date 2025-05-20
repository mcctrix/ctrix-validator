package validator

import (
	"regexp"
	"strconv"
)

type validationError struct {
	Field   string
	Message string
}

type validatorApp struct {
	fieldName     string
	data          interface{}
	requiredField bool
	errors        []validationError
	foundErr      bool
}

/*
This function checks if the field contains any special character

returns: *validatorApp
*/
func (v *validatorApp) HasSpecialChar() *validatorApp {
	if v.commonReturnCase() {
		return v
	}

	switch v.data.(type) {
	case string:
		if len(v.data.(string)) > 0 {
			specialCharRegex := regexp.MustCompile(`[^a-zA-Z0-9\s]+`)
			if specialCharRegex.MatchString(v.data.(string)) {
				return v
			}
		}
		v.appendError(validationError{
			Field:   v.fieldName,
			Message: "must contain at least one special character",
		})
	}
	return v
}

/*
This function checks if the field is a valid email

returns: *validatorApp
*/
func (v *validatorApp) Email() *validatorApp {
	if v.commonReturnCase() {
		return v
	}
	switch v.data.(type) {
	case string:
		if len(v.data.(string)) > 0 {
			str := v.data.(string)
			// Email pattern compliant with most RFCs while still using Go's RE2 engine
			emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
			if emailRegex.MatchString(str) {
				return v
			}
		}
		v.appendError(validationError{
			Field:   v.fieldName,
			Message: "must be a valid email",
		})
	}
	return v
}

/*
This function checks if the field has minimum length

returns: *validatorApp
*/
func (v *validatorApp) Min(length int) *validatorApp {

	if v.commonReturnCase() {
		return v
	}

	errMessage := "must be greater than or equal to " + strconv.Itoa(length)

	switch v.data.(type) {
	case string:
		if len(v.data.(string)) < length {
			v.appendError(validationError{
				Field:   v.fieldName,
				Message: errMessage,
			})
		}
	case int:
		if v.data.(int) < length {
			v.appendError(validationError{
				Field:   v.fieldName,
				Message: errMessage,
			})

		}
	case int16:
		if v.data.(int16) < int16(length) {
			v.appendError(validationError{
				Field:   v.fieldName,
				Message: errMessage,
			})

		}
	case int32:
		if v.data.(int32) < int32(length) {
			v.appendError(validationError{
				Field:   v.fieldName,
				Message: errMessage,
			})

		}
	case int64:
		if v.data.(int64) < int64(length) {
			v.appendError(validationError{
				Field:   v.fieldName,
				Message: errMessage,
			})

		}
	case int8:
		if v.data.(int8) < int8(length) {
			v.appendError(validationError{
				Field:   v.fieldName,
				Message: errMessage,
			})

		}
	case uint:
		if v.data.(uint8) < uint8(length) {
			v.appendError(validationError{
				Field:   v.fieldName,
				Message: errMessage,
			})

		}
	case uint8:
		if v.data.(uint8) < uint8(length) {
			v.appendError(validationError{
				Field:   v.fieldName,
				Message: errMessage,
			})

		}
	case uint16:
		if v.data.(uint16) < uint16(length) {
			v.appendError(validationError{
				Field:   v.fieldName,
				Message: errMessage,
			})

		}
	case uint32:
		if v.data.(uint32) < uint32(length) {
			v.appendError(validationError{
				Field:   v.fieldName,
				Message: errMessage,
			})
		}
	case uint64:
		if v.data.(uint64) < uint64(length) {
			v.appendError(validationError{
				Field:   v.fieldName,
				Message: errMessage,
			})

		}
	case float32:
		if v.data.(float32) < float32(length) {
			v.appendError(validationError{
				Field:   v.fieldName,
				Message: errMessage,
			})

		}

	case float64:
		if v.data.(float64) < float64(length) {
			v.appendError(validationError{
				Field:   v.fieldName,
				Message: errMessage,
			})
		}

	}
	return v
}

/*
This function checks if the field has maximum length

returns: *validatorApp
*/
func (v *validatorApp) Max(length int) *validatorApp {
	if v.commonReturnCase() {
		return v
	}

	errMessage := "must be less than or equal to " + strconv.Itoa(length)

	switch v.data.(type) {
	case string:
		if len(v.data.(string)) > length {
			v.appendError(validationError{
				Field:   v.fieldName,
				Message: errMessage,
			})

		}
	case int:
		if v.data.(int) > length {
			v.appendError(validationError{
				Field:   v.fieldName,
				Message: errMessage,
			})

		}
	case int16:
		if v.data.(int16) > int16(length) {
			v.appendError(validationError{
				Field:   v.fieldName,
				Message: errMessage,
			})

		}
	case int32:
		if v.data.(int32) > int32(length) {
			v.appendError(validationError{
				Field:   v.fieldName,
				Message: errMessage,
			})

		}
	case int64:
		if v.data.(int64) > int64(length) {
			v.appendError(validationError{
				Field:   v.fieldName,
				Message: errMessage,
			})

		}
	case int8:
		if v.data.(int8) > int8(length) {
			v.appendError(validationError{
				Field:   v.fieldName,
				Message: errMessage,
			})

		}
	case uint:
		if v.data.(uint8) > uint8(length) {
			v.appendError(validationError{
				Field:   v.fieldName,
				Message: errMessage,
			})

		}
	case uint8:
		if v.data.(uint8) > uint8(length) {
			v.appendError(validationError{
				Field:   v.fieldName,
				Message: errMessage,
			})

		}
	case uint16:
		if v.data.(uint16) > uint16(length) {
			v.appendError(validationError{
				Field:   v.fieldName,
				Message: errMessage,
			})

		}
	case uint32:
		if v.data.(uint32) > uint32(length) {
			v.appendError(validationError{
				Field:   v.fieldName,
				Message: errMessage,
			})
		}
	case uint64:
		if v.data.(uint64) > uint64(length) {
			v.appendError(validationError{
				Field:   v.fieldName,
				Message: errMessage,
			})

		}
	case float32:
		if v.data.(float32) > float32(length) {
			v.appendError(validationError{
				Field:   v.fieldName,
				Message: errMessage,
			})

		}

	case float64:
		if v.data.(float64) > float64(length) {
			v.appendError(validationError{
				Field:   v.fieldName,
				Message: errMessage,
			})
		}
	}
	return v
}

/*
This function checks if the field is a valid url

returns: *validatorApp
*/
func (v *validatorApp) Url() *validatorApp {
	if v.commonReturnCase() {
		return v
	}

	switch v.data.(type) {
	case string:
		if len(v.data.(string)) > 0 {
			urlRegex := regexp.MustCompile(`^(https?|ftp):\/\/` + // protocol
				`(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*` + // subdomains
				`([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])` + // domain name
				`(:\d+)?` + // port
				`(\/[-a-zA-Z0-9_%\.~#?&=]*)*$`) // path, query params, fragment

			if urlRegex.MatchString(v.data.(string)) {
				return v
			}
		}
		v.appendError(validationError{
			Field:   v.fieldName,
			Message: "must be a valid url",
		})

	}
	return v
}

/*
This function checks if the field contains only alphabets

returns: *validatorApp
*/
func (v *validatorApp) Alpha() *validatorApp {
	if v.commonReturnCase() {
		return v
	}

	switch v.data.(type) {
	case string:
		if len(v.data.(string)) > 0 {
			alphaRegex := regexp.MustCompile(`^[a-zA-Z]+$`)
			if alphaRegex.MatchString(v.data.(string)) {
				return v
			}
		}
		v.appendError(validationError{
			Field:   v.fieldName,
			Message: "must contain only alphabets",
		})
	}
	return v
}

/*
This function checks if the field contains only Numeric Values

returns: *validatorApp
*/
func (v *validatorApp) Numeric() *validatorApp {
	if v.commonReturnCase() {
		return v
	}

	switch v.data.(type) {
	case string:
		if len(v.data.(string)) > 0 {
			numericRegex := regexp.MustCompile(`^[0-9]+$`)
			if numericRegex.MatchString(v.data.(string)) {
				return v
			}
		}
		v.appendError(validationError{
			Field:   v.fieldName,
			Message: "must contain only numbers",
		})
	}
	return v
}

/*
This function checks if the field contains only AlphaNumeric Values

returns: *validatorApp
*/
func (v *validatorApp) AlphaNumeric() *validatorApp {
	if v.commonReturnCase() {
		return v
	}

	switch v.data.(type) {
	case string:
		if len(v.data.(string)) > 0 {
			alphaNumericRegex := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
			if alphaNumericRegex.MatchString(v.data.(string)) {
				return v
			}
		}
		v.appendError(validationError{
			Field:   v.fieldName,
			Message: "must contain only alphabets and numbers",
		})
	}
	return v
}

/*
This function checks if the field contains Valid Date

returns: *validatorApp
*/
func (v *validatorApp) Date() *validatorApp {
	if v.commonReturnCase() {
		return v
	}

	switch v.data.(type) {
	case string:
		if len(v.data.(string)) > 0 {
			dateRegex := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
			if dateRegex.MatchString(v.data.(string)) {
				return v
			}
		}
		v.appendError(validationError{
			Field:   v.fieldName,
			Message: "must be a valid date",
		})
	}
	return v
}

/*
This function checks if the field matches the pattern

returns: *validatorApp
*/
func (v *validatorApp) Match(pattern *regexp.Regexp) *validatorApp {
	if v.commonReturnCase() {
		return v
	}

	switch v.data.(type) {
	case string:
		if len(v.data.(string)) > 0 {
			if pattern.MatchString(v.data.(string)) {
				return v
			}
		}
		v.appendError(validationError{
			Field:   v.fieldName,
			Message: "must match the pattern",
		})
	}
	return v
}

/*
This function applies validations conditionally if the provided condition is true

- condition: A function that evaluates to a boolean
- validations: A function that applies validations to the validator

returns: *validatorApp

Example:

	validator := NewValidator("password", password).
	    When(
	        func() bool {
	            return userType == "admin"
	        },
	        func(v *validatorApp) *validatorApp {
	            return v.Min(12).HasSpecialChar()
	        },
	    ).
	    When(
	        func() bool {
	            return userType == "regular"
	        },
	        func(v *validatorApp) *validatorApp {
	            return v.Min(8)
	        },
	    )
*/
// func (v *validatorApp) When(condition func() bool, validations func(*validatorApp) *validatorApp) *validatorApp {
// 	// Skip this if we've already found errors or it's an optional empty field
// 	if v.commonReturnCase() {
// 		return v
// 	}

// 	// If the condition is true, run the validations
// 	if condition() {
// 		return validations(v)
// 	}

// 	// Otherwise, just return the validator as is
// 	return v
// }

func (v *validatorApp) appendError(err validationError) {
	v.errors = append(v.errors, err)
	v.foundErr = true
}

/*
This function makes the field optional

returns: *validatorApp
*/
func (v *validatorApp) NotRequired() *validatorApp {
	v.requiredField = false
	return v
}

/*
This function returns the errors

returns: *validatorApp
*/
func (v *validatorApp) GetError() []validationError {
	if len(v.errors) == 0 {
		return nil
	}
	return v.errors
}

/*
Allow custom error messages

- fieldName: name of the field

- message: custom error message

- returns: *validatorApp
*/
func (v *validatorApp) ChangeErrorMessage(fieldName string, message string) *validatorApp {
	if len(v.errors) > 0 {
		for index, err := range v.errors {
			if err.Field == fieldName {
				v.errors[index].Message = message
			}
		}
	}
	return v
}

/* Allow transformation during validation */
func (v *validatorApp) Transform(fn func(interface{}) interface{}) *validatorApp {
	if v.commonReturnCase() {
		return v
	}
	v.data = fn(v.data)
	return v
}

func (v *validatorApp) commonReturnCase() bool {
	var dataNullish bool

	switch v.data.(type) {
	case string:
		if v.data.(string) == "" {
			dataNullish = true
		}
	case int:
		if v.data.(int) == 0 {
			dataNullish = true
		}
	case float64:
		if v.data.(float64) == float64(0) {
			dataNullish = true
		}
	}

	if dataNullish && !v.requiredField {
		return true
	}
	if dataNullish && v.requiredField {
		v.appendError(validationError{
			Field:   v.fieldName,
			Message: "Field is Required",
		})
		return true
	}

	if v.foundErr {
		return true
	}
	return false
}

/*
  - - fieldName: name of the field
    *
  - - data: data to be validated
    *
  - - returns: validatorApp
  - Note : By Default field is considered required, please use NotRequired() immeadiately after to make it optional
*/
func (v *validatorApp) NextField(fieldName string, data interface{}) *validatorApp {
	v.fieldName = fieldName
	v.data = data
	v.requiredField = true
	v.foundErr = false
	return v
}

/*
- fieldName: name of the field

- data: data to be validated

- returns: validatorApp

- Note: By Default field is considered required, please use NotRequired() immeadiately after to make it optional
*/
func NewValidator(fieldName string, data interface{}) *validatorApp {
	vApp := &validatorApp{
		fieldName:     fieldName,
		data:          data,
		errors:        make([]validationError, 0),
		foundErr:      false,
		requiredField: true,
	}
	return vApp
}
