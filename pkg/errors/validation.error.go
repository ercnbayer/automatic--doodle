package errors

type ValidationErrorField struct {
	FailedField string `json:"field"`
	Tag         string `json:"condition"`
	Value       string `json:"value"`
}

func NewValidationError(module string, params []ValidationErrorField) *CustomError {
	return &CustomError{
		Module:  module,
		Message: "Validation Failed",
		Params:  params,
		Status:  400,
	}
}
