package validate

import (
	db "auth-and-gateway-microservice/db/sqlc"
	"github.com/go-playground/validator/v10"
	"unicode"
)

var validate = validator.New()

func init() {
	_ = validate.RegisterValidation("password", validatePassword)
}

func CreateUserTxParams(userParams *db.CreateUserTxParams) error {
	if err := validate.Struct(*userParams); err != nil {
		return err
	}
	return nil
}

func validatePassword(f1 validator.FieldLevel) bool {
	password := f1.Field().String()

	// check if password contains at least one uppercase letter
	hasUpperCase := false
	for _, char := range password {
		if unicode.IsUpper(char) {
			hasUpperCase = true
			break
		}
	}
	hasLowerCase := false
	for _, char := range password {
		if unicode.IsLower(char) {
			hasLowerCase = true
			break
		}
	}

	isAtLeast8Chars := len(password) >= 8

	return hasLowerCase && hasUpperCase && isAtLeast8Chars
}
