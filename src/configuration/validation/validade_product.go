package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	"github.com/karineaf/GoWithReactProject/src/configuration/handle_error"
)

var (
	Validade = validator.New()
	transl ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
        en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)

	}
}

func ValidadeProductError(validation_err error) *handle_error.HandleError {
	var jsonErr * json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr){
		return handle_error.BadRequestError("Invalid field type")
	} else if errors.As(validation_err, &jsonValidationError){
		errorsCauses := []handle_error.Causes{}
		for _, e := range validation_err.(validator.ValidationErrors){
			cause := handle_error.Causes{
                Field:  e.Field(),
                Message: e.Translate(transl),
            }
            errorsCauses = append(errorsCauses, cause)
		}
		return handle_error.BadRequestValidationError("Some fields are invalid", errorsCauses)
	} else {
		return handle_error.BadRequestError("Error trying to convert fields")
	}
}