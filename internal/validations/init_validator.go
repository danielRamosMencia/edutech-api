package validations

import (
	"log"
	"sync"

	"github.com/go-playground/locales/es"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	es_translations "github.com/go-playground/validator/v10/translations/es"
)

var (
	once       sync.Once
	Validate   *validator.Validate
	Translator ut.Translator
)

func InitValidator() {
	once.Do(func() {
		Validate = validator.New()
		spanish := es.New()
		uni := ut.New(spanish, spanish)

		var found bool
		Translator, found = uni.GetTranslator("es")
		if !found {
			log.Fatal("translator not found")

		}

		err := es_translations.RegisterDefaultTranslations(Validate, Translator)
		if err != nil {
			log.Fatal("error registering default translations", err)
		}
	})
}

func MapValidatorErrors(err error) []string {
	var errFields []string

	for _, err := range err.(validator.ValidationErrors) {
		errFields = append(errFields, err.Translate(Translator))
	}

	return errFields
}
