package validations

import (
	"api-sw/src/shared/tools/communication"
	"api-sw/src/shared/validations/validator"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
)

type ValidationError struct {
	Index   *int   `json:"index,omitempty"`
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationErrors []*ValidationError

func ValidateStruct(data interface{}) ValidationErrors {
	vl := reflect.ValueOf(data)

	return validateStruct(vl)
}

func validateStruct(vl reflect.Value) (validationErrors ValidationErrors) {
	if vl.Kind() == reflect.Ptr || vl.Kind() == reflect.Struct {
		var elm reflect.Value

		if vl.Kind() == reflect.Ptr {
			elm = vl.Elem()
		} else {
			elm = vl
		}

		for i := 0; i < elm.NumField(); i++ {
			valueField := elm.Field(i)
			typeField := elm.Type().Field(i)
			field := elm.Type().Field(i).Name
			tagValidate := typeField.Tag.Get("validate")

			if tagValidate == "" {
				continue
			}

			//verify field is public
			if validator.Matches(strings.Split(field, "")[0], "[A-Z]") {
				//If the field type is a struct then validator individual
				if valueField.Kind() == reflect.Struct {
					switch valueField.Interface().(type) {
					case time.Time:
						//Struct time
						validatorErrors := validateDate(field, tagValidate, valueField)
						validationErrors = append(validationErrors, validatorErrors...)
					default:
						//Simple struct
						validatorErrors := validateStruct(valueField)
						validationErrors = append(validationErrors, validatorErrors...)
					}
				} else {
					//If the type is a slice then loop and validator one by one
					if valueField.Kind() == reflect.Slice {
						//Simple slices
						validatorErrors := validateSimpleSlice(field, tagValidate, valueField)
						validationErrors = append(validationErrors, validatorErrors...)
					} else {
						//Simple Field
						validatorErrors := validateField(field, tagValidate, valueField)
						validationErrors = append(validationErrors, validatorErrors...)
					}
				}
			}
		}
	}
	return
}

func validateSimpleSlice(field string, tagValidate string, slice reflect.Value) (validationErrors ValidationErrors) {
	comm := communication.New()

	if slice.Kind() == reflect.Slice {
		elementsOfSlice := reflect.ValueOf(slice.Interface())
		quantityElements := elementsOfSlice.Len()
		tagsOfSlice := getTags(tagValidate)

		if quantityElements == 0 {
			for x := range tagsOfSlice {
				if tagsOfSlice[x] == "required" {
					validationErrors = append(validationErrors, &ValidationError{Field: field, Message: comm.Mapping["validate_required"].Message})
				}
			}
		} else {
			for i := 0; i < quantityElements; i++ {
				valueField := elementsOfSlice.Index(i)

				if valueField.Kind() == reflect.Struct {
					validateErrors := validateStruct(valueField)

					for _, err := range validateErrors {
						err.Index = getIntPointer(i)
					}

					validationErrors = append(validationErrors, validateErrors...)
				} else {
					if valueField.Elem().Kind() == reflect.Struct {
						validateErrors := validateStruct(valueField)

						for _, err := range validateErrors {
							err.Index = getIntPointer(i)
						}

						validationErrors = append(validationErrors, validateErrors...)
					} else {
						validateErrors := validateField(field, tagValidate, valueField)

						for _, err := range validateErrors {
							err.Index = getIntPointer(i)
						}

						validationErrors = append(validationErrors, validateErrors...)
					}
				}
			}
		}
	}
	return
}

func validateField(field, tagValidate string, valueField reflect.Value) (validationErrors ValidationErrors) {
	comm := communication.New()
	tags := getTags(tagValidate)

	re, _ := regexp.Compile("isLessThan|isLessOrEqualThan|isGreaterThan|isGreaterOrEqualThan")

	if len(tags) >= 1 {
		//All values to string
		value := fmt.Sprint(valueField)

		//Efetuando looping nas tags e efetuando as validações
		for x := range tags {
			var tagValue string

			tagsSplited := strings.Split(tags[x], "=")
			tag := tagsSplited[0]

			if len(tagsSplited) > 1 {
				tagValue = tagsSplited[1]
			}

			if tag == "isPassword" && value != "" {
				check := govalidator.IsByteLength(value, 6, 40)
				if !check {
					validationErrors = append(validationErrors, &ValidationError{Field: field, Message: comm.Mapping["validator_password_length"].Message})
				}
			} else if tag == "email" {
				check := govalidator.IsEmail(value)
				if !check {
					validationErrors = append(validationErrors, &ValidationError{Field: field, Message: comm.Mapping["validator_password_length"].Message})
				}
			} else if tag == "isNotZero" {
				check := value != "0" && value != "0.0"
				if !check {
					validationErrors = append(validationErrors, &ValidationError{Field: field, Message: comm.Mapping["validate_required"].Message})
				}
			} else if tag == "gte" || tag == "lte" {
				numberValidations, exists := numberValidator[tag]

				if exists {
					check := numberValidations.validator(value, tagValue)

					if check {
						validationErrors = append(validationErrors, &ValidationError{Field: field, Message: numberValidations.message(tagValue)})
					}
				}
			} else if re.MatchString(tag) {
				stringValidations, exists := stringLengthValidator[tag]

				if exists {
					invalid := stringValidations.validator(value, tagValue)

					if invalid {
						validationErrors = append(validationErrors, &ValidationError{Field: field, Message: stringValidations.message(tagValue)})
					}
				}
			} else {
				simpleValidations, exists := simpleValidator[tag]

				if exists {
					check := simpleValidations.validator(value)
					if check {
						validationErrors = append(validationErrors, &ValidationError{Field: field, Message: simpleValidations.message})
					}
				}
			}
		}
	}
	return
}

func validateDate(field, tagValidate string, valueField reflect.Value) (validationErrors ValidationErrors) {
	comm := communication.New()
	tags := getTags(tagValidate)

	if len(tags) >= 1 {
		//All values to string
		value := fmt.Sprint(valueField)

		//Efetuando looping nas tags e efetuando as validações
		for x := range tags {
			if tags[x] == "required" {
				valid, _ := validator.IsDate(value)
				fmt.Println(validator.IsDate(value))
				if !valid {
					validationErrors = append(validationErrors, &ValidationError{Field: field, Message: comm.Mapping["validator_date"].Message})
				}
			}
		}
	}
	return
}

func getTags(tagValidate string) (tags []string) {
	if tagValidate != "" {
		//Broken tags per |
		tags = strings.Split(tagValidate, ",")

		//If not broken per | so try per ;
		if len(tags) == 1 {
			tags = strings.Split(tags[0], ";")
		}
	}
	return
}

func getIntPointer(val int) *int {
	return &val
}
