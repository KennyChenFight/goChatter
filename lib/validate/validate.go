package validate

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

var updateValidate *validator.Validate

const (
	ValidatorTag string = "update"
)

func init() {
	updateValidate = validator.New()
}

func StructForUpdate(obj interface{}, structFieldNames map[string]bool) error {
	immutable := reflect.ValueOf(obj).Elem()
	immutableType := immutable.Type()
	updateValidate = validator.New()
	for fieldName := range structFieldNames {
		fieldName := strings.ToUpper(fieldName[0:1]) + fieldName[1:]
		field := immutable.FieldByName(fieldName)
		fieldType, _ := immutableType.FieldByName(fieldName)
		if fieldType.Tag.Get(ValidatorTag) == "fixed" {
			return errors.New("do not put something id field, it's fixed")
		}
		err := updateValidate.Var(field.Interface(), fieldType.Tag.Get(ValidatorTag))
		if err != nil {
			return errors.New(strings.ReplaceAll(err.Error(), "''", fieldName))
		}
	}
	return nil
}
