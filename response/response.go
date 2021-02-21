package response

type BaseResponse struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string
	Code    int
	Status  string
}

func InitResponse() *BaseResponse {
	return &BaseResponse{}
}

func ResponseAPI(message, status string, code int, formatter interface{}) BaseResponse {
	return BaseResponse{
		Meta: Meta{
			Message: message,
			Code:    code,
			Status:  status,
		},
		Data: formatter,
	}
}
