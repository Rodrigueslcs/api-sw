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
	data["success"] = mapItem{Message: "Success", Code: 100001}
	data["success_create"] = mapItem{Message: "Record successfully created", Code: 100002}
	data["success_update"] = mapItem{Message: "Registro atualizado com sucesso", Code: 100003}
	data["success_delete"] = mapItem{Message: "Record successfully deleted", Code: 100004}
	data["success_search"] = mapItem{Message: "Record found", Code: 100005}
	data["error_create"] = mapItem{Message: "Unable to create record", Code: 100006}
	data["error_list"] = mapItem{Message: "Unable to list record", Code: 100007}
	data["error_search"] = mapItem{Message: "Record not found", Code: 100008}
	data["validate_failed"] = mapItem{Message: "Validation failed", Code: 100009}
	data["planet_not_found"] = mapItem{Message: "Planet not found", Code: 1000010}
	data["error_update"] = mapItem{Message: "Unable to update record", Code: 100011}
	data["error_delete"] = mapItem{Message: "Unable to delete record", Code: 100012}

	response.Mapping = data
}
