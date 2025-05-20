# Ctrix Validator

A lightweight and chainable Go validation library designed for fast development, focusing on a "first error per field" approach.
This library is inspired Chaining Validation from other languages.


## Features

* **Fluent Chaining API:** Apply multiple validation rules to a field in a concise, readable chain.
* **"First Error Per Field":** Validation for a given field stops after the first error is detected for that field.
* **Default Required:** All fields are considered required by default.
* **Optional Fields:** Easily mark fields as optional using `.NotRequired()`.
* **Graceful `nil` Handling:** Optional fields with `nil` data are skipped entirely, incurring no errors.
* **Concise Error Reporting:** Collects and returns a slice of `validationError` structs.
* **Silent Type Mismatch Skipping:** Designed for rapid development, type mismatches during validation (e.g., calling `Email()` on an `int`) will result in the specific validation rule being silently skipped without adding an error.

## Installation

To use this library, you can fetch it using `go get`:

```bash
go get https://github.com/mcctrix/ctrix-validator

```

## Usage

### Basic Usage

```go
package main

import (
	"github.com/mcctrix/ctrix-validator"
)

func main() {
	vApp := validator.NewValidator("email", "test@test.com").Email().Min(5).Max(10).HasSpecialChar()
	vApp.NextField("url", "https://google.com").Url()
    vApp.NextField("number", 10).Min(5).Max(15)
    vApp.NextField("float", 10.5).Min(5).Max(15)
    vApp.NextField("username", "ctrix").min(5).max(10)
    
	if err := vApp.GetError(); err != nil {
		for _, err := range validator.GetError() {
			println(err.Field, err.Message)
		}
	}
}
```

Output:

```
email must be a valid email
```

```go
