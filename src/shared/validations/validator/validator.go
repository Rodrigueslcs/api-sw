package validator

var _instance *Validator = nil

type Validator struct {
	String *StringValidator
}

func New() *Validator {
	if _instance == nil {
		_instance = &Validator{
			String: &StringValidator{},
		}
	}

	return _instance
}
