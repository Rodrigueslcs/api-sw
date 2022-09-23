package communication

type mapItem struct {
	Message string
	Code    int
}

type ResponseMapping struct {
	Mapping map[string]mapItem `json:"mapping"`
}

var singletonResponseMapping *ResponseMapping

func New() ResponseMapping {
	if singletonResponseMapping == nil {
		mapping := ResponseMapping{}
		mapping.populate()
		singletonResponseMapping = &mapping
	}
	return *singletonResponseMapping
}

func (response *ResponseMapping) Response(status int, identifier string, data any) Response {
	return Response{
		Status:  status,
		Code:    response.Mapping[identifier].Code,
		Message: response.Mapping[identifier].Message,
		Data:    data,
	}
}

func (response *ResponseMapping) ResponseError(status int, identifier string, err error) Response {
	return Response{
		Status:  status,
		Code:    response.Mapping[identifier].Code,
		Message: response.Mapping[identifier].Message,
		Error:   err,
	}
}

func (response *ResponseMapping) populate() {
	data := make(map[string]mapItem)
	data["already_exist"] = mapItem{Message: "Alread exists", Code: 100000}
	data["success"] = mapItem{Message: "Success", Code: 100006}
	data["success_create"] = mapItem{Message: "Record successfully created", Code: 100009}
	data["error_create"] = mapItem{Message: "Unable to create record", Code: 100012}
	data["error_list"] = mapItem{Message: "Unable to list record", Code: 100015}
	data["validate_failed"] = mapItem{Message: "Validation failed", Code: 100018}

	response.Mapping = data
}
