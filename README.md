# Ctrix Validator

A lightweight and chainable Go validation library designed for fast development, focusing on a "first error per field" approach.
This library is inspired Chaining Validation from other languages.

## Features

- **Fluent Chaining API:** Apply multiple validation rules to a field in a concise, readable chain.
- **"First Error Per Field":** Validation for a given field stops after the first error is detected for that field.
- **Default Required:** All fields are considered required by default.
- **Optional Fields:** Easily mark fields as optional using `.NotRequired()`.
- **Graceful nullish Handling:** Optional fields with nullish data are skipped entirely, incurring no errors.
- **Concise Error Reporting:** Collects and returns a slice of `validationError` structs.
- **Silent Type Mismatch Skipping:** Designed for rapid development, type mismatches during validation (e.g., calling `Email()` on an `int`) will result in the specific validation rule being silently skipped without adding an error.

## Installation

To use this library, you can fetch it using `go get`:

```bash
go get github.com/mcctrix/ctrix-validator

```

## Usage

### Basic Usage

````go
package main

import (
	validator "github.com/mcctrix/ctrix-validator"
)

func main() {
	vApp := validator.NewValidator("email", "test@test.com").Email().Min(5).Max(30).HasSpecialChar()
	vApp.NextField("url", "https://google.com").Url()
	vApp.NextField("number", 10).Min(5).Max(15)
	vApp.NextField("float", 10.5).Min(5).Max(15)
	vApp.NextField("username", "ctrix").Min(4).Max(10).HasSpecialChar()

	if err := vApp.GetError(); err != nil {
		for _, err := range vApp.GetError() {
			println(err.Field, err.Message)
		}
	}
}
````

Output:
username must contain at least one special character

### Practical Usage

````go
package main

import (
	validator "github.com/mcctrix/ctrix-validator"
)

type userData struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	QuestNum int    `json:"questNum"`
}


func main() {
	app.Post("/user", func(c fiber.Ctx) error {
		data := &userData{}
		if err := json.Unmarshal(c.BodyRaw(), data); err != nil {
			fmt.Println(err)
			return fiber.ErrInternalServerError
		}

		vApp := validator.NewValidator("email", data.Email).Email()
		vApp.NextField("username", data.Username).Min(5).Max(25)
		vApp.NextField("password", data.Password).Min(8)
		vApp.NextField("Quest Number", data.QuestNum).NotRequired().Min(10).Max(20)

		if err := vApp.GetError(); err != nil {
			for _, errr := range err {
				fmt.Println("Error from Ctrix Validator: ", errr)
				return c.JSON(errr)
			}
		}

		return c.SendString("Data is Successfully added!")
	})
}
````

Payload:

```
{
   "username": "ctrix",
   "email": "ctrix@gmail.com",
   "password": "123456798",
   "questNum": 25
}
```

Error:

```
{
  "Field": "username",
  "Message": "must be greater than or equal to 5"
}
```

Note: In this example we are sending the first error that occurs in the chain. however getError() has details of all the errors that happened accross all the fields
