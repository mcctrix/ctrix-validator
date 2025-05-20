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
			emailRegex := regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$`)
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

	switch v.data.(type) {
	case string:
		if len(v.data.(string)) <= length {
			v.appendError(validationError{
				Field:   v.fieldName,
				Message: "must be greater than or equal to " + strconv.Itoa(length),
			})
		}
	case int:
		if v.data.(int) <= length {
			v.appendError(validationError{
				Field:   v.fieldName,
				Message: "must be greater than or equal to " + strconv.Itoa(length),
			})

		}
	case float64:
		if v.data.(float64) <= float64(length) {
			v.appendError(validationError{
				Field:   v.fieldName,
				Message: "must be greater than or equal to " + strconv.Itoa(length),
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
	switch v.data.(type) {
	case string:
		if len(v.data.(string)) >= length {
			v.appendError(validationError{
				Field:   v.fieldName,
				Message: "must be less than or equal to " + strconv.Itoa(length),
			})

		}
	case int:
		if v.data.(int) >= length {
			v.appendError(validationError{
				Field:   v.fieldName,
				Message: "must be less than or equal to " + strconv.Itoa(length),
			})

		}
	case float64:
		if v.data.(float64) >= float64(length) {
			v.appendError(validationError{
				Field:   v.fieldName,
				Message: "must be less than or equal to " + strconv.Itoa(length),
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
			urlRegex := regexp.MustCompile(`^(https?:\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)*\/?$`)
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

func (v *validatorApp) commonReturnCase() bool {
	if v.data == nil && !v.requiredField {
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
