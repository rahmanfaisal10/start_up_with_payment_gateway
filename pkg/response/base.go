package response

type BaseResponse struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data,omitempty"`
}

type Meta struct {
	Message string
	Code    int
	Status  string
}
