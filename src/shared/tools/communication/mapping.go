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
	data["success_create"] = mapItem{Message: "Record successfully created", Code: 100009}

	response.Mapping = data
}
