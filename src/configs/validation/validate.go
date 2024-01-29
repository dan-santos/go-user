package validation

import (
	"encoding/json"
	"errors"

	resterrors "github.com/dan-santos/go-user/src/configs/rest_errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translatation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	trans ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		trans, _ = unt.GetTranslator("en")
		en_translatation.RegisterDefaultTranslations(val, trans)
	}
}

func ValidateUserError(validation_err error) *resterrors.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return resterrors.NewBadRequestError("Invalid field type")
	} else if errors.As(validation_err, &jsonValidationError) {
		errorCauses := []resterrors.Cause{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := resterrors.Cause{
				Message: e.Translate(trans),
				Field: e.Field(),
			}

			errorCauses = append(errorCauses, cause)
		}

		return resterrors.NewBadRequestValidationError("Some field are invalid", errorCauses)
	}
	return resterrors.NewBadRequestError("Error trying to convert fields")
}