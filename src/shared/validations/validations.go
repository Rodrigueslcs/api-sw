package validations

import (
	"api-sw/src/shared/tools/communication"
	"api-sw/src/shared/validations/validator"
	"fmt"

	"github.com/asaskevich/govalidator"
)

var Validator = validator.New()

var simpleValidator = map[string]struct {
	validator func(str string) bool
	message   string
	code      int
}{
	"required": {
		validator: Validator.String.IsNotNull,
		message:   communication.New().Mapping["validate_required"].Message,
		code:      communication.New().Mapping["validate_required"].Code,
	},
	"isRequiredCreated": {
		validator: govalidator.IsNotNull,
		message:   communication.New().Mapping["validate_required"].Message,
		code:      communication.New().Mapping["validate_required"].Code,
	},
	"isRequiredUpdated": {
		validator: govalidator.IsNotNull,
		message:   communication.New().Mapping["validate_required"].Message,
		code:      communication.New().Mapping["validate_required"].Code,
	},
	"email": {
		validator: govalidator.IsEmail,
		message:   communication.New().Mapping["validate_invalid"].Message,
		code:      communication.New().Mapping["validate_invalid"].Code,
	},
}

var stringLengthValidator = map[string]struct {
	validator func(value string, threshold string) bool
	message   func(value string) string
}{
	"isLessThan": {
		validator: Validator.String.IsLessThan,
		message: func(value string) string {
			return fmt.Sprintf("%s %s", communication.New().Mapping["less_than"].Message, value)
		},
	},
	"isLessOrEqualThan": {
		validator: Validator.String.IsLessOrEqualThan,
		message: func(value string) string {
			return fmt.Sprintf("%s %s", communication.New().Mapping["less_or_equal_than"].Message, value)
		},
	},
	"isGreaterThan": {
		validator: Validator.String.IsGreaterThan,
		message: func(value string) string {
			return fmt.Sprintf("%s %s", communication.New().Mapping["greater_than"].Message, value)
		},
	},
	"isGreaterOrEqualThan": {
		validator: Validator.String.IsGreaterOrEqualThan,
		message: func(value string) string {
			return fmt.Sprintf("%s %s", communication.New().Mapping["greater_or_equal_than"].Message, value)
		},
	},
}

var numberValidator = map[string]struct {
	validator func(value string, compare string) bool
	message   func(value string) string
}{
	"gte": {
		validator: func(value string, compare string) bool {
			number, err := govalidator.ToInt(value)
			if err != nil {
				return false
			}

			limit, err := govalidator.ToInt(compare)
			if err != nil {
				return false
			}

			return number < limit
		},
		message: func(value string) string {
			return fmt.Sprintf("%s %s", communication.New().Mapping["greater_or_equal_than"].Message, value)
		},
	},
	"lte": {
		validator: func(value string, compare string) bool {
			number, err := govalidator.ToInt(value)
			if err != nil {
				return false
			}

			limit, err := govalidator.ToInt(compare)
			if err != nil {
				return false
			}

			return number > limit
		},
		message: func(value string) string {
			return fmt.Sprintf("%s %s", communication.New().Mapping["less_or_equal_than"].Message, value)
		},
	},
}
