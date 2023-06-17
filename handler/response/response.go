package response

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

type APIResponse struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

func Response(message string, code int, status string, data interface{}) APIResponse {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	response := APIResponse{
		Meta: meta,
		Data: data,
	}

	return response

}
