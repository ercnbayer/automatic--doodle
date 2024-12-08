package validator

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"unicode"

	"automatic-doodle/pkg/errors"

	_validator "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = _validator.New()

func ValidateStruct[T any](data *T) *errors.CustomError {
	errParams := []errors.ValidationErrorField{}
	err := validate.Struct(data)
	if err == nil {
		return nil
	}

	for _, err := range err.(_validator.ValidationErrors) {
		element := errors.ValidationErrorField{
			FailedField: err.Field(),
			Tag:         err.Tag(),
			Value:       fmt.Sprintf("%v", err.Value()),
		}
		errParams = append(errParams, element)
	}

	return errors.NewValidationError("ValidationMiddleware", errParams)
}

func ParseBody[T any](ctx *fiber.Ctx, pDestination *T) *errors.CustomError {
	if err := ctx.BodyParser(pDestination); err != nil {
		return errors.New(
			"BodyParser",
			err.Error(),
		)
	}

	if err := ValidateStruct(pDestination); err != nil {
		return err
	}

	return nil
}

func ParseQuery[T any](ctx *fiber.Ctx, pDestination *T) *errors.CustomError {
	if err := ctx.QueryParser(pDestination); err != nil {
		return errors.New("QueryParser", err.Error())
	}

	if err := ValidateStruct(pDestination); err != nil {
		return err
	}

	return nil
}

func ValidatePassword(fl _validator.FieldLevel) bool {
	password := fl.Field().String()

	var (
		hasMinLen      = false
		hasUpper       = false
		hasLower       = false
		hasNumber      = false
		hasSpecialChar = false
	)

	if len(password) >= 8 {
		hasMinLen = true
	}
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecialChar = true
		}
	}

	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecialChar
}

func validateFilename(fl _validator.FieldLevel) bool {
	filenameRegex := `^[a-zA-Z0-9_\-]+\.[a-zA-Z0-9]{1,10}$` // Simple regex for filenames
	matched, _ := regexp.MatchString(filenameRegex, fl.Field().String())
	return matched
}

// Custom validation for allowed file extensions.
func validateFileExtension(fl _validator.FieldLevel) bool {
	validExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
	}

	filename := fl.Field().String()
	for ext := range validExtensions {
		if len(filename) >= len(ext) && filename[len(filename)-len(ext):] == ext {
			return true
		}
	}
	return false
}

func init() {
	// use json tags as err.Field()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	validate.RegisterValidation("password", ValidatePassword)
	validate.RegisterValidation("filename", validateFilename)
	validate.RegisterValidation("fileext", validateFileExtension)
}
